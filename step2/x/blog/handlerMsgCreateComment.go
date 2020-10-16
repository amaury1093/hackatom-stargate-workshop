package blog

import (
	"github.com/amaurymartiny/step2/x/blog/keeper"
	"github.com/amaurymartiny/step2/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreateComment(ctx sdk.Context, k keeper.Keeper, comment *types.MsgComment) (*sdk.Result, error) {
	k.CreateComment(ctx, *comment)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
