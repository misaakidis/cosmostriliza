package cosmostriliza

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

  "github.com/misaakidis/cosmos-triliza/x/cosmostriliza/keeper"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

func handleMsgCommitMove(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCommitMove) (*sdk.Result, error) {
	// Check if gameId exists, gameState == ONGOING and is player's turn (nextPlayer == [creator/guest]Mark)

	// Check if [row, col] valid move: elem = row*col, check board[elem/row,elem%col] == EMPTY

	// Store move
	k.CommitMove(ctx, *msg)

	// If (moves > 2*strike - 1)
	// Check row: n = 1; for (i=1; i<=strike && col+-i ∈ [0,cols-1]; i++) { if (board[row][col+-i]] == mark) n++ else break; if (n >= strike) gameWin() }
	// Check col: n = 1; for (i=1; i<=strike && row+-i ∈ [0,rows-1]; i++) { if (board[row+-i][col-1] == mark) n++ else break; if (n >= strike) gameWin() }
	// Check diag: n = 1; for (i=1; i<=strike && col+-i ∈ [0,cols-1] && row+-i ∈ [0,rows-1]; i++) { if (board[row+-i][col+-i] == mark) n++ else break; if (n >= strike) gameWin() }

	// If (numOfMoves == rows*cols) { gameDraw() }

	// Set nextPlayer = opponent

	// Store updated game

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}


func gameWin() {
	// Set gameState

	// Transfer reward

	// Store updated game
}

func gameDraw() {
	// Set gameState

	// Return reward

	// Store updated game
}
