package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateToken{}, "asset/CreateToken", nil)
	cdc.RegisterConcrete(&MsgUpdateToken{}, "asset/UpdateToken", nil)
	cdc.RegisterConcrete(&MsgAuthorizeAddress{}, "asset/AuthorizeAddress", nil)
	cdc.RegisterConcrete(&MsgUnAuthorizeAddress{}, "asset/UnAuthorizeAddress", nil)
	cdc.RegisterConcrete(&MsgTransferToken{}, "asset/TransferToken", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthorizeAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnAuthorizeAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferToken{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
