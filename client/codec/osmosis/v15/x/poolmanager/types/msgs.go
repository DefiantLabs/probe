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
	return nil
}

func (msg MsgSwapExactAmountIn) GetSignBytes() []byte {
	return nil
}

func (msg MsgSwapExactAmountIn) GetSigners() []sdk.AccAddress {
	return nil
}

// SplitRouteSwapExactAmountIn
func (msg MsgSplitRouteSwapExactAmountIn) Route() string { return "" }

func (msg MsgSplitRouteSwapExactAmountIn) Type() string { return "" }

func (msg MsgSplitRouteSwapExactAmountIn) ValidateBasic() error {
	return nil
}

func (msg MsgSplitRouteSwapExactAmountIn) GetSignBytes() []byte {
	return nil
}

func (msg MsgSplitRouteSwapExactAmountIn) GetSigners() []sdk.AccAddress {
	return nil
}

// SwapExactAmountOut
func (msg MsgSwapExactAmountOut) Route() string { return "" }

func (msg MsgSwapExactAmountOut) Type() string { return "" }

func (msg MsgSwapExactAmountOut) ValidateBasic() error {
	return nil
}

func (msg MsgSwapExactAmountOut) GetSignBytes() []byte {
	return nil
}

func (msg MsgSwapExactAmountOut) GetSigners() []sdk.AccAddress {
	return nil
}

// SplitRouteSwapExactAmountOut
func (msg MsgSplitRouteSwapExactAmountOut) Route() string { return "" }

func (msg MsgSplitRouteSwapExactAmountOut) Type() string { return "" }

func (msg MsgSplitRouteSwapExactAmountOut) ValidateBasic() error {
	return nil
}

func (msg MsgSplitRouteSwapExactAmountOut) GetSignBytes() []byte {
	return nil
}

func (msg MsgSplitRouteSwapExactAmountOut) GetSigners() []sdk.AccAddress {
	return nil
}
