package keeper_test

import (
	"testing"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/perp/types"
	"github.com/NibiruChain/nibiru/x/testutil"
	"github.com/NibiruChain/nibiru/x/testutil/sample"
	vtypes "github.com/NibiruChain/nibiru/x/vpool/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func mocktwap(ctx sdk.Context, pair common.TokenPair, dir vtypes.Direction, abs sdk.Int) (sdk.Dec, error) {
	return sdk.NewDec(1), nil
}

func TestFullLiquidate(t *testing.T) {
	tests := []struct {
		name                      string
		positionSize              sdk.Dec
		initialEF                 sdk.Dec
		indexPrice                sdk.Dec
		otherPositionSize         sdk.Dec
		expectedErr               error
		expectedLiquidatorBalance sdk.Dec
	}{
		{
			name:              "happy path",
			initialEF:         sdk.NewDec(100),
			indexPrice:        sdk.MustNewDecFromStr("1"),
			positionSize:      sdk.NewDec(1000),
			otherPositionSize: sdk.NewDec(1000),
			expectedErr:       types.MarginHighEnough,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)

			tokenPair, err := common.NewTokenPairFromStr("atom:unusd")
			require.NoError(t, err)

			t.Log("add margin funds (NUSD) to traders' account")
			traderAddr := sample.AccAddress()
			otherTraderAddr := sample.AccAddress()
			liquidatorAddr := sample.AccAddress()

			// Create vPool to get the spot price
			app.VpoolKeeper.CreatePool(
				ctx,
				tokenPair.String(),
				sdk.MustNewDecFromStr("0.9"),  // 0.9 ratio
				sdk.NewDec(10_000_000),        // 10 tokens
				sdk.NewDec(5_000_000),         // 5 tokens
				sdk.MustNewDecFromStr("0.25"), // 0.25 ratio
				sdk.MustNewDecFromStr("0.25"), // 0.25 ratio
			)
			require.True(t, app.VpoolKeeper.ExistsPool(ctx, tokenPair))

			app.PerpKeeper.PairMetadata().Set(ctx, &types.PairMetadata{
				Pair:                       tokenPair.String(),
				CumulativePremiumFractions: []sdk.Dec{sdk.OneDec()},
			})

			err = simapp.FundAccount(
				app.BankKeeper,
				ctx,
				traderAddr,
				sdk.NewCoins(
					sdk.NewCoin(common.StableDenom, tc.positionSize.TruncateInt().Add(sdk.NewInt(2))), // Adding 2 for both fees (toll and spread)
				),
			)
			require.NoErrorf(t, err, "fund account call should work")
			err = simapp.FundAccount(
				app.BankKeeper,
				ctx,
				otherTraderAddr,
				sdk.NewCoins(
					sdk.NewCoin(common.StableDenom, tc.otherPositionSize.TruncateInt().Add(sdk.NewInt(2))),
				),
			)
			require.NoErrorf(t, err, "fund account call should work")

			t.Log("add liquidation funds to perp EF")
			err = simapp.FundModuleAccount(
				app.BankKeeper,
				ctx,
				common.TreasuryPoolModuleAccount,
				sdk.NewCoins(sdk.NewCoin(common.StableDenom, tc.initialEF.TruncateInt())),
			)
			require.NoErrorf(t, err, "fund module call should work")

			t.Log("establish initial position")
			err = app.PerpKeeper.OpenPosition(
				ctx, tokenPair, types.Side_BUY, traderAddr.String(), tc.positionSize, sdk.OneDec(), sdk.NewDec(150),
			)
			require.NoError(t, err, "initial position should be opened")
			err = app.PerpKeeper.OpenPosition(
				ctx, tokenPair, types.Side_BUY, otherTraderAddr.String(), tc.otherPositionSize, sdk.OneDec(), sdk.NewDec(150),
			)
			require.NoError(t, err, "second position should be opened")

			t.Log("liquidate position")
			err = app.PerpKeeper.Liquidate(ctx, tokenPair, traderAddr.String(), liquidatorAddr)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
