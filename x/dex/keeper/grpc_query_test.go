package keeper_test

import (
	"testing"

	"github.com/NibiruChain/nibiru/x/dex/keeper"
	"github.com/NibiruChain/nibiru/x/dex/types"
	"github.com/NibiruChain/nibiru/x/testutil"
	"github.com/NibiruChain/nibiru/x/testutil/mock"
	"github.com/NibiruChain/nibiru/x/testutil/sample"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	app, ctx := testutil.NewNibiruApp(true)

	params := types.DefaultParams()
	app.DexKeeper.SetParams(ctx, params)

	queryServer := keeper.NewQuerier(app.DexKeeper)

	response, err := queryServer.Params(sdk.WrapSDKContext(ctx), &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

func TestQueryPoolHappyPath(t *testing.T) {
	tests := []struct {
		name         string
		existingPool types.Pool
	}{
		{
			name: "correct fetch pool",
			existingPool: types.Pool{
				Id:      1,
				Address: sample.AccAddress().String(),
				PoolParams: types.PoolParams{
					SwapFee: sdk.MustNewDecFromStr("0.03"),
					ExitFee: sdk.MustNewDecFromStr("0.03"),
				},
				PoolAssets: []types.PoolAsset{
					{
						Token:  sdk.NewInt64Coin("bar", 100),
						Weight: sdk.OneInt(),
					},
					{
						Token:  sdk.NewInt64Coin("bar", 100),
						Weight: sdk.OneInt(),
					},
				},
				TotalWeight: sdk.NewInt(2),
				TotalShares: sdk.NewInt64Coin("nibiru/pool/1", 200),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)
			app.DexKeeper.SetPool(ctx, tc.existingPool)

			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.Pool(sdk.WrapSDKContext(ctx), &types.QueryPoolRequest{
				PoolId: 1,
			})
			require.NoError(t, err)
			require.Equal(t, tc.existingPool, *resp.Pool)
		})
	}
}

func TestQueryPoolFail(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "invalid request",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)
			queryServer := keeper.NewQuerier(app.DexKeeper)
			resp, err := queryServer.Pool(sdk.WrapSDKContext(ctx), nil)
			require.Error(t, err)
			require.Nil(t, resp)
		})
	}
}

func TestQueryPools(t *testing.T) {
	tests := []struct {
		name          string
		existingPools []types.Pool
		pagination    query.PageRequest
		expectedPools []types.Pool
	}{
		{
			name: "successful query single pool",
			existingPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("foo", 100),
						sdk.NewInt64Coin("bar", 100),
					),
					/*shares=*/ 100,
				),
			},
			expectedPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("foo", 100),
						sdk.NewInt64Coin("bar", 100),
					),
					/*shares=*/ 100,
				),
			},
		},
		{
			name: "successful query multiple pools",
			existingPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("foo", 100),
						sdk.NewInt64Coin("bar", 100),
					),
					/*shares=*/ 100,
				),
				mock.DexPool(
					/*poolId=*/ 2,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("bar", 100),
						sdk.NewInt64Coin("baz", 100),
					),
					/*shares=*/ 100,
				),
			},
			expectedPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("foo", 100),
						sdk.NewInt64Coin("bar", 100),
					),
					/*shares=*/ 100,
				),
				mock.DexPool(
					/*poolId=*/ 2,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("bar", 100),
						sdk.NewInt64Coin("baz", 100),
					),
					/*shares=*/ 100,
				),
			},
		},
		{
			name: "query pools with pagination",
			existingPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("foo", 100),
						sdk.NewInt64Coin("bar", 100),
					),
					/*shares=*/ 100,
				),
				mock.DexPool(
					/*poolId=*/ 2,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("bar", 100),
						sdk.NewInt64Coin("baz", 100),
					),
					/*shares=*/ 100,
				),
			},
			pagination: query.PageRequest{
				Limit: 1,
			},
			expectedPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("foo", 100),
						sdk.NewInt64Coin("bar", 100),
					),
					/*shares=*/ 100,
				),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)
			for _, existingPool := range tc.existingPools {
				app.DexKeeper.SetPool(ctx, existingPool)
			}

			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.Pools(
				sdk.WrapSDKContext(ctx),
				&types.QueryPoolsRequest{
					Pagination: &tc.pagination,
				},
			)
			require.NoError(t, err)

			var responsePools []types.Pool
			for _, p := range resp.Pools {
				responsePools = append(responsePools, *p)
			}
			require.Equal(t, tc.expectedPools, responsePools)
		})
	}
}

