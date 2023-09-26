package types

import (
	"testing"

	"ethState/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateStorageSlot_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateStorageSlot
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateStorageSlot{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateStorageSlot{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateStorageSlot_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateStorageSlot
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateStorageSlot{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateStorageSlot{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteStorageSlot_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteStorageSlot
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteStorageSlot{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteStorageSlot{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
