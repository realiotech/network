package keeper_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/realiotech/realio-network/app"
	"github.com/realiotech/realio-network/x/asset/keeper"
	"github.com/realiotech/realio-network/x/asset/types"
)

const PRIV_NAME = "test"

type KeeperTestSuite struct {
	suite.Suite
	app *app.RealioNetwork
	ctx sdk.Context

	assetKeeper *keeper.Keeper
	govkeeper   govkeeper.Keeper
	msgServer   types.MsgServer
	bankKeeper  bankkeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	app := app.Setup(false, nil, 3)

	suite.app = app
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{Height: app.LastBlockHeight() + 1})
	suite.assetKeeper = keeper.NewKeeper(
		app.AppCodec(), app.InterfaceRegistry(), app.GetKey(types.StoreKey),
		app.GetMemKey(types.MemStoreKey), app.GetSubspace(types.ModuleName), app.BankKeeper, app.AccountKeeper,
	)
	suite.govkeeper = app.GovKeeper
	suite.bankKeeper = app.BankKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func init() {
	proto.RegisterType((*MockPrivilegeMsg)(nil), "MockPrivilegeMsg")
}

// MockPrivilegeMsg defines a mock type PrivilegeMsg
type MockPrivilegeMsg struct {
	privName string
}

var _ proto.Message = &MockPrivilegeMsg{}

func (m *MockPrivilegeMsg) ValidateBasic() error {
	if m.privName == "" {
		return fmt.Errorf("empty")
	}
	return nil
}

func (m *MockPrivilegeMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{userAddr1}
}

func (m *MockPrivilegeMsg) Reset()         { *m = MockPrivilegeMsg{} }
func (m *MockPrivilegeMsg) String() string { return proto.CompactTextString(m) }
func (m *MockPrivilegeMsg) ProtoMessage()  {}

func (m *MockPrivilegeMsg) NeedPrivilege() string {
	return m.privName
}
func (m *MockPrivilegeMsg) XXX_MessageName() string {
	return "MockPrivilegeMsg"
}

func (m *MockPrivilegeMsg) XXX_Unmarshal(b []byte) error {
	return nil
}

// // MockPrivilegeI defines a mock type PrivilegeI
type MockPrivilegeI struct {
	count    uint64
	privName string
}

var _ types.PrivilegeI = MockPrivilegeI{}

func (m MockPrivilegeI) Name() string {
	return m.privName
}

func (m MockPrivilegeI) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&types.MsgMock{},
	)
}

func (m MockPrivilegeI) MsgHandler() types.MsgHandler {
	return func(context context.Context, privMsg proto.Message, tokenID string, privAcc sdk.AccAddress) (proto.Message, error) {
		typeUrl := "/" + proto.MessageName(privMsg)
		if typeUrl != sdk.MsgTypeURL(&types.MsgMock{}) {
			return nil, errors.New("unsupport msg type")
		}

		msg, ok := privMsg.(*types.MsgMock)
		if !ok {
			return nil, errors.New("unable to cast msg type")
		}

		return &types.MsgMockResponse{
			Count: msg.Count,
		}, nil
	}
}

func (m MockPrivilegeI) QueryHandler() types.QueryHandler {
	return func(context context.Context, privQuery proto.Message, tokenID string) (proto.Message, error) {
		return nil, nil
	}
}

func (m MockPrivilegeI) CLI() *cobra.Command {
	return &cobra.Command{
		Use: "mock",
	}
}
