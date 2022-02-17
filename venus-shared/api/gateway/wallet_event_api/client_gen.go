// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package wallet_event_api

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"

	"github.com/filecoin-project/venus/venus-shared/api"
)

// NewIWalletEventAPIRPC creates a new httpparse jsonrpc remotecli.
func NewIWalletEventAPIRPC(ctx context.Context, addr string, requestHeader http.Header, opts ...jsonrpc.Option) (IWalletEventAPI, jsonrpc.ClientCloser, error) {
	if requestHeader == nil {
		requestHeader = http.Header{}
	}
	requestHeader.Set(api.VenusAPINamespaceHeader, "wallet_event_api.IWalletEventAPI")

	var res IWalletEventAPIStruct
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Gateway", api.GetInternalStructs(&res), requestHeader, opts...)

	return &res, closer, err
}
