package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgLockTokens)(nil)
	_ sdk.Msg = (*MsgBeginUnlockingAll)(nil)
	_ sdk.Msg = (*MsgBeginUnlocking)(nil)
	_ sdk.Msg = (*MsgExtendLockup)(nil)
	_ sdk.Msg = (*MsgForceUnlock)(nil)
	//Old Osmosis Lockup types
	_ sdk.Msg = (*MsgUnlockPeriodLock)(nil)
	_ sdk.Msg = (*MsgUnlockTokens)(nil)
)

// WARNING: The functions in here are interface stub definitions, not usable in their current state

// MsgLockTokens
func (msg MsgLockTokens) Route() string { return "" }

func (msg MsgLockTokens) Type() string { return "" }

func (msg MsgLockTokens) ValidateBasic() error {
	panic("MsgLockTokens ValidateBasic Unimplemented")
}

func (msg MsgLockTokens) GetSignBytes() []byte {
	panic("MsgLockTokens GetSignBytes Unimplemented")
}

func (msg MsgLockTokens) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgBeginUnlockingAll
func (msg MsgBeginUnlockingAll) Route() string { return "" }

func (msg MsgBeginUnlockingAll) Type() string { return "" }

func (msg MsgBeginUnlockingAll) ValidateBasic() error {
	panic("MsgBeginUnlockingAll ValidateBasic Unimplemented")
}

func (msg MsgBeginUnlockingAll) GetSignBytes() []byte {
	panic("MsgBeginUnlockingAll GetSignBytes Unimplemented")
}

func (msg MsgBeginUnlockingAll) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgBeginUnlocking
func (msg MsgBeginUnlocking) Route() string { return "" }

func (msg MsgBeginUnlocking) Type() string { return "" }

func (msg MsgBeginUnlocking) ValidateBasic() error {
	panic("MsgBeginUnlocking ValidateBasic Unimplemented")
}

func (msg MsgBeginUnlocking) GetSignBytes() []byte {
	panic("MsgBeginUnlocking GetSignBytes Unimplemented")
}

func (msg MsgBeginUnlocking) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgExtendLockup
func (msg MsgExtendLockup) Route() string { return "" }

func (msg MsgExtendLockup) Type() string { return "" }

func (msg MsgExtendLockup) ValidateBasic() error {
	panic("MsgExtendLockup ValidateBasic Unimplemented")
}

func (msg MsgExtendLockup) GetSignBytes() []byte {
	panic("MsgExtendLockup GetSignBytes Unimplemented")
}

func (msg MsgExtendLockup) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgForceUnlock
func (msg MsgForceUnlock) Route() string { return "" }

func (msg MsgForceUnlock) Type() string { return "" }

func (msg MsgForceUnlock) ValidateBasic() error {
	panic("MsgForceUnlock ValidateBasic Unimplemented")
}

func (msg MsgForceUnlock) GetSignBytes() []byte {
	panic("MsgForceUnlock GetSignBytes Unimplemented")
}

func (msg MsgForceUnlock) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// Old Osmosis Lockup types
// MsgUnlockPeriodLock
func (msg MsgUnlockPeriodLock) Route() string { return "" }

func (msg MsgUnlockPeriodLock) Type() string { return "" }

func (msg MsgUnlockPeriodLock) ValidateBasic() error {
	panic("MsgUnlockPeriodLock ValidateBasic Unimplemented")
}

func (msg MsgUnlockPeriodLock) GetSignBytes() []byte {
	panic("MsgUnlockPeriodLock GetSignBytes Unimplemented")
}

func (msg MsgUnlockPeriodLock) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgUnlockTokens
func (msg MsgUnlockTokens) Route() string { return "" }

func (msg MsgUnlockTokens) Type() string { return "" }

func (msg MsgUnlockTokens) ValidateBasic() error {
	panic("MsgUnlockTokens ValidateBasic Unimplemented")
}

func (msg MsgUnlockTokens) GetSignBytes() []byte {
	panic("MsgUnlockTokens GetSignBytes Unimplemented")
}

func (msg MsgUnlockTokens) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
