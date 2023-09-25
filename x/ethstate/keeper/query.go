package keeper

import (
	"ethState/x/ethstate/types"
)

var _ types.QueryServer = Keeper{}
