package v0

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/venus/venus-shared/chain"
	"github.com/filecoin-project/venus/venus-shared/wallet"
)

type IWallet interface {
	WalletSign(ctx context.Context, k address.Address, msg []byte, meta wallet.MsgMeta) (*crypto.Signature, error) //perm:sign
	WalletExport(addr address.Address, password string) (*wallet.KeyInfo, error)                                   //perm:admin
	WalletImport(key *wallet.KeyInfo) (address.Address, error)                                                     //perm:admin
	WalletHas(ctx context.Context, addr address.Address) (bool, error)                                             //perm:write
	WalletNewAddress(protocol address.Protocol) (address.Address, error)                                           //perm:write
	WalletBalance(ctx context.Context, addr address.Address) (abi.TokenAmount, error)                              //perm:read
	WalletDefaultAddress(ctx context.Context) (address.Address, error)                                             //perm:write
	WalletAddresses(ctx context.Context) []address.Address                                                         //perm:admin
	WalletSetDefault(ctx context.Context, addr address.Address) error                                              //perm:write
	WalletSignMessage(ctx context.Context, k address.Address, msg *chain.Message) (*chain.SignedMessage, error)    //perm:sign
	LockWallet(ctx context.Context) error                                                                          //perm:admin
	UnLockWallet(ctx context.Context, password []byte) error                                                       //perm:admin
	SetPassword(Context context.Context, password []byte) error                                                    //perm:admin
	HasPassword(Context context.Context) bool                                                                      //perm:admin
	WalletState(Context context.Context) int                                                                       //perm:admin
}
