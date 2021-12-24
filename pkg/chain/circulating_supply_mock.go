package chain

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/venus/pkg/state/tree"
	types "github.com/filecoin-project/venus/venus-shared/chain"
)

var _ ICirculatingSupplyCalcualtor = (*MockCirculatingSupplyCalculator)(nil)

type MockCirculatingSupplyCalculator struct {
}

func NewMockCirculatingSupplyCalculator() ICirculatingSupplyCalcualtor {
	return &MockCirculatingSupplyCalculator{}
}

func (m MockCirculatingSupplyCalculator) GetCirculatingSupplyDetailed(ctx context.Context, height abi.ChainEpoch, st tree.Tree) (types.CirculatingSupply, error) {
	return types.CirculatingSupply{
		FilVested:           abi.TokenAmount{},
		FilMined:            abi.TokenAmount{},
		FilBurnt:            abi.TokenAmount{},
		FilLocked:           abi.TokenAmount{},
		FilCirculating:      abi.TokenAmount{},
		FilReserveDisbursed: abi.TokenAmount{},
	}, nil
}
