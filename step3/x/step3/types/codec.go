package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(MsgComment{}, "step3/CreateComment", nil)
	cdc.RegisterConcrete(MsgPost{}, "step3/CreatePost", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgComment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPost{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)