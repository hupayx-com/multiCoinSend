package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend/types"
)

func (k msgServer) CoinSend(goCtx context.Context, msg *types.MsgCoinSend) (*types.MsgCoinSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	logger := k.Logger(ctx)

	totalAmount := sdk.NewCoin("stake", sdk.NewInt(0))
	receivers := msg.Receivers

	for _, receiver := range receivers {
		totalAmount.AddAmount(receiver.Amount[0].Amount)
	}

	from, _ := sdk.AccAddressFromBech32(msg.Creator)

	// 모듈로 옮겼다가
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, sdk.NewCoins(totalAmount)); err != nil {
		return nil, err
	}

	for _, receiver := range receivers {

		to, _ := sdk.AccAddressFromBech32(receiver.To)

		logger.Debug(
			"multi sender",
			"address", receiver.To,
		)

		// 개인계좌로 옮기자
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, receiver.Amount); err != nil {
			return nil, err
		}

	}

	// totalAmount.Amount
	return &types.MsgCoinSendResponse{CntValue: strconv.Itoa(len(receivers)), TotalAmount: totalAmount.Amount.String()}, nil
}
