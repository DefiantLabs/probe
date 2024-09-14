// APACHE NOTICE
// Sourced with modifications from https://github.com/strangelove-ventures/lens
package client

import (
	"fmt"

	osmosisGammTypes "github.com/DefiantLabs/probe/client/codec/osmosis/v15/x/gamm/types"
	osmosisLockupTypes "github.com/DefiantLabs/probe/client/codec/osmosis/v15/x/lockup/types"
	osmosisPoolManagerTypes "github.com/DefiantLabs/probe/client/codec/osmosis/v15/x/poolmanager/types"
	tendermintLiquidityTypes "github.com/DefiantLabs/probe/client/codec/tendermint/liquidity/x/liquidity/types"
	probeCodecTypes "github.com/DefiantLabs/probe/client/codec/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type Codec struct {
	ProbeInterfaceRegistry *probeCodecTypes.ProbeInterfaceRegistry
	InterfaceRegistry      types.InterfaceRegistry
	Marshaler              codec.Codec
	TxConfig               client.TxConfig
	Amino                  *codec.LegacyAmino
}

func MakeCodec(moduleBasics []module.AppModuleBasic, customMsgTypeRegistry map[string]sdkTypes.Msg) (Codec, error) {
	modBasic := module.NewBasicManager(moduleBasics...)
	encodingConfig := MakeCodecConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	modBasic.RegisterLegacyAminoCodec(encodingConfig.Amino)
	modBasic.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	for typeURL, msg := range customMsgTypeRegistry {
		if encodingConfig.ProbeInterfaceRegistry.TypeURLIsRegistered(typeURL) {
			return Codec{}, fmt.Errorf("error registering custom message type in codec, typeURL %s is already registered", typeURL)
		}
		encodingConfig.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdkTypes.Msg)(nil), typeURL, msg)
	}

	return encodingConfig, nil
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
	cosmosInterfaceRegistry, probeRegistry := probeCodecTypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(cosmosInterfaceRegistry)
	return Codec{
		ProbeInterfaceRegistry: probeRegistry,
		InterfaceRegistry:      cosmosInterfaceRegistry,
		Marshaler:              marshaler,
		TxConfig:               tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:                  codec.NewLegacyAmino(),
	}
}
