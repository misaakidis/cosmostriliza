package cosmostriliza

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/keeper"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateGame) (*sdk.Result, error) {
	k.CreateGame(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateGame) (*sdk.Result, error) {
	var game = types.Game{
		Creator:    msg.Creator,
		Id:         msg.Id,
		Rows:       msg.Rows,
		Cols:       msg.Cols,
		Strike:     msg.Strike,
		Reward:     msg.Reward,
		GameState:  msg.GameState,
		Guest:      msg.Guest,
		Oplayer:    msg.Oplayer,
		Xplayer:    msg.Xplayer,
		NextPlayer: msg.NextPlayer,
	}

	// Checks that the element exists
	if !k.HasGame(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetGameOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteGame) (*sdk.Result, error) {
	if !k.HasGame(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetGameOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeleteGame(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
