module github.com/filecoin-project/venus/venus-devtool

go 1.16

require (
	github.com/filecoin-project/venus/venus-shared v0.0.1
	github.com/whyrusleeping/cbor-gen v0.0.0-20211110122933-f57984553008
)

replace github.com/filecoin-project/venus/venus-shared => ../venus-shared/
