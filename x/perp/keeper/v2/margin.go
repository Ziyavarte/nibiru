package keeper

import (
	"fmt"

	"github.com/NibiruChain/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/common/asset"
	"github.com/NibiruChain/nibiru/x/perp/types/v1"
	v2types "github.com/NibiruChain/nibiru/x/perp/types/v2"
)

/*
	AddMargin deleverages an existing position by adding margin (collateral)

to it. Adding margin increases the margin ratio of the corresponding position.
*/
func (k Keeper) AddMargin(
	ctx sdk.Context, pair asset.Pair, traderAddr sdk.AccAddress, marginToAdd sdk.Coin,
) (res *v2types.MsgAddMarginResponse, err error) {
	market, err := k.Markets.Get(ctx, pair)
	if err != nil {
		return nil, types.ErrPairNotFound
	}

	amm, err := k.AMMs.Get(ctx, pair)
	if err != nil {
		return nil, types.ErrPairNotFound
	}

	position, err := k.Positions.Get(ctx, collections.Join(pair, traderAddr))
	if err != nil {
		return nil, err
	}

	fundingPayment := FundingPayment(position, market.LatestCumulativePremiumFraction)
	remainingMargin := position.Margin.Sub(fundingPayment).Add(marginToAdd.Amount.ToDec())

	if remainingMargin.IsNegative() {
		return nil, fmt.Errorf("failed to add margin; resultant position would still have bad debt; consider adding more margin")
	}

	if err = k.BankKeeper.SendCoinsFromAccountToModule(
		ctx,
		/* from */ traderAddr,
		/* to */ v2types.VaultModuleAccount,
		/* amount */ sdk.NewCoins(marginToAdd),
	); err != nil {
		return nil, err
	}

	// apply funding payment and add margin
	position.Margin = remainingMargin
	position.LatestCumulativePremiumFraction = market.LatestCumulativePremiumFraction
	position.LastUpdatedBlockNumber = ctx.BlockHeight()
	k.Positions.Insert(ctx, collections.Join(position.Pair, traderAddr), position)

	positionNotional, err := PositionNotionalSpot(amm, position)
	if err != nil {
		return nil, err
	}

	if err = ctx.EventManager().EmitTypedEvent(
		&types.PositionChangedEvent{
			Pair:               pair,
			TraderAddress:      traderAddr.String(),
			Margin:             sdk.NewCoin(pair.QuoteDenom(), position.Margin.RoundInt()),
			PositionNotional:   positionNotional,
			ExchangedNotional:  sdk.ZeroDec(),                                 // always zero when adding margin
			ExchangedSize:      sdk.ZeroDec(),                                 // always zero when adding margin
			TransactionFee:     sdk.NewCoin(pair.QuoteDenom(), sdk.ZeroInt()), // always zero when adding margin
			PositionSize:       position.Size_,
			RealizedPnl:        sdk.ZeroDec(), // always zero when adding margin
			UnrealizedPnlAfter: UnrealizedPnl(position, positionNotional),
			BadDebt:            sdk.NewCoin(pair.QuoteDenom(), sdk.ZeroInt()), // always zero when adding margin
			FundingPayment:     fundingPayment,
			MarkPrice:          amm.MarkPrice(),
			BlockHeight:        ctx.BlockHeight(),
			BlockTimeMs:        ctx.BlockTime().UnixMilli(),
		},
	); err != nil {
		return nil, err
	}

	return &v2types.MsgAddMarginResponse{
		FundingPayment: fundingPayment,
		Position:       &position,
	}, nil
}

/*
	RemoveMargin further leverages an existing position by directly removing

the margin (collateral) that backs it from the vault. This also decreases the
margin ratio of the position.

Fails if the position goes underwater.

args:
  - ctx: the cosmos-sdk context
  - pair: the asset pair
  - traderAddr: the trader's address
  - margin: the amount of margin to withdraw. Must be positive.

ret:
  - marginOut: the amount of margin removed
  - fundingPayment: the funding payment that was applied with this position interaction
  - err: error if any
*/
func (k Keeper) RemoveMargin(
	ctx sdk.Context, pair asset.Pair, traderAddr sdk.AccAddress, marginToRemove sdk.Coin,
) (res *v2types.MsgRemoveMarginResponse, err error) {
	// fetch objects from state
	market, err := k.Markets.Get(ctx, pair)
	if err != nil {
		return nil, types.ErrPairNotFound
	}

	amm, err := k.AMMs.Get(ctx, pair)
	if err != nil {
		return nil, types.ErrPairNotFound
	}

	position, err := k.Positions.Get(ctx, collections.Join(pair, traderAddr))
	if err != nil {
		return nil, err
	}

	// ensure we have enough free collateral
	spotNotional, err := PositionNotionalSpot(amm, position)
	if err != nil {
		return nil, err
	}
	twapNotional, err := k.PositionNotionalTWAP(ctx, position, market.TwapLookbackWindow)
	if err != nil {
		return nil, err
	}
	minPositionNotional := sdk.MinDec(spotNotional, twapNotional)

	// account for funding payment
	fundingPayment := FundingPayment(position, market.LatestCumulativePremiumFraction)
	remainingMargin := position.Margin.Sub(fundingPayment)

	// account for negative PnL
	unrealizedPnl := UnrealizedPnl(position, minPositionNotional)
	if unrealizedPnl.IsNegative() {
		remainingMargin = remainingMargin.Add(unrealizedPnl)
	}

	if remainingMargin.LT(marginToRemove.Amount.ToDec()) {
		return nil, fmt.Errorf("not enough free collateral")
	}

	if err = k.Withdraw(ctx, market, traderAddr, marginToRemove.Amount); err != nil {
		return nil, err
	}

	// apply funding payment and remove margin
	position.Margin = position.Margin.Sub(fundingPayment).Sub(marginToRemove.Amount.ToDec())
	position.LatestCumulativePremiumFraction = market.LatestCumulativePremiumFraction
	position.LastUpdatedBlockNumber = ctx.BlockHeight()
	k.Positions.Insert(ctx, collections.Join(position.Pair, traderAddr), position)

	if err = ctx.EventManager().EmitTypedEvent(
		&types.PositionChangedEvent{
			Pair:               pair,
			TraderAddress:      traderAddr.String(),
			Margin:             sdk.NewCoin(pair.QuoteDenom(), position.Margin.RoundInt()),
			PositionNotional:   spotNotional,
			ExchangedNotional:  sdk.ZeroDec(),                                 // always zero when removing margin
			ExchangedSize:      sdk.ZeroDec(),                                 // always zero when removing margin
			TransactionFee:     sdk.NewCoin(pair.QuoteDenom(), sdk.ZeroInt()), // always zero when removing margin
			PositionSize:       position.Size_,
			RealizedPnl:        sdk.ZeroDec(), // always zero when removing margin
			UnrealizedPnlAfter: UnrealizedPnl(position, spotNotional),
			BadDebt:            sdk.NewCoin(pair.QuoteDenom(), sdk.ZeroInt()), // always zero when removing margin
			FundingPayment:     fundingPayment,
			MarkPrice:          amm.MarkPrice(),
			BlockHeight:        ctx.BlockHeight(),
			BlockTimeMs:        ctx.BlockTime().UnixMilli(),
		},
	); err != nil {
		return nil, err
	}

	return &v2types.MsgRemoveMarginResponse{
		FundingPayment: fundingPayment,
		Position:       &position,
	}, nil
}