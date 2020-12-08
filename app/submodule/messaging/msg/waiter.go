package msg

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/venus/pkg/block"
	"github.com/filecoin-project/venus/pkg/constants"
	"github.com/ipfs/go-cid"
	bstore "github.com/ipfs/go-ipfs-blockstore"
	cbor "github.com/ipfs/go-ipld-cbor"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/venus/pkg/chain"
	"github.com/filecoin-project/venus/pkg/types"
	"github.com/filecoin-project/venus/pkg/vm/state"
)

var log = logging.Logger("messageimpl")

// Abstracts over a store of blockchain state.
type waiterChainReader interface {
	GetHead() block.TipSetKey
	GetTipSet(block.TipSetKey) (*block.TipSet, error)
	ResolveAddressAt(ctx context.Context, tipKey block.TipSetKey, addr address.Address) (address.Address, error)
	GetActorAt(ctx context.Context, tipKey block.TipSetKey, addr address.Address) (*types.Actor, error)
	GetTipSetState(context.Context, block.TipSetKey) (state.Tree, error)
	GetTipSetReceiptsRoot(block.TipSetKey) (cid.Cid, error)
	SubHeadChanges(ctx context.Context) chan []*chain.HeadChange
}

// Waiter waits for a message to appear on chain.
type Waiter struct {
	chainReader     waiterChainReader
	messageProvider chain.MessageProvider
	cst             cbor.IpldStore
	bs              bstore.Blockstore
}

// ChainMessage is an on-chain message with its block and receipt.
type ChainMessage struct {
	Ts      *block.TipSet
	Message types.ChainMsg
	Block   *block.Block
	Receipt *types.MessageReceipt
}

// WaitPredicate is a function that identifies a message and returns true when found.
type WaitPredicate func(msg *types.UnsignedMessage, msgCid cid.Cid) bool

// NewWaiter returns a new Waiter.
func NewWaiter(chainStore waiterChainReader, messages chain.MessageProvider, bs bstore.Blockstore, cst cbor.IpldStore) *Waiter {
	return &Waiter{
		chainReader:     chainStore,
		cst:             cst,
		bs:              bs,
		messageProvider: messages,
	}
}

// Find searches the blockchain history (but doesn't wait).
func (w *Waiter) Find(ctx context.Context, msg types.ChainMsg, lookback abi.ChainEpoch, ts *block.TipSet) (*ChainMessage, bool, error) {
	if ts == nil {
		var err error
		ts, err = w.chainReader.GetTipSet(w.chainReader.GetHead())
		if err != nil {
			return nil, false, err
		}
	}

	return w.findMessage(ctx, ts, msg, lookback)
}

// WaitPredicate invokes the callback when the passed predicate succeeds.
// See api description.
//
// Note: this method does too much -- the callback should just receive the tipset
// containing the message and the caller should pull the receipt out of the block
// if in fact that's what it wants to do, using something like receiptFromTipset.
// Something like receiptFromTipset is necessary because not every message in
// a block will have a receipt in the tipset: it might be a duplicate message.
// This method will always check for the message in the current head tipset.
// A lookback parameter > 1 will cause this method to check for the message in
// up to that many previous tipsets on the chain of the current head.
func (w *Waiter) WaitPredicate(ctx context.Context, msg types.ChainMsg, confidence abi.ChainEpoch, lookback abi.ChainEpoch) (*ChainMessage, error) {
	ch := w.chainReader.SubHeadChanges(ctx)
	chainMsg, found, err := w.waitForMessage(ctx, ch, msg, confidence, lookback)
	if err != nil {
		return nil, err
	}
	if found {
		return chainMsg, nil
	}
	return nil, nil
}

// Wait uses WaitPredicate to invoke the callback when a message with the given cid appears on chain.
func (w *Waiter) Wait(ctx context.Context, msg types.ChainMsg, confidence abi.ChainEpoch, lookbackLimit abi.ChainEpoch) (*ChainMessage, error) {
	mid, _ := msg.VMMessage().Cid()
	log.Infof("Calling Waiter.Wait CID: %s", mid.String())

	return w.WaitPredicate(ctx, msg, confidence, lookbackLimit)
}

