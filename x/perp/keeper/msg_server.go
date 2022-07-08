package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/perp/types"
	vpooltypes "github.com/NibiruChain/nibiru/x/vpool/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (m msgServer) RemoveMargin(ctx context.Context, margin *types.MsgRemoveMargin,
) (*types.MsgRemoveMarginResponse, error) {
	return m.k.RemoveMargin(ctx, margin)
}

func (m msgServer) AddMargin(ctx context.Context, margin *types.MsgAddMargin,
) (*types.MsgAddMarginResponse, error) {
	return m.k.AddMargin(ctx, margin)
}

func (m msgServer) OpenPosition(goCtx context.Context, req *types.MsgOpenPosition,
) (response *types.MsgOpenPositionResponse, err error) {
	pair, err := common.NewAssetPair(req.TokenPair)
	if err != nil {
		panic(err) // must not happen
	}
	sender, err := sdk.AccAddressFromBech32(req.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	err = m.k.OpenPosition(
		ctx,
		pair,
		req.Side,
		sender,
		req.QuoteAssetAmount,
		req.Leverage,
		req.BaseAssetAmountLimit.ToDec(),
	)
	if err != nil {
		return nil, sdkerrors.Wrap(vpooltypes.ErrOpeningPosition, err.Error())
	}

	return &types.MsgOpenPositionResponse{}, nil
}

func (m msgServer) ClosePosition(goCtx context.Context, position *types.MsgClosePosition) (*types.MsgClosePositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, err := sdk.AccAddressFromBech32(position.Sender)
	if err != nil {
		panic(err)
	}

	tokenPair, err := common.NewAssetPair(position.TokenPair)
	if err != nil {
		panic(err)
	}

	_, err = m.k.ClosePosition(ctx, tokenPair, addr)

	return &types.MsgClosePositionResponse{}, err
}

func (m msgServer) Liquidate(goCtx context.Context, msg *types.MsgLiquidate,
) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	liquidatorAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	traderAddr, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		return nil, err
	}

	pair, err := common.NewAssetPair(msg.TokenPair)
	if err != nil {
		return nil, err
	}

	feeToLiquidator, feeToFund, err := m.k.Liquidate(ctx, liquidatorAddr, pair, traderAddr)
	if err != nil {
		return nil, err
	}

	return &types.MsgLiquidateResponse{
		FeeToLiquidator:        feeToLiquidator,
		FeeToPerpEcosystemFund: feeToFund,
	}, nil
}
