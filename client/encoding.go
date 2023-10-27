// APACHE NOTICE
// Sourced with modifications from https://github.com/strangelove-ventures/lens
package client

import (
	osmosisGammTypes "github.com/DefiantLabs/probe/client/codec/osmosis/v15/x/gamm/types"
	osmosisLockupTypes "github.com/DefiantLabs/probe/client/codec/osmosis/v15/x/lockup/types"
	osmosisPoolManagerTypes "github.com/DefiantLabs/probe/client/codec/osmosis/v15/x/poolmanager/types"
	tendermintLiquidityTypes "github.com/DefiantLabs/probe/client/codec/tendermint/liquidity/x/liquidity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type Codec struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeCodec(moduleBasics []module.AppModuleBasic) Codec {
	modBasic := module.NewBasicManager(moduleBasics...)
	encodingConfig := MakeCodecConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	modBasic.RegisterLegacyAminoCodec(encodingConfig.Amino)
	modBasic.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}

// Split out from base codec to not include explicitly.
// Should be included only when needed.
func RegisterOsmosisInterfaces(registry types.InterfaceRegistry) {
	// Needs to be extended in order to cover all the modules
	osmosisGammTypes.RegisterInterfaces(registry)
	osmosisPoolManagerTypes.RegisterInterfaces(registry)
	osmosisLockupTypes.RegisterInterfaces(registry)
}

// Split out from base codec to not include explicitly.
// Should be included only when needed.
func RegisterTendermintLiquidityInterfaces(aminoCodec *codec.LegacyAmino, registry types.InterfaceRegistry) {
	tendermintLiquidityTypes.RegisterLegacyAminoCodec(aminoCodec)
	tendermintLiquidityTypes.RegisterInterfaces(registry)
}

func MakeCodecConfig() Codec {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return Codec{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}
}
