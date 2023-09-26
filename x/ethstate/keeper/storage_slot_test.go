package keeper_test

import (
	"testing"

	keepertest "ethState/testutil/keeper"
	"ethState/testutil/nullify"
	"ethState/x/ethstate/keeper"
	"ethState/x/ethstate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNStorageSlot(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StorageSlot {
	items := make([]types.StorageSlot, n)
	for i := range items {
		items[i].Id = keeper.AppendStorageSlot(ctx, items[i])
	}
	return items
}

func TestStorageSlotGet(t *testing.T) {
	keeper, ctx := keepertest.EthstateKeeper(t)
	items := createNStorageSlot(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetStorageSlot(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestStorageSlotRemove(t *testing.T) {
	keeper, ctx := keepertest.EthstateKeeper(t)
	items := createNStorageSlot(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStorageSlot(ctx, item.Id)
		_, found := keeper.GetStorageSlot(ctx, item.Id)
		require.False(t, found)
	}
}

func TestStorageSlotGetAll(t *testing.T) {
	keeper, ctx := keepertest.EthstateKeeper(t)
	items := createNStorageSlot(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStorageSlot(ctx)),
	)
}

func TestStorageSlotCount(t *testing.T) {
	keeper, ctx := keepertest.EthstateKeeper(t)
	items := createNStorageSlot(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetStorageSlotCount(ctx))
}
