package keeper

import (
	"github.com/amaurymartiny/step2/x/blog/types"
)

var _ types.QueryServer = Keeper{}
