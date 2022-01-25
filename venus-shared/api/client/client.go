package client

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/venus/venus-shared/api"
	v0 "github.com/filecoin-project/venus/venus-shared/api/chain/v0"
	v1 "github.com/filecoin-project/venus/venus-shared/api/chain/v1"
)

func NewFullRPCV0(ctx context.Context, addr string, header http.Header) (v0.FullNode, jsonrpc.ClientCloser, error) {
	var full v0.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin", api.GetInternalStructs(&full),
		header)
	return &full, closer, err
}

func NewFullRPCV1(ctx context.Context, addr string, header http.Header) (v1.FullNode, jsonrpc.ClientCloser, error) {
	var full v1.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin", api.GetInternalStructs(&full),
		header)
	return &full, closer, err
}

func NewWalletRPC(ctx context.Context, addr string, header http.Header) (v1.IWallet, jsonrpc.ClientCloser, error) {
	var wallet v1.IWalletStruct
	closer, err := jsonrpc.NewClient(ctx, addr, "Filecoin", &wallet.Internal, header)
	return &wallet, closer, err
}
