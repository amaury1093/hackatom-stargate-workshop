package step3

import (
	"github.com/amaurymartiny/step3/x/step3/keeper"
	"github.com/amaurymartiny/step3/x/step3/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreateComment(ctx sdk.Context, k keeper.Keeper, comment *types.MsgComment) (*sdk.Result, error) {
	k.CreateComment(ctx, *comment)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
