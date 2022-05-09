package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/vamm/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func NewKeeper(codec codec.BinaryCodec, storeKey sdk.StoreKey, memKey sdk.StoreKey, ps paramtypes.Subspace) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		codec:      codec,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

type Keeper struct {
	codec      codec.BinaryCodec
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	paramstore paramtypes.Subspace
}

func (k Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

// SwapInput swaps pair token
func (k Keeper) SwapInput(
	ctx sdk.Context,
	pair string,
	dir types.Direction,
	quoteAssetAmount sdk.Int,
	baseAmountLimit sdk.Int,
) (sdk.Int, error) {
	if !k.existsPool(ctx, pair) {
		return sdk.Int{}, types.ErrPairNotSupported
	}

	if quoteAssetAmount.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), nil
	}

	pool, err := k.getPool(ctx, pair)
	if err != nil {
		return sdk.Int{}, err
	}

	if dir == types.Direction_REMOVE_FROM_AMM {
		enoughReserve, err := pool.HasEnoughQuoteReserve(quoteAssetAmount)
		if err != nil {
			return sdk.Int{}, err
		}
		if !enoughReserve {
			return sdk.Int{}, types.ErrOvertradingLimit
		}
	}

	baseAssetAmount, err := pool.GetBaseAmountByQuoteAmount(dir, quoteAssetAmount)
	if err != nil {
		return sdk.Int{}, err
	}

	if !baseAmountLimit.Equal(sdk.ZeroInt()) {
		if dir == types.Direction_ADD_TO_AMM {
			if baseAssetAmount.LT(baseAmountLimit) {
				return sdk.Int{}, fmt.Errorf(
					"base amount (%s) is less than selected limit (%s)",
					baseAssetAmount.String(),
					baseAmountLimit.String(),
				)
			}
		} else {
			if baseAssetAmount.GT(baseAmountLimit) {
				return sdk.Int{}, fmt.Errorf(
					"base amount (%s) is greater than selected limit (%s)",
					baseAssetAmount.String(),
					baseAmountLimit.String(),
				)
			}
		}
	}

	err = k.updateReserve(ctx, pool, dir, quoteAssetAmount, baseAssetAmount, false)
	if err != nil {
		return sdk.Int{}, fmt.Errorf("error updating reserve: %w", err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSwapInput,
			sdk.NewAttribute(types.AttributeToken0Amount, quoteAssetAmount.String()),
			sdk.NewAttribute(types.AttributeToken1Amount, baseAssetAmount.String()),
		),
	)

	return baseAssetAmount, nil
}

// getPool returns the pool from database
func (k Keeper) getPool(ctx sdk.Context, pair string) (*types.Pool, error) {
	store := k.getStore(ctx)

	bz := store.Get(types.GetPoolKey(pair))
	var pool types.Pool

	err := k.codec.Unmarshal(bz, &pool)
	if err != nil {
		return nil, err
	}

	return &pool, nil
}

// CreatePool creates a pool for a specific pair.
func (k Keeper) CreatePool(
	ctx sdk.Context,
	pair string,
	tradeLimitRatio sdk.Dec, // integer with 6 decimals, 1_000_000 means 1.0
	quoteAssetReserve sdk.Int,
	baseAssetReserve sdk.Int,
	fluctuationLimitRation sdk.Dec,
) error {
	pool := types.NewPool(pair, tradeLimitRatio, quoteAssetReserve, baseAssetReserve, fluctuationLimitRation)

	err := k.savePool(ctx, pool)
	if err != nil {
		return err
	}

	err = k.saveReserveSnapshot(ctx, 0, pool)
	if err != nil {
		return fmt.Errorf("error saving snapshot on pool creation: %w", err)
	}

	return nil
}

func (k Keeper) savePool(
	ctx sdk.Context,
	pool *types.Pool,
) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.codec.Marshal(pool)
	if err != nil {
		return err
	}

	store.Set(types.GetPoolKey(pool.Pair), bz)

	return nil
}

