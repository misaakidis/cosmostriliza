package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, rows string, cols string, strike string, reward string, gameState string, guest string, Oplayer string, Xplayer string, nextPlayer string) *MsgCreateGame {
	return &MsgCreateGame{
		Creator:    creator,
		Rows:       rows,
		Cols:       cols,
		Strike:     strike,
		Reward:     reward,
		GameState:  gameState,
		Guest:      guest,
		Oplayer:    Oplayer,
		Xplayer:    Xplayer,
		NextPlayer: nextPlayer,
	}
}

func (msg *MsgCreateGame) Route() string {
	return RouterKey
}

func (msg *MsgCreateGame) Type() string {
	return "CreateGame"
}

func (msg *MsgCreateGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGame{}

func NewMsgUpdateGame(creator string, id string, rows string, cols string, strike string, reward string, gameState string, guest string, Oplayer string, Xplayer string, nextPlayer string) *MsgUpdateGame {
	return &MsgUpdateGame{
		Id:         id,
		Creator:    creator,
		Rows:       rows,
		Cols:       cols,
		Strike:     strike,
		Reward:     reward,
		GameState:  gameState,
		Guest:      guest,
		Oplayer:    Oplayer,
		Xplayer:    Xplayer,
		NextPlayer: nextPlayer,
	}
}

func (msg *MsgUpdateGame) Route() string {
	return RouterKey
}

func (msg *MsgUpdateGame) Type() string {
	return "UpdateGame"
}

func (msg *MsgUpdateGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgDeleteGame(creator string, id string) *MsgDeleteGame {
	return &MsgDeleteGame{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteGame) Route() string {
	return RouterKey
}

func (msg *MsgDeleteGame) Type() string {
	return "DeleteGame"
}

func (msg *MsgDeleteGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
