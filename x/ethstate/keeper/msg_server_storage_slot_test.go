package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"ethState/x/ethstate/types"
)

func TestStorageSlotMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateStorageSlot(ctx, &types.MsgCreateStorageSlot{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestStorageSlotMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateStorageSlot
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateStorageSlot{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateStorageSlot{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateStorageSlot{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateStorageSlot(ctx, &types.MsgCreateStorageSlot{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateStorageSlot(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStorageSlotMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteStorageSlot
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteStorageSlot{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteStorageSlot{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteStorageSlot{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateStorageSlot(ctx, &types.MsgCreateStorageSlot{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteStorageSlot(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
