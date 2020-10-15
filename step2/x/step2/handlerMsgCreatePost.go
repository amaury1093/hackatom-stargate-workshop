package step2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/amaurymartiny/step2/x/step2/types"
	"github.com/amaurymartiny/step2/x/step2/keeper"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, post *types.MsgPost) (*sdk.Result, error) {
	k.CreatePost(ctx, *post)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
