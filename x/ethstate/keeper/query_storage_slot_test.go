package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "ethState/testutil/keeper"
	"ethState/testutil/nullify"
	"ethState/x/ethstate/types"
)

func TestStorageSlotQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EthstateKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStorageSlot(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetStorageSlotRequest
		response *types.QueryGetStorageSlotResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetStorageSlotRequest{Id: msgs[0].Id},
			response: &types.QueryGetStorageSlotResponse{StorageSlot: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetStorageSlotRequest{Id: msgs[1].Id},
			response: &types.QueryGetStorageSlotResponse{StorageSlot: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetStorageSlotRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.StorageSlot(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestStorageSlotQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EthstateKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStorageSlot(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStorageSlotRequest {
		return &types.QueryAllStorageSlotRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StorageSlotAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.StorageSlot), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.StorageSlot),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StorageSlotAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.StorageSlot), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.StorageSlot),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StorageSlotAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.StorageSlot),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StorageSlotAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
