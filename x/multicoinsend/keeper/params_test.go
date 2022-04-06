package keeper_test

import (
	"testing"

	testkeeper "github.com/hupayx-com/multiCoinSend/testutil/keeper"
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MulticoinsendKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
