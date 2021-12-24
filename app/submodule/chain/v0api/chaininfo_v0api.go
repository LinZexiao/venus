package v0api

import (
	"context"

	v0api "github.com/filecoin-project/venus/venus-shared/api/chain/v0"
	v1api "github.com/filecoin-project/venus/venus-shared/api/chain/v1"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/venus/pkg/constants"
	apitypes "github.com/filecoin-project/venus/venus-shared/api/chain"
	types "github.com/filecoin-project/venus/venus-shared/chain"
)

var _ v0api.IChain = &WrapperV1IChain{}

type WrapperV1IChain struct { //nolint
	v1api.IChain
}

func (a *WrapperV1IChain) StateSearchMsg(ctx context.Context, msg cid.Cid) (*apitypes.MsgLookup, error) {
	return a.IChain.StateSearchMsg(ctx, types.EmptyTSK, msg, constants.LookbackNoLimit, true)
}

func (a *WrapperV1IChain) StateSearchMsgLimited(ctx context.Context, msg cid.Cid, limit abi.ChainEpoch) (*apitypes.MsgLookup, error) {
	return a.IChain.StateSearchMsg(ctx, types.EmptyTSK, msg, limit, true)
}

func (a *WrapperV1IChain) StateWaitMsg(ctx context.Context, msg cid.Cid, confidence uint64) (*apitypes.MsgLookup, error) {
	return a.IChain.StateWaitMsg(ctx, msg, confidence, constants.LookbackNoLimit, true)
}

func (a *WrapperV1IChain) StateWaitMsgLimited(ctx context.Context, msg cid.Cid, confidence uint64, limit abi.ChainEpoch) (*apitypes.MsgLookup, error) {
	return a.IChain.StateWaitMsg(ctx, msg, confidence, limit, true)
}

func (a *WrapperV1IChain) StateGetReceipt(ctx context.Context, msg cid.Cid, from types.TipSetKey) (*types.MessageReceipt, error) {
	ml, err := a.IChain.StateSearchMsg(ctx, from, msg, constants.LookbackNoLimit, true)
	if err != nil {
		return nil, err
	}

	if ml == nil {
		return nil, nil
	}

	return &ml.Receipt, nil
}
