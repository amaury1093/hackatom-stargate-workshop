package step2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/amaurymartiny/step2/x/step2/types"
	"github.com/amaurymartiny/step2/x/step2/keeper"
)

func handleMsgCreateName(ctx sdk.Context, k keeper.Keeper, name *types.MsgName) (*sdk.Result, error) {
	k.CreateName(ctx, *name)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
