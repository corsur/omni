package ethstate_test

import (
	"testing"

	keepertest "ethState/testutil/keeper"
	"ethState/testutil/nullify"
	"ethState/x/ethstate"
	"ethState/x/ethstate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StorageSlotList: []types.StorageSlot{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		StorageSlotCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthstateKeeper(t)
	ethstate.InitGenesis(ctx, *k, genesisState)
	got := ethstate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StorageSlotList, got.StorageSlotList)
	require.Equal(t, genesisState.StorageSlotCount, got.StorageSlotCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
