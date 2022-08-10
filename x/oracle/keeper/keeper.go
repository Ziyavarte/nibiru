package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/NibiruChain/nibiru/x/oracle/types"
)

// Keeper of the oracle store
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramstypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	distrKeeper   types.DistributionKeeper
	StakingKeeper types.StakingKeeper

	distrName string
}

// NewKeeper constructs a new keeper for oracle
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey,
	paramspace paramstypes.Subspace, accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper, distrKeeper types.DistributionKeeper,
	stakingKeeper types.StakingKeeper, distrName string) Keeper {
	// ensure oracle module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// set KeyTable if it has not already been set
	if !paramspace.HasKeyTable() {
		paramspace = paramspace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		paramSpace:    paramspace,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		distrKeeper:   distrKeeper,
		StakingKeeper: stakingKeeper,
		distrName:     distrName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

//-----------------------------------
// ExchangeRate logic

// GetExchangeRate gets the consensus exchange rate of the given pair from the store.
func (k Keeper) GetExchangeRate(ctx sdk.Context, pair string) (sdk.Dec, error) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetExchangeRateKey(pair))
	if b == nil {
		return sdk.ZeroDec(), sdkerrors.Wrap(types.ErrUnknownPair, pair)
	}

	dp := sdk.DecProto{}
	k.cdc.MustUnmarshal(b, &dp)
	return dp.Dec, nil
}

// SetExchangeRate sets the consensus exchange rate of the given pair.
func (k Keeper) SetExchangeRate(ctx sdk.Context, pair string, exchangeRate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: exchangeRate})
	store.Set(types.GetExchangeRateKey(pair), bz)
}

// SetExchangeRateWithEvent calls SetExchangeRate and emits an event.
func (k Keeper) SetExchangeRateWithEvent(ctx sdk.Context, pair string, exchangeRate sdk.Dec) {
	k.SetExchangeRate(ctx, pair, exchangeRate)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeExchangeRateUpdate,
			sdk.NewAttribute(types.AttributeKeyPair, pair),
			sdk.NewAttribute(types.AttributeKeyExchangeRate, exchangeRate.String()),
		),
	)
}

// DeleteExchangeRate deletes the consensus exchange rate of the provided pair from the store.
func (k Keeper) DeleteExchangeRate(ctx sdk.Context, pair string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetExchangeRateKey(pair))
}

// IterateExchangeRates iterates over pair's exchange rates
func (k Keeper) IterateExchangeRates(ctx sdk.Context, handler func(pair string, exchangeRate sdk.Dec) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ExchangeRateKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		rawKey := iter.Key()
		pair := string(rawKey[len(types.ExchangeRateKey) : len(rawKey)-1])
		dp := sdk.DecProto{}
		k.cdc.MustUnmarshal(iter.Value(), &dp)
		if handler(pair, dp.Dec) {
			break
		}
	}
}

//-----------------------------------
// Oracle delegation logic

// GetFeederDelegation gets the account address that the validator operator delegated oracle vote rights to
func (k Keeper) GetFeederDelegation(ctx sdk.Context, operator sdk.ValAddress) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetFeederDelegationKey(operator))
	if bz == nil {
		// By default the right is delegated to the validator itself
		return sdk.AccAddress(operator)
	}

	return sdk.AccAddress(bz)
}

// SetFeederDelegation sets the account address that the validator operator delegated oracle vote rights to
func (k Keeper) SetFeederDelegation(ctx sdk.Context, operator sdk.ValAddress, delegatedFeeder sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetFeederDelegationKey(operator), delegatedFeeder.Bytes())
}

// IterateFeederDelegations iterates over the feed delegates and performs a callback function.
func (k Keeper) IterateFeederDelegations(ctx sdk.Context,
	handler func(delegator sdk.ValAddress, delegate sdk.AccAddress) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.FeederDelegationKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		delegator := sdk.ValAddress(iter.Key()[2:])
		delegate := sdk.AccAddress(iter.Value())

		if handler(delegator, delegate) {
			break
		}
	}
}

//-----------------------------------
// Miss counter logic

// GetMissCounter retrieves the # of vote periods missed in this oracle slash window
func (k Keeper) GetMissCounter(ctx sdk.Context, operator sdk.ValAddress) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetMissCounterKey(operator))
	if bz == nil {
		// By default the counter is zero
		return 0
	}

	var missCounter gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &missCounter)
	return missCounter.Value
}

// SetMissCounter updates the # of vote periods missed in this oracle slash window
func (k Keeper) SetMissCounter(ctx sdk.Context, operator sdk.ValAddress, missCounter uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: missCounter})
	store.Set(types.GetMissCounterKey(operator), bz)
}

// DeleteMissCounter removes miss counter for the validator
func (k Keeper) DeleteMissCounter(ctx sdk.Context, operator sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMissCounterKey(operator))
}

// IterateMissCounters iterates over the miss counters and performs a callback function.
func (k Keeper) IterateMissCounters(ctx sdk.Context,
	handler func(operator sdk.ValAddress, missCounter uint64) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.MissCounterKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		operator := sdk.ValAddress(iter.Key()[2:])

		var missCounter gogotypes.UInt64Value
		k.cdc.MustUnmarshal(iter.Value(), &missCounter)

		if handler(operator, missCounter.Value) {
			break
		}
	}
}

