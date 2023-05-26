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
	//Old Osmosis GAMM message types
	_ sdk.Msg = (*MsgCreatePool)(nil)
	_ sdk.Msg = (*MsgCreateBalancerPool)(nil)
)

// WARNING: The functions in here are interface stub definitions, not usable in their current state

// Join
func (msg MsgJoinPool) Route() string { return "" }

func (msg MsgJoinPool) Type() string { return "" }

func (msg MsgJoinPool) ValidateBasic() error {
	panic("MsgJoinPool ValidateBasic Unimplemented")
}

func (msg MsgJoinPool) GetSignBytes() []byte {
	panic("MsgJoinPool GetSignBytes Unimplemented")
}

func (msg MsgJoinPool) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// Exit
func (msg MsgExitPool) Route() string { return "" }

func (msg MsgExitPool) Type() string { return "" }

func (msg MsgExitPool) ValidateBasic() error {
	panic("MsgExitPool ValidateBasic Unimplemented")
}

func (msg MsgExitPool) GetSignBytes() []byte {
	panic("MsgExitPool GetSignBytes Unimplemented")
}

func (msg MsgExitPool) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// SwapExactIn
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

// SwapExactOut
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

// JoinSwapExternIn
func (msg MsgJoinSwapExternAmountIn) Route() string { return "" }

func (msg MsgJoinSwapExternAmountIn) Type() string { return "" }

func (msg MsgJoinSwapExternAmountIn) ValidateBasic() error {
	panic("MsgJoinSwapExternAmountIn ValidateBasic Unimplemented")
}

func (msg MsgJoinSwapExternAmountIn) GetSignBytes() []byte {
	panic("MsgJoinSwapExternAmountIn GetSignBytes Unimplemented")
}

func (msg MsgJoinSwapExternAmountIn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// JoinSwapShareOut
func (msg MsgJoinSwapShareAmountOut) Route() string { return "" }

func (msg MsgJoinSwapShareAmountOut) Type() string { return "" }

func (msg MsgJoinSwapShareAmountOut) ValidateBasic() error {
	panic("MsgJoinSwapShareAmountOut ValidateBasic Unimplemented")
}

func (msg MsgJoinSwapShareAmountOut) GetSignBytes() []byte {
	panic("MsgJoinSwapShareAmountOut GetSignBytes Unimplemented")
}

func (msg MsgJoinSwapShareAmountOut) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ExitSwapShareIn
func (msg MsgExitSwapShareAmountIn) Route() string { return "" }

func (msg MsgExitSwapShareAmountIn) Type() string { return "" }

func (msg MsgExitSwapShareAmountIn) ValidateBasic() error {
	panic("MsgExitSwapShareAmountIn ValidateBasic Unimplemented")
}

func (msg MsgExitSwapShareAmountIn) GetSignBytes() []byte {
	panic("MsgExitSwapShareAmountIn GetSignBytes Unimplemented")
}

func (msg MsgExitSwapShareAmountIn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ExitSwapExternOut
func (msg MsgExitSwapExternAmountOut) Route() string { return "" }

func (msg MsgExitSwapExternAmountOut) Type() string { return "" }

func (msg MsgExitSwapExternAmountOut) ValidateBasic() error {
	panic("MsgExitSwapExternAmountOut ValidateBasic Unimplemented")
}

func (msg MsgExitSwapExternAmountOut) GetSignBytes() []byte {
	panic("MsgExitSwapExternAmountOut GetSignBytes Unimplemented")
}

func (msg MsgExitSwapExternAmountOut) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// Old Message Types

// CreatePool
func (msg MsgCreatePool) Route() string { return "" }

func (msg MsgCreatePool) Type() string { return "" }

func (msg MsgCreatePool) ValidateBasic() error {
	panic("MsgCreatePool ValidateBasic Unimplemented")
}

func (msg MsgCreatePool) GetSignBytes() []byte {
	panic("MsgCreatePool GetSignBytes Unimplemented")
}

func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// CreatePool
func (msg MsgCreateBalancerPool) Route() string { return "" }

func (msg MsgCreateBalancerPool) Type() string { return "" }

func (msg MsgCreateBalancerPool) ValidateBasic() error {
	panic("MsgCreateBalancerPool ValidateBasic Unimplemented")
}

func (msg MsgCreateBalancerPool) GetSignBytes() []byte {
	panic("MsgCreateBalancerPool GetSignBytes Unimplemented")
}

func (msg MsgCreateBalancerPool) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
