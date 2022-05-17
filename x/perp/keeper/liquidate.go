package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/perp/events"
	"github.com/NibiruChain/nibiru/x/perp/types"
	vtypes "github.com/NibiruChain/nibiru/x/vpool/types"
)

type LiquidationOutput struct {
	FeeToPerpEcosystemFund sdk.Dec
	BadDebt                sdk.Dec
	FeeToLiquidator        sdk.Dec
	PositionResp           *types.PositionResp
}

func (l *LiquidationOutput) String() string {
	return fmt.Sprintf(`
	liquidationOutput {
		FeeToPerpEcosystemFund: %v,
		BadDebt: %v,
		FeeToLiquidator: %v,
		PositionResp: %v,
	}
	`,
		l.FeeToPerpEcosystemFund,
		l.BadDebt,
		l.FeeToLiquidator,
		l.PositionResp,
	)
}

func (l *LiquidationOutput) Validate() error {
	for _, field := range []sdk.Dec{
		l.FeeToPerpEcosystemFund, l.BadDebt, l.FeeToLiquidator} {
		if field.IsNil() {
			return fmt.Errorf(
				`invalid liquidationOutput: %v,
				must not have nil fields`, l.String())
		}
	}
	return nil
}

/* Liquidate allows to liquidate the trader position if the margin is below the
required margin maintenance ratio.
*/
func (k Keeper) Liquidate(
	goCtx context.Context, msg *types.MsgLiquidate,
) (res *types.MsgLiquidateResponse, err error) {
	// ------------- Liquidation Message Setup -------------

	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate liquidator (msg.Sender)
	liquidator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return res, err
	}

	// validate trader (msg.PositionOwner)
	trader, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		return res, err
	}

	// validate pair
	pair, err := common.NewTokenPairFromStr(msg.TokenPair)
	if err != nil {
		return res, err
	}
	err = k.requireVpool(ctx, pair)
	if err != nil {
		return res, err
	}

	position, err := k.GetPosition(ctx, pair, trader.String())
	if err != nil {
		return res, err
	}

	marginRatio, err := k.GetMarginRatio(ctx, *position, types.MarginCalculationPriceOption_MAX_PNL)
	if err != nil {
		return res, err
	}

	if k.VpoolKeeper.IsOverSpreadLimit(ctx, pair) {
		marginRatioBasedOnOracle, err := k.GetMarginRatio(
			ctx, *position, types.MarginCalculationPriceOption_INDEX)
		if err != nil {
			return res, err
		}

		marginRatio = sdk.MaxDec(marginRatio, marginRatioBasedOnOracle)
	}

	params := k.GetParams(ctx)
	err = requireMoreMarginRatio(marginRatio, params.MaintenanceMarginRatio, false)
	if err != nil {
		return res, types.MarginHighEnough
	}

	marginRatioBasedOnSpot, err := k.GetMarginRatio(
		ctx, *position, types.MarginCalculationPriceOption_SPOT)
	if err != nil {
		return res, err
	}

	fmt.Println("marginRatioBasedOnSpot", marginRatioBasedOnSpot)

	var (
		liquidationOutput LiquidationOutput
	)

	if marginRatioBasedOnSpot.GTE(params.GetPartialLiquidationRatioAsDec()) {
		liquidationOutput, err = k.CreatePartialLiquidation(ctx, pair, trader, position)
		if err != nil {
			return res, err
		}
	} else {
		liquidationOutput, err = k.CreateLiquidation(ctx, pair, trader, position)
		if err != nil {
			return res, err
		}
	}

	// Transfer fee from vault to PerpEF
	feeToPerpEF := liquidationOutput.FeeToPerpEcosystemFund.TruncateInt()
	if feeToPerpEF.IsPositive() {
		coinToPerpEF := sdk.NewCoin(
			pair.GetQuoteTokenDenom(), feeToPerpEF)
		err = k.BankKeeper.SendCoinsFromModuleToModule(
			ctx,
			types.VaultModuleAccount,
			types.PerpEFModuleAccount,
			sdk.NewCoins(coinToPerpEF),
		)
		if err != nil {
			return res, err
		}
		// TODO: emit transfer event from vault to PerpEF
	}

	// Transfer fee from PerpEF to liquidator
	feeToLiquidator := liquidationOutput.FeeToLiquidator.TruncateInt()
	if feeToLiquidator.IsPositive() {
		coinToLiquidator := sdk.NewCoin(
			pair.GetQuoteTokenDenom(), liquidationOutput.FeeToLiquidator.TruncateInt())
		err = k.BankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.PerpEFModuleAccount,
			liquidator,
			sdk.NewCoins(coinToLiquidator),
		)
		if err != nil {
			return res, err
		}
		// TODO: emit transfer event from PerpEF to liquidator
	}

	events.EmitPositionLiquidate(
		/* ctx */ ctx,
		/* vpool */ pair.String(),
		/* owner */ trader,
		/* notional */ liquidationOutput.PositionResp.ExchangedQuoteAssetAmount,
		/* vsize */ liquidationOutput.PositionResp.ExchangedPositionSize,
		/* liquidator */ liquidator,
		/* liquidationFee */ liquidationOutput.FeeToLiquidator.TruncateInt(),
		/* badDebt */ liquidationOutput.BadDebt,
	)

	return res, nil
}

