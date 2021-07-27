package networks

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/venus/pkg/constants"

	"github.com/filecoin-project/venus/pkg/config"
)

func Net2k() *NetworkConf {
	return &NetworkConf{
		Bootstrap: config.BootstrapConfig{
			Addresses:        []string{},
			MinPeerThreshold: 0,
			Period:           "30s",
		},
		Network: config.NetworkParamsConfig{
			NetworkType:            constants.Network2k,
			BlockDelay:             4,
			ConsensusMinerMinPower: 2048,
			MinVerifiedDealSize:    256,
			ReplaceProofTypes: []abi.RegisteredSealProof{
				abi.RegisteredSealProof_StackedDrg2KiBV1,
			},
			ForkUpgradeParam: &config.ForkUpgradeConfig{
				UpgradeBreezeHeight:        -1,
				UpgradeSmokeHeight:         -1,
				UpgradeIgnitionHeight:      -2,
				UpgradeRefuelHeight:        -3,
				UpgradeTapeHeight:          -4,
				UpgradeAssemblyHeight:      -5,
				UpgradeLiftoffHeight:       -6,
				UpgradeKumquatHeight:       -7,
				UpgradePriceListOopsHeight: -8,
				UpgradeCalicoHeight:        -9,
				UpgradePersianHeight:       -10,
				UpgradeOrangeHeight:        -11,
				UpgradeClausHeight:         -12,
				UpgradeTrustHeight:         -13,
				UpgradeNorwegianHeight:     -14,
				UpgradeTurboHeight:         -15,
				UpgradeHyperdriveHeight:    -16,

				BreezeGasTampingDuration: 0,
			},
			DrandSchedule:           map[abi.ChainEpoch]config.DrandEnum{0: 1},
			AddressNetwork:          address.Testnet,
			PreCommitChallengeDelay: abi.ChainEpoch(150),
		},
	}
}