func TestQueryNumPools(t *testing.T) {
	tests := []struct {
		name             string
		newPools         []types.Pool
		expectedNumPools uint64
	}{
		{
			name: "one pool",
			newPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("unibi", 100),
						sdk.NewInt64Coin("unusd", 100),
					),
					/*shares=*/ 100,
				),
			},
			expectedNumPools: 1,
		},
		{
			name: "multiple pools",
			newPools: []types.Pool{
				mock.DexPool(
					/*poolId=*/ 1,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("unibi", 100),
						sdk.NewInt64Coin("uust", 100),
					),
					/*shares=*/ 100,
				),
				mock.DexPool(
					/*poolId=*/ 2,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("uust", 100),
						sdk.NewInt64Coin("unusd", 100),
					),
					/*shares=*/ 100,
				),
				mock.DexPool(
					/*poolId=*/ 3,
					/*assets=*/ sdk.NewCoins(
						sdk.NewInt64Coin("unusd", 100),
						sdk.NewInt64Coin("unibi", 100),
					),
					/*shares=*/ 100,
				),
			},
			expectedNumPools: 3,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)
			sender := sample.AccAddress()
			// need funds to create pools
			require.NoError(t, simapp.FundAccount(
				app.BankKeeper,
				ctx,
				sender,
				sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 1e18),
					sdk.NewInt64Coin("unusd", 1e18),
					sdk.NewInt64Coin("uust", 1e18),
				),
			))

			for _, newPool := range tc.newPools {
				_, err := app.DexKeeper.NewPool(
					ctx,
					sender,
					newPool.PoolParams,
					newPool.PoolAssets,
				)
				require.NoError(t, err)
			}

			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.NumPools(
				sdk.WrapSDKContext(ctx),
				&types.QueryNumPoolsRequest{},
			)
			require.NoError(t, err)
			require.Equal(t, tc.expectedNumPools, resp.NumPools)
		})
	}
}

func TestQueryPoolParams(t *testing.T) {
	tests := []struct {
		name               string
		existingPool       types.Pool
		expectedPoolParams types.PoolParams
	}{
		{
			name: "successful fetch pool params",
			existingPool: types.Pool{
				Id:      1,
				Address: sample.AccAddress().String(),
				PoolParams: types.PoolParams{
					SwapFee: sdk.MustNewDecFromStr("0.03"),
					ExitFee: sdk.MustNewDecFromStr("0.03"),
				},
				PoolAssets: []types.PoolAsset{
					{
						Token:  sdk.NewInt64Coin("bar", 100),
						Weight: sdk.OneInt(),
					},
					{
						Token:  sdk.NewInt64Coin("bar", 100),
						Weight: sdk.OneInt(),
					},
				},
				TotalWeight: sdk.NewInt(2),
				TotalShares: sdk.NewInt64Coin("nibiru/pool/1", 200),
			},
			expectedPoolParams: types.PoolParams{
				SwapFee: sdk.MustNewDecFromStr("0.03"),
				ExitFee: sdk.MustNewDecFromStr("0.03"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)
			app.DexKeeper.SetPool(ctx, tc.existingPool)

			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.PoolParams(sdk.WrapSDKContext(ctx), &types.QueryPoolParamsRequest{
				PoolId: 1,
			})
			require.NoError(t, err)
			require.Equal(t, tc.expectedPoolParams, *resp.PoolParams)
		})
	}
}

func TestQueryTotalShares(t *testing.T) {
	tests := []struct {
		name                string
		existingPool        types.Pool
		expectedTotalShares sdk.Coin
	}{
		{
			name: "successfully get existing shares",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin("unusd", 100),
				),
				/*shares=*/ 100,
			),
			expectedTotalShares: sdk.NewInt64Coin("nibiru/pool/1", 100),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)

			app.DexKeeper.SetPool(ctx, tc.existingPool)

			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.TotalShares(
				sdk.WrapSDKContext(ctx),
				&types.QueryTotalSharesRequest{
					PoolId: 1,
				},
			)
			require.NoError(t, err)
			require.Equal(t, tc.expectedTotalShares, resp.TotalShares)
		})
	}
}

