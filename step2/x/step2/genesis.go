package step2

import (
	"github.com/amaurymartiny/step2/x/step2/keeper"
	"github.com/amaurymartiny/step2/x/step2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesis()
}
