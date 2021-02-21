package cosmostriliza

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/keeper"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the game
	for _, elem := range genState.GameList {
		k.SetGame(ctx, *elem)
	}

	// Set game count
	k.SetGameCount(ctx, int64(len(genState.GameList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all game
	gameList := k.GetAllGame(ctx)
	for _, elem := range gameList {
		elem := elem
		genesis.GameList = append(genesis.GameList, &elem)
	}

	return genesis
}
