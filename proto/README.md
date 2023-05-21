# Proto Definitions and Generation

## Why Module Protobuf definitions and code generation is needed:

The main use-case of this repo is for generalized querying on Cosmos blockchains.

For the use-case of Transaction and Message type querying, the package needs some way to unpack the message types found on-chain into their underlying Go types. The message types all are expected to follow the Cosmos Msg interface, with their representation stored onchain in byte form. These bytes need to be converted to the Go Msg types by this package to support transaction querying.

When messages are received from the chain during Transaction querying, they are unpacked into their underlying type. The message type is found using the TypeUrl for the message, pulled from the client Codec, and unpacked.

For base Cosmos SDK types, this process is trivial. This package imports the Cosmos SDK and sets up the Codec to handle these types.

For customized modules, the process is more involved. The types cannot always be imported directly into this codebase due to (amongst other considerations): Random blockchains are not always using Cosmos SDK v47, meaning they are incompatible with this package if imported.

## Our Solution

We are making use of the following solution where needed:

1. Determine which Blockchain modules need to be supported by this package
2. Bring in the Module Msg proto definitions
3. Generate the .pb.go files using `buf`
4. Fill out the remaining Msg interface function definition requirements
5. Register the Msg types into the application Codec

This allows Transaction querying to unpack these messages into their underlying type so that values can be pulled out of them.

## Requirements

Before contributing proto definitions here, make sure your environment is setup for generating Cosmos Proto files by following these requirements:

1. Download Cosmos Gogoproto

```sh
git clone https://github.com/cosmos/gogoproto.git
cd gogoproto
go mod download
make install
```

2. Make sure your proto definitions fit the application package requirements. Primarily, due to the location of auto-generated code living under client/codec/\<repo\>/\<module\>, make sure your proto package name reflects this location.
