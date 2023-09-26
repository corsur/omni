package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateStorageSlot = "create_storage_slot"
	TypeMsgUpdateStorageSlot = "update_storage_slot"
	TypeMsgDeleteStorageSlot = "delete_storage_slot"
)

var _ sdk.Msg = &MsgCreateStorageSlot{}

func NewMsgCreateStorageSlot(creator string, address string, storage string) *MsgCreateStorageSlot {
	return &MsgCreateStorageSlot{
		Creator: creator,
		Address: address,
		Storage: storage,
	}
}

func (msg *MsgCreateStorageSlot) Route() string {
	return RouterKey
}

func (msg *MsgCreateStorageSlot) Type() string {
	return TypeMsgCreateStorageSlot
}

func (msg *MsgCreateStorageSlot) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStorageSlot) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStorageSlot) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStorageSlot{}

func NewMsgUpdateStorageSlot(creator string, id uint64, address string, storage string) *MsgUpdateStorageSlot {
	return &MsgUpdateStorageSlot{
		Id:      id,
		Creator: creator,
		Address: address,
		Storage: storage,
	}
}

func (msg *MsgUpdateStorageSlot) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStorageSlot) Type() string {
	return TypeMsgUpdateStorageSlot
}

func (msg *MsgUpdateStorageSlot) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStorageSlot) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStorageSlot) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStorageSlot{}

func NewMsgDeleteStorageSlot(creator string, id uint64) *MsgDeleteStorageSlot {
	return &MsgDeleteStorageSlot{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteStorageSlot) Route() string {
	return RouterKey
}

func (msg *MsgDeleteStorageSlot) Type() string {
	return TypeMsgDeleteStorageSlot
}

func (msg *MsgDeleteStorageSlot) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteStorageSlot) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteStorageSlot) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
