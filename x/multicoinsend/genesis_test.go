package multicoinsend_test

import (
	"testing"

	keepertest "github.com/hupayx-com/multiCoinSend/testutil/keeper"
	"github.com/hupayx-com/multiCoinSend/testutil/nullify"
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend"
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MulticoinsendKeeper(t)
	multicoinsend.InitGenesis(ctx, *k, genesisState)
	got := multicoinsend.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
