package step3

import (
	"github.com/amaurymartiny/step3/x/step3/keeper"
	"github.com/amaurymartiny/step3/x/step3/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, post *types.MsgPost) (*sdk.Result, error) {
	k.CreatePost(ctx, *post)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