//-----------------------------------
// AggregateExchangeRatePrevote logic

// GetAggregateExchangeRatePrevote retrieves an oracle prevote from the store
func (k Keeper) GetAggregateExchangeRatePrevote(ctx sdk.Context, voter sdk.ValAddress) (aggregatePrevote types.AggregateExchangeRatePrevote, err error) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetAggregateExchangeRatePrevoteKey(voter))
	if b == nil {
		err = sdkerrors.Wrap(types.ErrNoAggregatePrevote, voter.String())
		return
	}
	k.cdc.MustUnmarshal(b, &aggregatePrevote)
	return
}

// SetAggregateExchangeRatePrevote set an oracle aggregate prevote to the store
func (k Keeper) SetAggregateExchangeRatePrevote(ctx sdk.Context, voter sdk.ValAddress, prevote types.AggregateExchangeRatePrevote) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&prevote)

	store.Set(types.GetAggregateExchangeRatePrevoteKey(voter), bz)
}

// DeleteAggregateExchangeRatePrevote deletes an oracle prevote from the store
func (k Keeper) DeleteAggregateExchangeRatePrevote(ctx sdk.Context, voter sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAggregateExchangeRatePrevoteKey(voter))
}

// IterateAggregateExchangeRatePrevotes iterates rate over prevotes in the store
func (k Keeper) IterateAggregateExchangeRatePrevotes(ctx sdk.Context, handler func(voterAddr sdk.ValAddress, aggregatePrevote types.AggregateExchangeRatePrevote) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.AggregateExchangeRatePrevoteKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		voterAddr := sdk.ValAddress(iter.Key()[2:])

		var aggregatePrevote types.AggregateExchangeRatePrevote
		k.cdc.MustUnmarshal(iter.Value(), &aggregatePrevote)
		if handler(voterAddr, aggregatePrevote) {
			break
		}
	}
}

//-----------------------------------
// AggregateExchangeRateVote logic

// GetAggregateExchangeRateVote retrieves an oracle prevote from the store
func (k Keeper) GetAggregateExchangeRateVote(ctx sdk.Context, voter sdk.ValAddress) (aggregateVote types.AggregateExchangeRateVote, err error) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetAggregateExchangeRateVoteKey(voter))
	if b == nil {
		err = sdkerrors.Wrap(types.ErrNoAggregateVote, voter.String())
		return
	}
	k.cdc.MustUnmarshal(b, &aggregateVote)
	return
}

// SetAggregateExchangeRateVote adds an oracle aggregate prevote to the store
func (k Keeper) SetAggregateExchangeRateVote(ctx sdk.Context, voter sdk.ValAddress, vote types.AggregateExchangeRateVote) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&vote)
	store.Set(types.GetAggregateExchangeRateVoteKey(voter), bz)
}

// DeleteAggregateExchangeRateVote deletes an oracle prevote from the store
func (k Keeper) DeleteAggregateExchangeRateVote(ctx sdk.Context, voter sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAggregateExchangeRateVoteKey(voter))
}

// IterateAggregateExchangeRateVotes iterates rate over prevotes in the store
func (k Keeper) IterateAggregateExchangeRateVotes(ctx sdk.Context, handler func(voterAddr sdk.ValAddress, aggregateVote types.AggregateExchangeRateVote) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.AggregateExchangeRateVoteKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		voterAddr := sdk.ValAddress(iter.Key()[2:])

		var aggregateVote types.AggregateExchangeRateVote
		k.cdc.MustUnmarshal(iter.Value(), &aggregateVote)
		if handler(voterAddr, aggregateVote) {
			break
		}
	}
}

// PairExists return tobin tax for the pair
// TODO(mercilex): use AssetPair
func (k Keeper) PairExists(ctx sdk.Context, pair string) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPairKey(pair))
	if bz == nil {
		return false
	}

	return true
}

// SetPair updates tobin tax for the pair
// TODO(mercilex): use AssetPair
func (k Keeper) SetPair(ctx sdk.Context, pair string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPairKey(pair), []byte{})
}

// IteratePairs iterates rate over tobin taxes in the store
func (k Keeper) IteratePairs(ctx sdk.Context, handler func(pair string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PairsKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		pair := types.ExtractPairFromPairKey(iter.Key())
		if handler(pair) {
			break
		}
	}
}

// ClearPairs clears tobin taxes
func (k Keeper) ClearPairs(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PairsKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// ValidateFeeder return the given feeder is allowed to feed the message or not
func (k Keeper) ValidateFeeder(ctx sdk.Context, feederAddr sdk.AccAddress, validatorAddr sdk.ValAddress) error {
	if !feederAddr.Equals(validatorAddr) {
		delegate := k.GetFeederDelegation(ctx, validatorAddr)
		if !delegate.Equals(feederAddr) {
			return sdkerrors.Wrap(types.ErrNoVotingPermission, feederAddr.String())
		}
	}

	// Check that the given validator exists
	if val := k.StakingKeeper.Validator(ctx, validatorAddr); val == nil || !val.IsBonded() {
		return sdkerrors.Wrapf(stakingtypes.ErrNoValidatorFound, "validator %s is not active set", validatorAddr.String())
	}

	return nil
}
