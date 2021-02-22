package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCommitMove{}

func NewMsgCommitMove(player string, id string, row uint32, col uint32) *MsgCommitMove {
	return &MsgCommitMove{
		Player:			player,
		GameId:			id,
		Row:				row,
		Col:				col,
	}
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
