package keeper

import (
	"github.com/amaurymartiny/step1/x/blog/types"
)

var _ types.QueryServer = Keeper{}
