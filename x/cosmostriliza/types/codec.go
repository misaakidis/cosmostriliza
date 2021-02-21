package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateGame{}, "cosmostriliza/CreateGame", nil)
	cdc.RegisterConcrete(&MsgUpdateGame{}, "cosmostriliza/UpdateGame", nil)
	cdc.RegisterConcrete(&MsgDeleteGame{}, "cosmostriliza/DeleteGame", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGame{},
		&MsgUpdateGame{},
		&MsgDeleteGame{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