func TestQuerySpotPrice(t *testing.T) {
	tests := []struct {
		name          string
		existingPool  types.Pool
		tokenInDenom  string
		tokenOutDenom string
		expectedPrice sdk.Dec
	}{
		{
			name: "same quantity",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin("unusd", 100),
				),
				/*shares=*/ 100,
			),
			tokenInDenom:  "unusd",
			tokenOutDenom: "unibi",
			expectedPrice: sdk.MustNewDecFromStr("1"),
		},
		{
			name: "price of 2 unusd per unibi",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin("unusd", 200),
				),
				/*shares=*/ 100,
			),
			tokenInDenom:  "unusd",
			tokenOutDenom: "unibi",
			expectedPrice: sdk.MustNewDecFromStr("2"),
		},
		{
			name: "price of 0.5 unibi per unusd",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin("unusd", 200),
				),
				/*shares=*/ 100,
			),
			tokenInDenom:  "unibi",
			tokenOutDenom: "unusd",
			expectedPrice: sdk.MustNewDecFromStr("0.5"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)

			app.DexKeeper.SetPool(ctx, tc.existingPool)

			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.SpotPrice(
				sdk.WrapSDKContext(ctx),
				&types.QuerySpotPriceRequest{
					PoolId:        1,
					TokenInDenom:  tc.tokenInDenom,
					TokenOutDenom: tc.tokenOutDenom,
				},
			)
			require.NoError(t, err)
			require.Equal(t, tc.expectedPrice, sdk.MustNewDecFromStr(resp.SpotPrice))
		})
	}
}

func TestQueryEstimateSwapExactAmountIn(t *testing.T) {
	tests := []struct {
		name             string
		existingPool     types.Pool
		tokenIn          sdk.Coin
		tokenOutDenom    string
		expectedTokenOut sdk.Coin
	}{
		{
			name: "simple swap",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin("unusd", 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:          sdk.NewInt64Coin("unusd", 100),
			tokenOutDenom:    "unibi",
			expectedTokenOut: sdk.NewInt64Coin("unibi", 50),
		},
		{
			name: "complex swap",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 34844867),
					sdk.NewInt64Coin("unusd", 4684496849),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("unibi", 586848),
			tokenOutDenom: "unusd",
			// https://www.wolframalpha.com/input?i=4684496849+-+%2834844867+*+4684496849+%2F+%2834844867%2B586848%29+%29
			expectedTokenOut: sdk.NewInt64Coin("unusd", 77588330),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)
			app.DexKeeper.SetPool(ctx, tc.existingPool)
			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.EstimateSwapExactAmountIn(
				sdk.WrapSDKContext(ctx),
				&types.QuerySwapExactAmountInRequest{
					PoolId:        1,
					TokenIn:       tc.tokenIn,
					TokenOutDenom: tc.tokenOutDenom,
				},
			)

			require.NoError(t, err)
			require.Equal(t, tc.expectedTokenOut, resp.TokenOut)
		})
	}
}

func TestQueryEstimateSwapExactAmountOut(t *testing.T) {
	tests := []struct {
		name            string
		existingPool    types.Pool
		tokenOut        sdk.Coin
		tokenInDenom    string
		expectedTokenIn sdk.Coin
	}{
		{
			name: "simple swap",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin("unusd", 100),
				),
				/*shares=*/ 100,
			),
			tokenOut:     sdk.NewInt64Coin("unibi", 50),
			tokenInDenom: "unusd",
			// there's a swap fee that we take the ceiling of to round int
			expectedTokenIn: sdk.NewInt64Coin("unusd", 101),
		},
		{
			name: "complex swap",
			existingPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 34844867),
					sdk.NewInt64Coin("unusd", 4684496849),
				),
				/*shares=*/ 100,
			),
			tokenOut:     sdk.NewInt64Coin("unusd", 77588330),
			tokenInDenom: "unibi",
			// https://www.wolframalpha.com/input?i=4684496849+-+%2834844867+*+4684496849+%2F+%2834844867%2B586848%29+%29
			expectedTokenIn: sdk.NewInt64Coin("unibi", 586848),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testutil.NewNibiruApp(true)
			app.DexKeeper.SetPool(ctx, tc.existingPool)
			queryServer := keeper.NewQuerier(app.DexKeeper)

			resp, err := queryServer.EstimateSwapExactAmountOut(
				sdk.WrapSDKContext(ctx),
				&types.QuerySwapExactAmountOutRequest{
					PoolId:       1,
					TokenOut:     tc.tokenOut,
					TokenInDenom: tc.tokenInDenom,
				},
			)

			require.NoError(t, err)
			require.Equal(t, tc.expectedTokenIn, resp.TokenIn)
		})
	}
}
