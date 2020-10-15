package step2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/amaurymartiny/step2/x/step2/types"
	"github.com/amaurymartiny/step2/x/step2/keeper"
)

func handleMsgCreateComment(ctx sdk.Context, k keeper.Keeper, comment *types.MsgComment) (*sdk.Result, error) {
	k.CreateComment(ctx, *comment)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
