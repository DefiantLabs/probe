// APACHE NOTICE
// Sourced with modifications from https://github.com/strangelove-ventures/lens
package client

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	ibcTransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcTm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
)

var (
	// Provides a default set of AppModuleBasics that are included in the ChainClientConfig
	// This is used to provide a default set of modules that will be used for protobuf registration and in-app decoding of RPC responses
	DefaultModuleBasics = []module.AppModuleBasic{
		auth.AppModuleBasic{},
		authz.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		gov.AppModuleBasic{},
		crisis.AppModuleBasic{},
		distribution.AppModuleBasic{},
		feegrant.AppModuleBasic{},
		mint.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		staking.AppModuleBasic{},
		vesting.AppModuleBasic{},
		wasm.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibcTransfer.AppModuleBasic{},
		ibcTm.AppModuleBasic{},
	}
)

type ChainClientConfig struct {
	Key                   string                  `json:"key" yaml:"key"`
	ChainID               string                  `json:"chain-id" yaml:"chain-id"`
	RPCAddr               string                  `json:"rpc-addr" yaml:"rpc-addr"`
	AccountPrefix         string                  `json:"account-prefix" yaml:"account-prefix"`
	KeyringBackend        string                  `json:"keyring-backend" yaml:"keyring-backend"`
	KeyDirectory          string                  `json:"key-directory" yaml:"key-directory"`
	Debug                 bool                    `json:"debug" yaml:"debug"`
	Timeout               string                  `json:"timeout" yaml:"timeout"`
	OutputFormat          string                  `json:"output-format" yaml:"output-format"`
	Modules               []module.AppModuleBasic `json:"-" yaml:"-"`
	CustomMsgTypeRegistry map[string]sdkTypes.Msg `json:"-" yaml:"-"`
}
