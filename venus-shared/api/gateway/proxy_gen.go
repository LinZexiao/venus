// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package gateway

import (
	"context"

	"github.com/filecoin-project/venus/venus-shared/types/gateway"
)

type IProofEventAPIStruct struct {
	Internal struct {
		ListenProofEvent   func(ctx context.Context, policy *gateway.ProofRegisterPolicy) (<-chan *gateway.RequestEvent, error) `perm:"write"`
		ResponseProofEvent func(ctx context.Context, resp *gateway.ResponseEvent) error                                         `perm:"write"`
	}
}

func (s *IProofEventAPIStruct) ListenProofEvent(p0 context.Context, p1 *gateway.ProofRegisterPolicy) (<-chan *gateway.RequestEvent, error) {
	return s.Internal.ListenProofEvent(p0, p1)
}
func (s *IProofEventAPIStruct) ResponseProofEvent(p0 context.Context, p1 *gateway.ResponseEvent) error {
	return s.Internal.ResponseProofEvent(p0, p1)
}