// findMessage looks for a matching in the chain and returns the message,
// block and receipt, when it is found. Returns the found message/block or nil
// if now block with the given CID exists in the chain.
// The lookback parameter is the number of tipsets in the past this method will check before giving up.
func (w *Waiter) findMessage(ctx context.Context, from *block.TipSet, m types.ChainMsg, lookback abi.ChainEpoch) (*ChainMessage, bool, error) {
	limitHeight := from.EnsureHeight() - lookback
	noLimit := lookback == constants.LookbackNoLimit

	cur := from
	curActor, err := w.chainReader.GetActorAt(ctx, cur.Key(), m.VMMessage().From)
	if err != nil {
		return nil, false, xerrors.Errorf("failed to load initital tipset")
	}

	mFromID, err := w.chainReader.ResolveAddressAt(ctx, from.Key(), m.VMMessage().From)
	if err != nil {
		return nil, false, xerrors.Errorf("looking up From id address: %w", err)
	}

	mNonce := m.VMMessage().Nonce

	for {
		// If we've reached the genesis block, or we've reached the limit of
		// how far back to look
		if cur.EnsureHeight() == 0 || !noLimit && cur.EnsureHeight() <= limitHeight {
			// it ain't here!
			return nil, false, nil
		}

		select {
		case <-ctx.Done():
			return nil, false, nil
		default:
		}

		// we either have no messages from the sender, or the latest message we found has a lower nonce than the one being searched for,
		// either way, no reason to lookback, it ain't there
		if curActor == nil || curActor.Nonce == 0 || curActor.Nonce < mNonce {
			return nil, false, nil
		}

		pts, err := w.chainReader.GetTipSet(cur.EnsureParents())
		if err != nil {
			return nil, false, xerrors.Errorf("failed to load tipset during msg wait searchback: %w", err)
		}

		act, err := w.chainReader.GetActorAt(ctx, pts.Key(), mFromID)
		actorNoExist := errors.Is(err, types.ErrActorNotFound)
		if err != nil && !actorNoExist {
			return nil, false, xerrors.Errorf("failed to load the actor: %w", err)
		}

		// check that between cur and parent tipset the nonce fell into range of our message
		if actorNoExist || (curActor.Nonce > mNonce && act.Nonce <= mNonce) {
			msg, found, err := w.receiptForTipset(ctx, cur, m)
			if err != nil {
				log.Errorf("Waiter.Wait: %s", err)
				return nil, false, err
			}
			if found {
				return msg, true, nil
			}
		}

		cur = pts
		curActor = act
	}
}

