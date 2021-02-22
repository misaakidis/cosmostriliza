package cosmostriliza

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/keeper"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreateGame:
			return handleMsgCreateGame(ctx, k, msg)
		case *types.MsgJoinGame:
			return handleMsgJoinGame(ctx, k, msg)
		case *types.MsgCommitMove:
			return handleMsgCommitMove(ctx, k, msg)

		/* TODO Close games (i) that have not yet started, or (ii) that are idle for over 10 days.
		   Transfer reward to (i) the host / (ii) the opponent of the player who abandoned the game.
		 	 This is to avoid locking the reward (e.g. by players who are about to lose).
		 */

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
