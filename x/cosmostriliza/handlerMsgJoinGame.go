package cosmostriliza

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

  "github.com/misaakidis/cosmos-triliza/x/cosmostriliza/keeper"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

func handleMsgJoinGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgJoinGame) (*sdk.Result, error) {
  // Check if gameId exists and is OPEN

  // Check ErrInsufficientFunds for reward

  // Determine who will be X (the first player)

  // Set game fields: guest, creatorMark, guestMark

  // Set gameState to ONGOING

  // Lock reward

  // Store updated game

	k.JoinGame(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
