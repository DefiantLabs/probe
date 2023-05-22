package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgCreatePool)(nil)
	_ sdk.Msg = (*MsgDepositWithinBatch)(nil)
	_ sdk.Msg = (*MsgWithdrawWithinBatch)(nil)
	_ sdk.Msg = (*MsgSwapWithinBatch)(nil)
)

// Message types for the liquidity module
const (
	TypeMsgCreatePool          = "create_pool"
	TypeMsgDepositWithinBatch  = "deposit_within_batch"
	TypeMsgWithdrawWithinBatch = "withdraw_within_batch"
	TypeMsgSwapWithinBatch     = "swap_within_batch"
)

func (msg MsgCreatePool) Route() string { return "" }

func (msg MsgCreatePool) Type() string { return TypeMsgCreatePool }

func (msg MsgCreatePool) ValidateBasic() error {
	return nil
}

func (msg MsgCreatePool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.PoolCreatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgCreatePool) GetPoolCreator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.PoolCreatorAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgDepositWithinBatch creates a new MsgDepositWithinBatch.
func NewMsgDepositWithinBatch(depositor sdk.AccAddress, poolID uint64, depositCoins sdk.Coins) *MsgDepositWithinBatch {
	return &MsgDepositWithinBatch{
		DepositorAddress: depositor.String(),
		PoolId:           poolID,
		DepositCoins:     depositCoins,
	}
}

func (msg MsgDepositWithinBatch) Route() string { return "" }

func (msg MsgDepositWithinBatch) Type() string { return TypeMsgDepositWithinBatch }

func (msg MsgDepositWithinBatch) ValidateBasic() error {
	return nil
}

func (msg MsgDepositWithinBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDepositWithinBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.DepositorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgDepositWithinBatch) GetDepositor() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.DepositorAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgWithdrawWithinBatch creates a new MsgWithdrawWithinBatch.
func NewMsgWithdrawWithinBatch(withdrawer sdk.AccAddress, poolID uint64, poolCoin sdk.Coin) *MsgWithdrawWithinBatch {
	return &MsgWithdrawWithinBatch{
		WithdrawerAddress: withdrawer.String(),
		PoolId:            poolID,
		PoolCoin:          poolCoin,
	}
}

func (msg MsgWithdrawWithinBatch) Route() string { return "" }

func (msg MsgWithdrawWithinBatch) Type() string { return TypeMsgWithdrawWithinBatch }

func (msg MsgWithdrawWithinBatch) ValidateBasic() error {
	return nil
}

func (msg MsgWithdrawWithinBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawWithinBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.WithdrawerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgWithdrawWithinBatch) GetWithdrawer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.WithdrawerAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

func (msg MsgSwapWithinBatch) Route() string { return "" }

func (msg MsgSwapWithinBatch) Type() string { return TypeMsgSwapWithinBatch }

func (msg MsgSwapWithinBatch) ValidateBasic() error {
	return nil
}

func (msg MsgSwapWithinBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSwapWithinBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.SwapRequesterAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgSwapWithinBatch) GetSwapRequester() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.SwapRequesterAddress)
	if err != nil {
		panic(err)
	}
	return addr
}
