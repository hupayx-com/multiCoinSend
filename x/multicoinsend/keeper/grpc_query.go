package keeper

import (
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend/types"
)

var _ types.QueryServer = Keeper{}
