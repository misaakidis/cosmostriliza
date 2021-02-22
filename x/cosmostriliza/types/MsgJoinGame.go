package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgJoinGame{}

func NewMsgJoinGame(guest string, id string) *MsgJoinGame {
	return &MsgJoinGame{
		Guest: 	guest,
		GameId:	id,
	}
}

func (msg *MsgJoinGame) Route() string {
	return RouterKey
}

func (msg *MsgJoinGame) Type() string {
	return "JoinGame"
}

func (msg *MsgJoinGame) GetSigners() []sdk.AccAddress {
	guest, err := sdk.AccAddressFromBech32(msg.Guest)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{guest}
}

func (msg *MsgJoinGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJoinGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Guest)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Check gameId >= 0

	return nil
}
