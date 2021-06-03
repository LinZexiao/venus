package consensus

import (
	"context"

	proof5 "github.com/filecoin-project/specs-actors/v5/actors/runtime/proof"
	"github.com/filecoin-project/venus/pkg/constants"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/venus/pkg/crypto/sigs"
)

// Interface to PoSt verification, modify by force EPoStVerifier -> ProofVerifier
type ProofVerifier interface {
	VerifySeal(info proof5.SealVerifyInfo) (bool, error)
	VerifyAggregateSeals(aggregate proof5.AggregateSealVerifyProofAndInfos) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof5.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof5.WindowPoStVerifyInfo) (bool, error)
	GenerateWinningPoStSectorChallenge(ctx context.Context, proofType abi.RegisteredPoStProof, minerID abi.ActorID, randomness abi.PoStRandomness, eligibleSectorCount uint64) ([]uint64, error)
}

type SignFunc func(context.Context, address.Address, []byte) (*crypto.Signature, error)

func VerifyVRF(ctx context.Context, worker address.Address, vrfBase, vrfproof []byte) error {
	_, span := trace.StartSpan(ctx, "VerifyVRF")
	defer span.End()

	sig := &crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: vrfproof,
	}

	if err := sigs.Verify(sig, worker, vrfBase); err != nil {
		return xerrors.Errorf("vrf was invalid: %w", err)
	}

	return nil
}

func VerifyElectionPoStVRF(ctx context.Context, worker address.Address, rand []byte, evrf []byte) error {
	if constants.InsecurePoStValidation {
		return nil
	}
	return VerifyVRF(ctx, worker, rand, evrf)
}