func (k Keeper) updateReserve(
	ctx sdk.Context,
	pool *types.Pool,
	dir types.Direction,
	quoteAssetAmount sdk.Int,
	baseAssetAmount sdk.Int,
	skipFluctuationCheck bool,
) error {
	if dir == types.Direction_ADD_TO_AMM {
		pool.IncreaseToken0Reserve(quoteAssetAmount)
		pool.DecreaseToken1Reserve(baseAssetAmount)
		// TODO baseAssetDeltaThisFunding
		// TODO totalPositionSize
		// TODO cumulativeNotional
	} else {
		pool.DecreaseToken0Reserve(quoteAssetAmount)
		pool.IncreaseToken1Reserve(baseAssetAmount)
		// TODO baseAssetDeltaThisFunding
		// TODO totalPositionSize
		// TODO cumulativeNotional
	}

	// Check if its over Fluctuation Limit Ratio.
	if !skipFluctuationCheck {
		err := k.checkFluctuationLimitRatio(ctx, pool)
		if err != nil {
			return err
		}
	}

	err := k.addReserveSnapshot(ctx, pool)
	if err != nil {
		return fmt.Errorf("error creating snapshot: %w", err)
	}

	return k.savePool(ctx, pool)
}

// existsPool returns true if pool exists, false if not.
func (k Keeper) existsPool(ctx sdk.Context, pair string) bool {
	store := k.getStore(ctx)
	return store.Has(types.GetPoolKey(pair))
}

func (k Keeper) checkFluctuationLimitRatio(ctx sdk.Context, pool *types.Pool) error {
	fluctuationLimitRatio, err := sdk.NewDecFromStr(pool.FluctuationLimitRatio)
	if err != nil {
		return fmt.Errorf("error getting fluctuation limit ratio for pool: %s", pool.Pair)
	}

	if fluctuationLimitRatio.GT(sdk.ZeroDec()) {
		latestSnapshot, counter, err := k.getLastReserveSnapshot(ctx, pool.Pair)
		if err != nil {
			return fmt.Errorf("error getting last snapshot number for pair %s", pool.Pair)
		}

		if latestSnapshot.BlockNumber == ctx.BlockHeight() && counter > 1 {
			latestSnapshot, err = k.getSnapshotByCounter(ctx, pool.Pair, counter-1)
			if err != nil {
				return fmt.Errorf("error getting snapshot number %d from pair %s", counter-1, pool.Pair)
			}
		}

		if isOverFluctuationLimit(pool, latestSnapshot) {
			return types.ErrOverFluctuationLimit
		}
	}

	return nil
}

func isOverFluctuationLimit(pool *types.Pool, snapshot types.ReserveSnapshot) bool {
	fluctuationLimitRatio, _ := sdk.NewDecFromStr(pool.FluctuationLimitRatio)
	quoteAssetReserve, _ := pool.GetPoolToken0ReserveAsInt()
	baseAssetReserve, _ := pool.GetPoolToken1ReserveAsInt()
	price := quoteAssetReserve.ToDec().Quo(baseAssetReserve.ToDec())

	snapshotQuote, _ := sdk.NewDecFromStr(snapshot.Token0Reserve)
	snapshotBase, _ := sdk.NewDecFromStr(snapshot.Token1Reserve)
	lastPrice := snapshotQuote.Quo(snapshotBase)
	upperLimit := lastPrice.Mul(sdk.OneDec().Add(fluctuationLimitRatio))
	lowerLimit := lastPrice.Mul(sdk.OneDec().Sub(fluctuationLimitRatio))

	if price.GT(upperLimit) || price.LT(lowerLimit) {
		return true
	}

	return false
}

/* CalcFee calculates the total tx fee for exchanging 'quoteAmt' of tokens on
the exchange.

Args:
	quoteAmt (sdk.Int):

Returns:
	toll (sdk.Int): Amount of tokens transferred to the the fee pool.
	spread (sdk.Int): Amount of tokens transferred to the PerpEF.
*/
func (k Keeper) CalcFee(ctx sdk.Context, quoteAmt sdk.Int) (toll sdk.Int, spread sdk.Int, err error) {
	if quoteAmt.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt(), nil
	}

	params := k.GetParams(ctx)

	tollRatio := params.GetTollRatioAsDec()
	spreadRatio := params.GetSpreadRatioAsDec()

	return sdk.NewDecFromInt(quoteAmt).Mul(tollRatio).TruncateInt(), sdk.NewDecFromInt(quoteAmt).Mul(spreadRatio).TruncateInt(), nil
}
