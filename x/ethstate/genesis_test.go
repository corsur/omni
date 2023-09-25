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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthstateKeeper(t)
	ethstate.InitGenesis(ctx, *k, genesisState)
	got := ethstate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
