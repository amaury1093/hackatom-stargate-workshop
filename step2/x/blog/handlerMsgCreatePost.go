package blog

import (
	"github.com/amaurymartiny/step2/x/blog/keeper"
	"github.com/amaurymartiny/step2/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, post *types.MsgPost) (*sdk.Result, error) {
	k.CreatePost(ctx, *post)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