//CreateLiquidation create a liquidation of a position and compute the fee to ecosystem fund
func (k Keeper) CreateLiquidation(
	ctx sdk.Context, pair common.TokenPair, owner sdk.AccAddress, position *types.Position,
) (LiquidationOutput, error) {
	params := k.GetParams(ctx)

	positionResp, err := k.closePositionEntirely(ctx, *position, sdk.ZeroDec())
	if err != nil {
		return LiquidationOutput{}, err
	}

	remainMargin := positionResp.MarginToVault.Abs()

	feeToLiquidator := positionResp.ExchangedQuoteAssetAmount.
		Mul(params.GetLiquidationFeeAsDec()).
		Quo(sdk.MustNewDecFromStr("2"))
	totalBadDebt := positionResp.BadDebt

	if feeToLiquidator.GT(remainMargin) {
		// if the remainMargin is not enough for liquidationFee, count it as bad debt
		liquidationBadDebt := feeToLiquidator.Sub(remainMargin)
		totalBadDebt = totalBadDebt.Add(liquidationBadDebt)
	} else {
		// Otherwise, the remaining margin rest will be transferred to ecosystemFund
		remainMargin = remainMargin.Sub(feeToLiquidator)
	}

	var feeToPerpEcosystemFund sdk.Dec
	if remainMargin.GT(sdk.ZeroDec()) {
		feeToPerpEcosystemFund = remainMargin
	} else {
		feeToPerpEcosystemFund = sdk.ZeroDec()
	}

	output := LiquidationOutput{
		FeeToPerpEcosystemFund: feeToPerpEcosystemFund,
		BadDebt:                totalBadDebt,
		FeeToLiquidator:        feeToLiquidator,
		PositionResp:           positionResp,
	}

	err = output.Validate()
	if err != nil {
		return LiquidationOutput{}, err
	}
	return output, err
}

// CreatePartialLiquidation create a partial liquidation of a position and compute the fee to ecosystem fund
func (k Keeper) CreatePartialLiquidation(
	ctx sdk.Context, pair common.TokenPair, trader sdk.AccAddress, position *types.Position,
) (liquidationOutput LiquidationOutput, err error) {
	params := k.GetParams(ctx)
	var (
		dir vtypes.Direction
	)

	if position.Size_.GTE(sdk.ZeroDec()) {
		dir = vtypes.Direction_ADD_TO_POOL
	} else {
		dir = vtypes.Direction_REMOVE_FROM_POOL
	}

	partiallyLiquidatedPositionNotional, err := k.VpoolKeeper.GetBaseAssetPrice(
		ctx,
		pair,
		dir,
		/*abs= */ position.Size_.Mul(params.GetPartialLiquidationRatioAsDec()).Abs(),
	)
	if err != nil {
		return
	}

	positionResp, err := k.openReversePosition(
		/* ctx */ ctx,
		/* currentPosition */ *position,
		/* quoteAssetAmount */ partiallyLiquidatedPositionNotional,
		/* leverage */ sdk.OneDec(),
		/* baseAssetAmountLimit */ sdk.ZeroDec(),
		/* canOverFluctuationLimit */ true,
	)
	if err != nil {
		return
	}

	// half of the liquidationFee goes to liquidator & another half goes to ecosystem fund
	liquidationPenalty := positionResp.ExchangedQuoteAssetAmount.Mul(params.GetLiquidationFeeAsDec())
	feeToLiquidator := liquidationPenalty.Quo(sdk.MustNewDecFromStr("2"))

	positionResp.Position.Margin = positionResp.Position.Margin.Sub(liquidationPenalty)
	k.SetPosition(ctx, pair, trader.String(), positionResp.Position)

	liquidationOutput.PositionResp = positionResp

	liquidationOutput.FeeToPerpEcosystemFund = liquidationPenalty.Sub(feeToLiquidator)
	liquidationOutput.FeeToLiquidator = feeToLiquidator

	return
}
