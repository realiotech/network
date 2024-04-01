package ante_test

import (
	evmante "github.com/evmos/ethermint/app/ante"
	"github.com/evmos/ethermint/encoding"
	"github.com/evmos/ethermint/types"
	"github.com/realiotech/realio-network/app"
	"github.com/realiotech/realio-network/app/ante"
)

func (suite *AnteTestSuite) TestValidateHandlerOptions() {
	cases := []struct {
		name    string
		options ante.HandlerOptions
		expPass bool
	}{
		{
			"fail - empty options",
			ante.HandlerOptions{},
			false,
		},
		{
			"fail - empty account keeper",
			ante.HandlerOptions{
				AccountKeeper: nil,
			},
			false,
		},
		{
			"fail - empty bank keeper",
			ante.HandlerOptions{
				AccountKeeper: suite.app.AccountKeeper,
				BankKeeper:    nil,
			},
			false,
		},
		{
			"fail - empty distribution keeper",
			ante.HandlerOptions{
				AccountKeeper: suite.app.AccountKeeper,
				BankKeeper:    suite.app.BankKeeper,

				IBCKeeper: nil,
			},
			false,
		},
		{
			"fail - empty IBC keeper",
			ante.HandlerOptions{
				AccountKeeper: suite.app.AccountKeeper,
				BankKeeper:    suite.app.BankKeeper,

				IBCKeeper: nil,
			},
			false,
		},
		{
			"fail - empty staking keeper",
			ante.HandlerOptions{
				AccountKeeper: suite.app.AccountKeeper,
				BankKeeper:    suite.app.BankKeeper,

				IBCKeeper: suite.app.IBCKeeper,
			},
			false,
		},
		{
			"fail - empty fee market keeper",
			ante.HandlerOptions{
				AccountKeeper: suite.app.AccountKeeper,
				BankKeeper:    suite.app.BankKeeper,

				IBCKeeper:       suite.app.IBCKeeper,
				FeeMarketKeeper: nil,
			},
			false,
		},
		{
			"fail - empty EVM keeper",
			ante.HandlerOptions{
				AccountKeeper:   suite.app.AccountKeeper,
				BankKeeper:      suite.app.BankKeeper,
				IBCKeeper:       suite.app.IBCKeeper,
				FeeMarketKeeper: suite.app.FeeMarketKeeper,
				EvmKeeper:       nil,
			},
			false,
		},
		{
			"fail - empty signature gas consumer",
			ante.HandlerOptions{
				AccountKeeper:   suite.app.AccountKeeper,
				BankKeeper:      suite.app.BankKeeper,
				IBCKeeper:       suite.app.IBCKeeper,
				FeeMarketKeeper: suite.app.FeeMarketKeeper,
				EvmKeeper:       suite.app.EvmKeeper,
				SigGasConsumer:  nil,
			},
			false,
		},
		{
			"fail - empty signature mode handler",
			ante.HandlerOptions{
				AccountKeeper:   suite.app.AccountKeeper,
				BankKeeper:      suite.app.BankKeeper,
				IBCKeeper:       suite.app.IBCKeeper,
				FeeMarketKeeper: suite.app.FeeMarketKeeper,
				EvmKeeper:       suite.app.EvmKeeper,
				SigGasConsumer:  evmante.DefaultSigVerificationGasConsumer,
				SignModeHandler: nil,
			},
			false,
		},
		{
			"fail - empty tx fee checker",
			ante.HandlerOptions{
				AccountKeeper:   suite.app.AccountKeeper,
				BankKeeper:      suite.app.BankKeeper,
				IBCKeeper:       suite.app.IBCKeeper,
				FeeMarketKeeper: suite.app.FeeMarketKeeper,
				EvmKeeper:       suite.app.EvmKeeper,
				SigGasConsumer:  evmante.DefaultSigVerificationGasConsumer,
				SignModeHandler: suite.app.GetTxConfig().SignModeHandler(),
				TxFeeChecker:    nil,
			},
			false,
		},
		{
			"success - default app options",
			ante.HandlerOptions{
				AccountKeeper:          suite.app.AccountKeeper,
				BankKeeper:             suite.app.BankKeeper,
				ExtensionOptionChecker: types.HasDynamicFeeExtensionOption,
				EvmKeeper:              suite.app.EvmKeeper,
				FeegrantKeeper:         suite.app.FeeGrantKeeper,
				IBCKeeper:              suite.app.IBCKeeper,
				FeeMarketKeeper:        suite.app.FeeMarketKeeper,
				SignModeHandler:        encoding.MakeConfig(app.ModuleBasics).TxConfig.SignModeHandler(),
				SigGasConsumer:         evmante.DefaultSigVerificationGasConsumer,
				MaxTxGasWanted:         40000000,
				TxFeeChecker:           evmante.NewDynamicFeeChecker(suite.app.EvmKeeper),
			},
			true,
		},
	}

	for _, tc := range cases {
		err := tc.options.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}
