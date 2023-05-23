package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgJoinPool)(nil)
	_ sdk.Msg = (*MsgExitPool)(nil)
	_ sdk.Msg = (*MsgSwapExactAmountIn)(nil)
	_ sdk.Msg = (*MsgSwapExactAmountOut)(nil)
	_ sdk.Msg = (*MsgJoinSwapExternAmountIn)(nil)
	_ sdk.Msg = (*MsgJoinSwapShareAmountOut)(nil)
	_ sdk.Msg = (*MsgExitSwapShareAmountIn)(nil)
	_ sdk.Msg = (*MsgExitSwapExternAmountOut)(nil)
)

// WARNING: The functions in here are interface stub definitions, not usable in their current state

// Join
func (msg MsgJoinPool) Route() string { return "" }

func (msg MsgJoinPool) Type() string { return "" }

func (msg MsgJoinPool) ValidateBasic() error {
	return nil
}

func (msg MsgJoinPool) GetSignBytes() []byte {
	return nil
}

func (msg MsgJoinPool) GetSigners() []sdk.AccAddress {
	return nil
}

// Exit
func (msg MsgExitPool) Route() string { return "" }

func (msg MsgExitPool) Type() string { return "" }

func (msg MsgExitPool) ValidateBasic() error {
	return nil
}

func (msg MsgExitPool) GetSignBytes() []byte {
	return nil
}

func (msg MsgExitPool) GetSigners() []sdk.AccAddress {
	return nil
}

// SwapExactIn
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

// SwapExactOut
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

// JoinSwapExternIn
func (msg MsgJoinSwapExternAmountIn) Route() string { return "" }

func (msg MsgJoinSwapExternAmountIn) Type() string { return "" }

func (msg MsgJoinSwapExternAmountIn) ValidateBasic() error {
	return nil
}

func (msg MsgJoinSwapExternAmountIn) GetSignBytes() []byte {
	return nil
}

func (msg MsgJoinSwapExternAmountIn) GetSigners() []sdk.AccAddress {
	return nil
}

// JoinSwapShareOut
func (msg MsgJoinSwapShareAmountOut) Route() string { return "" }

func (msg MsgJoinSwapShareAmountOut) Type() string { return "" }

func (msg MsgJoinSwapShareAmountOut) ValidateBasic() error {
	return nil
}

func (msg MsgJoinSwapShareAmountOut) GetSignBytes() []byte {
	return nil
}

func (msg MsgJoinSwapShareAmountOut) GetSigners() []sdk.AccAddress {
	return nil
}

// ExitSwapShareIn
func (msg MsgExitSwapShareAmountIn) Route() string { return "" }

func (msg MsgExitSwapShareAmountIn) Type() string { return "" }

func (msg MsgExitSwapShareAmountIn) ValidateBasic() error {
	return nil
}

func (msg MsgExitSwapShareAmountIn) GetSignBytes() []byte {
	return nil
}

func (msg MsgExitSwapShareAmountIn) GetSigners() []sdk.AccAddress {
	return nil
}

// ExitSwapExternOut
func (msg MsgExitSwapExternAmountOut) Route() string { return "" }

func (msg MsgExitSwapExternAmountOut) Type() string { return "" }

func (msg MsgExitSwapExternAmountOut) ValidateBasic() error {
	return nil
}

func (msg MsgExitSwapExternAmountOut) GetSignBytes() []byte {
	return nil
}

func (msg MsgExitSwapExternAmountOut) GetSigners() []sdk.AccAddress {
	return nil
}
