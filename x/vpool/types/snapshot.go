package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/common"
)

func NewReserveSnapshot(ctx sdk.Context, pair common.AssetPair, baseAssetReserve, quoteAssetReserve sdk.Dec) ReserveSnapshot {
	return ReserveSnapshot{
		Pair:              pair.String(),
		BaseAssetReserve:  baseAssetReserve,
		QuoteAssetReserve: quoteAssetReserve,
		TimestampMs:       ctx.BlockTime().UnixMilli(),
		BlockNumber:       ctx.BlockHeight(),
	}
}
