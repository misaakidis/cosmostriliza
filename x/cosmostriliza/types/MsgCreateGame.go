package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, rows uint32, cols uint32, strike uint32, reward uint32) *MsgCreateGame {
	// Check rows > 3
	// Check cols > 3
	// Check strike > 3 && (<=rows || <= cols)

	// Lock reward

	return &MsgCreateGame{
		Creator:    creator,
		Rows:       rows,
		Cols:       cols,
		Strike:     strike,
		Reward:     reward,
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

	// Check ErrInsufficientFunds

	return nil
}
