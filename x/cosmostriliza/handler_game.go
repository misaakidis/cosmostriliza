package cosmostriliza

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/keeper"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateGame) (*sdk.Result, error) {
	//TODO Validate parameters
	
	k.CreateGame(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
