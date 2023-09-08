package perp

import (
	"time"

	"github.com/NibiruChain/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/common/asset"
	"github.com/NibiruChain/nibiru/x/perp/v2/keeper"
	types "github.com/NibiruChain/nibiru/x/perp/v2/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	for _, m := range genState.Markets {
		k.Markets.Insert(ctx, collections.Join(m.Pair, m.Version), m)
	}

	for _, g := range genState.MarketLastVersions {
		k.MarketLastVersion.Insert(ctx, g.Pair, types.MarketLastVersion{Version: g.Version})
	}

	for _, a := range genState.Amms {
		pair := a.Pair
		k.AMMs.Insert(ctx, collections.Join(a.Pair, a.Version), a)
		timestampMs := ctx.BlockTime().UnixMilli()
		k.ReserveSnapshots.Insert(
			ctx,
			collections.Join(pair, time.UnixMilli(timestampMs)),
			types.ReserveSnapshot{
				Amm:         a,
				TimestampMs: timestampMs,
			},
		)
	}

	for _, p := range genState.Positions {
		k.Positions.Insert(
			ctx,
			collections.Join(p.Pair, sdk.MustAccAddressFromBech32(p.TraderAddress)),
			p,
		)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := new(types.GenesisState)

	genesis.Markets = k.Markets.Iterate(ctx, collections.Range[collections.Pair[asset.Pair, uint64]]{}).Values()

	kv := k.MarketLastVersion.Iterate(ctx, collections.Range[asset.Pair]{}).KeyValues()
	for _, kv := range kv {
		genesis.MarketLastVersions = append(genesis.MarketLastVersions, types.GenesisMarketLastVersion{
			Pair:    kv.Key,
			Version: kv.Value.Version,
		})
	}

	genesis.Amms = k.AMMs.Iterate(ctx, collections.Range[collections.Pair[asset.Pair, uint64]]{}).Values()
	genesis.Positions = k.Positions.Iterate(ctx, collections.PairRange[asset.Pair, sdk.AccAddress]{}).Values()
	genesis.ReserveSnapshots = k.ReserveSnapshots.Iterate(ctx, collections.PairRange[asset.Pair, time.Time]{}).Values()

	return genesis
}
