package common

import (
	"github.com/holiman/uint256"
)

const (
	ModuleName                = "common"
	TreasuryPoolModuleAccount = "treasury_pool"
	// Precision for int representation in sdk.Int objects
	Precision = int64(1_000_000)
)

var (
	APrecision = uint256.NewInt().SetUint64(1)
)