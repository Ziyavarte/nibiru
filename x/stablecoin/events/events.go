package events

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// x/stablecoin attributes for events
const (
	AttributeFromAddr    = "from"
	AttributeToAddr      = "to"
	AttributeTokenAmount = "amount"
	AttributeTokenDenom  = "denom"
)

func EmitTransfer(
	ctx sdk.Context, coin sdk.Coin, from string, to string,
) {
	const EventTypeTransfer = "transfer"
	event := sdk.NewEvent(
		EventTypeTransfer,
		sdk.NewAttribute(AttributeFromAddr, from),
		sdk.NewAttribute(AttributeToAddr, to),
		sdk.NewAttribute(AttributeTokenDenom, coin.Denom),
		sdk.NewAttribute(AttributeTokenAmount, coin.Amount.String()),
	)
	ctx.EventManager().EmitEvents(sdk.Events{event})
}

func _mintOrBurnEvent(eventType string, coin sdk.Coin) sdk.Event {
	event := sdk.NewEvent(
		eventType,
		sdk.NewAttribute(AttributeTokenDenom, coin.Denom),
		sdk.NewAttribute(AttributeTokenAmount, coin.Amount.String()),
	)
	return event
}

// EmitBurnMtrx emits an event when a Matrix Stablecoin is minted.
func EmitMintStable(ctx sdk.Context, coin sdk.Coin) {
	const EventTypeMintStable = "mint_stable"
	event := _mintOrBurnEvent(EventTypeMintStable, coin)
	ctx.EventManager().EmitEvents(sdk.Events{event})
}

// EmitBurnMtrx emits an event when a Matrix Stablecoin is burned.
func EmitBurnStable(ctx sdk.Context, coin sdk.Coin) {
	const EventTypeBurnStable = "burn_stable"
	event := _mintOrBurnEvent(EventTypeBurnStable, coin)
	ctx.EventManager().EmitEvents(sdk.Events{event})
}

// EmitBurnMtrx emits an event when MTRX is minted.
func EmitMintMtrx(ctx sdk.Context, coin sdk.Coin) {
	const EventTypeMintMtrx = "mint_mtrx"
	event := _mintOrBurnEvent(EventTypeMintMtrx, coin)
	ctx.EventManager().EmitEvents(sdk.Events{event})
}

// EmitBurnMtrx emits an event when MTRX is burned.
func EmitBurnMtrx(ctx sdk.Context, coin sdk.Coin) {
	const EventTypeBurnMtrx = "burn_mtrx"
	event := _mintOrBurnEvent(EventTypeBurnMtrx, coin)
	ctx.EventManager().EmitEvents(sdk.Events{event})
}
