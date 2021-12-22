package apiface

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	types "github.com/filecoin-project/venus/venus-shared/chain"
	"github.com/filecoin-project/venus/venus-shared/messagepool"
)

type IMultiSig interface {
	// Rule[perm:sign]
	MsigCreate(ctx context.Context, req uint64, addrs []address.Address, duration abi.ChainEpoch, val types.BigInt, src address.Address, gp types.BigInt) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigPropose(ctx context.Context, msig address.Address, to address.Address, amt types.BigInt, src address.Address, method uint64, params []byte) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigAddPropose(ctx context.Context, msig address.Address, src address.Address, newAdd address.Address, inc bool) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigAddApprove(ctx context.Context, msig address.Address, src address.Address, txID uint64, proposer address.Address, newAdd address.Address, inc bool) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigAddCancel(ctx context.Context, msig address.Address, src address.Address, txID uint64, newAdd address.Address, inc bool) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigSwapPropose(ctx context.Context, msig address.Address, src address.Address, oldAdd address.Address, newAdd address.Address) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigSwapApprove(ctx context.Context, msig address.Address, src address.Address, txID uint64, proposer address.Address, oldAdd address.Address, newAdd address.Address) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigSwapCancel(ctx context.Context, msig address.Address, src address.Address, txID uint64, oldAdd address.Address, newAdd address.Address) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigApprove(ctx context.Context, msig address.Address, txID uint64, src address.Address) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigApproveTxnHash(ctx context.Context, msig address.Address, txID uint64, proposer address.Address, to address.Address, amt types.BigInt, src address.Address, method uint64, params []byte) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigCancel(ctx context.Context, msig address.Address, txID uint64, to address.Address, amt types.BigInt, src address.Address, method uint64, params []byte) (*messagepool.MessagePrototype, error)
	// Rule[perm:sign]
	MsigRemoveSigner(ctx context.Context, msig address.Address, proposer address.Address, toRemove address.Address, decrease bool) (*messagepool.MessagePrototype, error)
	// Rule[perm:read]
	MsigGetVested(ctx context.Context, addr address.Address, start types.TipSetKey, end types.TipSetKey) (types.BigInt, error)
}
