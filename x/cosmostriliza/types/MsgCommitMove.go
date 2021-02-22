package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCommitMove{}

func NewMsgCommitMove(player string, id string, row uint32, col uint32) *MsgCommitMove {

	// Check if gameId exists, gameState == ONGOING and is player's turn (nextPlayer == [creator/guest]Mark)

	// Check if [row, col] valid move: elem = row*col, check board[elem/row,elem%col] == EMPTY

	// If (moves > 2*strike - 1)
	// Check row: n = 1; for (i=1; i<=strike; i++) { if (board[row][col+/-i ∈ [0,cols-1]] == mark) n++ else break; if (n >= strike) WinGame() }
	// Check col: n = 1; for (i=1; i<=strike; i++) { if (board[row+/-i ∈ [0,rows-1]][col-1] == mark) n++ else break; if (n >= strike) WinGame() }
	// Check diag: n = 1; for (i=1; i<=strike; i++) { if (board[row&col+/-1 ∈ [0,rows-1][0,cols-1]] == mark) n++ else break; if (n >= strike) WinGame() }

	// If (numOfMoves == rows*cols) { DrawGame() }

	// Set nextPlayer = opponent

	// Store updated game

	return &MsgCommitMove{
		Player:			player,
		GameId:			id,
		Row:				row,
		Col:				col,
	}
}

func WinGame() {
	// Set gameState

	// Transfer reward

	// Store updated game

}

func DrawGame() {
	// Set gameState

	// Return reward

	// Store updated game
}

func (msg *MsgCommitMove) Route() string {
	return RouterKey
}

func (msg *MsgCommitMove) Type() string {
	return "CommitMove"
}

func (msg *MsgCommitMove) GetSigners() []sdk.AccAddress {
	player, err := sdk.AccAddressFromBech32(msg.Player)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{player}
}

func (msg *MsgCommitMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCommitMove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Player)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
