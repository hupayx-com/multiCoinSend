package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hupayx-com/multiCoinSend/testutil/keeper"
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend/keeper"
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MulticoinsendKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