// waitForMessage looks for a matching message in a channel of tipsets and returns
// the message, block and receipt, when it is found. Reads until the channel is
// closed or the context done. Returns the found message/block (or nil if the
// channel closed without finding it), whether it was found, or an error.
func (w *Waiter) waitForMessage(ctx context.Context, ch <-chan []*chain.HeadChange, msg types.ChainMsg, confidence abi.ChainEpoch, lookbackLimit abi.ChainEpoch) (*ChainMessage, bool, error) {
	current, ok := <-ch
	if !ok {
		return nil, false, fmt.Errorf("SubHeadChanges stream was invalid")
	}
	//todo message wait
	if len(current) != 1 {
		return nil, false, fmt.Errorf("SubHeadChanges first entry should have been one item")
	}

	if current[0].Type != chain.HCCurrent {
		return nil, false, fmt.Errorf("expected current head on SHC stream (got %s)", current[0].Type)
	}

	currentHead := current[0].Val
	chainMsg, found, err := w.receiptForTipset(ctx, currentHead, msg)
	if err != nil {
		return nil, false, err
	}
	if found {
		return chainMsg, found, nil
	}

	var backRcp *ChainMessage
	backSearchWait := make(chan struct{})
	go func() {
		r, foundMsg, err := w.findMessage(ctx, currentHead, msg, lookbackLimit)
		if err != nil {
			log.Warnf("failed to look back through chain for message: %w", err)
			return
		}
		if foundMsg {
			backRcp = r
			close(backSearchWait)
		}
	}()

	var candidateTs *block.TipSet
	var candidateRcp *ChainMessage
	heightOfHead := currentHead.EnsureHeight()
	reverts := map[string]bool{}

	for {
		select {
		case notif, ok := <-ch:
			if !ok {
				return nil, false, err
			}
			for _, val := range notif {
				switch val.Type {
				case chain.HCRevert:
					if val.Val.Equals(candidateTs) {
						candidateTs = nil
						candidateRcp = nil
					}
					if backSearchWait != nil {
						reverts[val.Val.Key().String()] = true
					}
				case chain.HCApply:
					if candidateTs != nil && val.Val.EnsureHeight() >= candidateTs.EnsureHeight()+confidence {
						return candidateRcp, true, nil
					}

					r, foundMsg, err := w.receiptForTipset(ctx, val.Val, msg)
					if err != nil {
						return nil, false, err
					}
					if r != nil {
						if confidence == 0 {
							return r, foundMsg, err
						}
						candidateTs = val.Val
						candidateRcp = r
					}
					heightOfHead = val.Val.EnsureHeight()
				}
			}
		case <-backSearchWait:
			// check if we found the message in the chain and that is hasn't been reverted since we started searching
			if backRcp != nil && !reverts[backRcp.Ts.Key().String()] {
				// if head is at or past confidence interval, return immediately
				if heightOfHead >= backRcp.Ts.EnsureHeight()+confidence {
					return backRcp, true, nil
				}

				// wait for confidence interval
				candidateTs = backRcp.Ts
				candidateRcp = backRcp
			}
			reverts = nil
			backSearchWait = nil
		case <-ctx.Done():
			return nil, false, err
		}
	}
}

func (w *Waiter) receiptForTipset(ctx context.Context, ts *block.TipSet, msg types.ChainMsg) (*ChainMessage, bool, error) {
	blockMessageInfos, err := w.messageProvider.LoadTipSetMessage(ctx, ts)
	if err != nil {
		return nil, false, err
	}
	expectedCid, _ := msg.Cid()

	for _, bms := range blockMessageInfos {
		for _, msg := range append(bms.BlsMessages, bms.SecpkMessages...) {
			msgCid, err := msg.Cid()
			if err != nil {
				return nil, false, err
			}
			if expectedCid == msgCid {
				recpt, err := w.receiptByIndex(ctx, ts.Key(), msgCid, blockMessageInfos)
				if err != nil {
					return nil, false, errors.Wrap(err, "error retrieving receipt from tipset")
				}
				return &ChainMessage{ts, msg, bms.Block, recpt}, true, nil
			}
		}

	}
	return nil, false, nil
}

func (w *Waiter) receiptByIndex(ctx context.Context, tsKey block.TipSetKey, targetCid cid.Cid, blockMsgs []block.BlockMessagesInfo) (*types.MessageReceipt, error) {
	receiptCid, err := w.chainReader.GetTipSetReceiptsRoot(tsKey)
	if err != nil {
		return nil, err
	}

	receipts, err := w.messageProvider.LoadReceipts(ctx, receiptCid)
	if err != nil {
		return nil, err
	}

	receiptIndex := 0
	for _, blkInfo := range blockMsgs {
		//todo aggrate bls and secp msg to one msg
		for _, msg := range append(blkInfo.BlsMessages, blkInfo.SecpkMessages...) {
			msgCid, err := msg.Cid()
			if err != nil {
				return nil, err
			}

			if msgCid.Equals(targetCid) {
				if receiptIndex >= len(receipts) {
					return nil, errors.Errorf("could not find message receipt at index %d", receiptIndex)
				}
				return &receipts[receiptIndex], nil
			}
			receiptIndex++
		}
	}
	return nil, errors.Errorf("could not find message cid %s in dedupped messages", targetCid.String())
}
