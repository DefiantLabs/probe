package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgSwapExactAmountIn)(nil)
	_ sdk.Msg = (*MsgSplitRouteSwapExactAmountIn)(nil)
	_ sdk.Msg = (*MsgSwapExactAmountOut)(nil)
	_ sdk.Msg = (*MsgSplitRouteSwapExactAmountOut)(nil)
)

// WARNING: The functions in here are interface stub definitions, not usable in their current state

// SwapExactAmountIn
func (msg MsgSwapExactAmountIn) Route() string { return "" }

func (msg MsgSwapExactAmountIn) Type() string { return "" }

func (msg MsgSwapExactAmountIn) ValidateBasic() error {
	panic("MsgSwapExactAmountIn ValidateBasic Unimplemented")
}

func (msg MsgSwapExactAmountIn) GetSignBytes() []byte {
	panic("MsgSwapExactAmountIn GetSignBytes Unimplemented")
}

func (msg MsgSwapExactAmountIn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// SplitRouteSwapExactAmountIn
func (msg MsgSplitRouteSwapExactAmountIn) Route() string { return "" }

func (msg MsgSplitRouteSwapExactAmountIn) Type() string { return "" }

func (msg MsgSplitRouteSwapExactAmountIn) ValidateBasic() error {
	panic("MsgSplitRouteSwapExactAmountIn ValidateBasic Unimplemented")
}

func (msg MsgSplitRouteSwapExactAmountIn) GetSignBytes() []byte {
	panic("MsgSplitRouteSwapExactAmountIn GetSignBytes Unimplemented")
}

func (msg MsgSplitRouteSwapExactAmountIn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// SwapExactAmountOut
func (msg MsgSwapExactAmountOut) Route() string { return "" }

func (msg MsgSwapExactAmountOut) Type() string { return "" }

func (msg MsgSwapExactAmountOut) ValidateBasic() error {
	panic("MsgSwapExactAmountOut ValidateBasic Unimplemented")
}

func (msg MsgSwapExactAmountOut) GetSignBytes() []byte {
	panic("MsgSwapExactAmountOut GetSignBytes Unimplemented")
}

func (msg MsgSwapExactAmountOut) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// SplitRouteSwapExactAmountOut
func (msg MsgSplitRouteSwapExactAmountOut) Route() string { return "" }

func (msg MsgSplitRouteSwapExactAmountOut) Type() string { return "" }

func (msg MsgSplitRouteSwapExactAmountOut) ValidateBasic() error {
	panic("MsgSplitRouteSwapExactAmountOut ValidateBasic Unimplemented")
}

func (msg MsgSplitRouteSwapExactAmountOut) GetSignBytes() []byte {
	panic("MsgSplitRouteSwapExactAmountOut GetSignBytes Unimplemented")
}

func (msg MsgSplitRouteSwapExactAmountOut) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
