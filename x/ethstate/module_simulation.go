package ethstate

import (
	"math/rand"

	"ethState/testutil/sample"
	ethstatesimulation "ethState/x/ethstate/simulation"
	"ethState/x/ethstate/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = ethstatesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateStorageSlot = "op_weight_msg_storage_slot"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStorageSlot int = 100

	opWeightMsgUpdateStorageSlot = "op_weight_msg_storage_slot"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStorageSlot int = 100

	opWeightMsgDeleteStorageSlot = "op_weight_msg_storage_slot"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStorageSlot int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ethstateGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		StorageSlotList: []types.StorageSlot{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		StorageSlotCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ethstateGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateStorageSlot int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateStorageSlot, &weightMsgCreateStorageSlot, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStorageSlot = defaultWeightMsgCreateStorageSlot
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStorageSlot,
		ethstatesimulation.SimulateMsgCreateStorageSlot(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStorageSlot int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateStorageSlot, &weightMsgUpdateStorageSlot, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStorageSlot = defaultWeightMsgUpdateStorageSlot
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStorageSlot,
		ethstatesimulation.SimulateMsgUpdateStorageSlot(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteStorageSlot int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteStorageSlot, &weightMsgDeleteStorageSlot, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStorageSlot = defaultWeightMsgDeleteStorageSlot
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStorageSlot,
		ethstatesimulation.SimulateMsgDeleteStorageSlot(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStorageSlot,
			defaultWeightMsgCreateStorageSlot,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ethstatesimulation.SimulateMsgCreateStorageSlot(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateStorageSlot,
			defaultWeightMsgUpdateStorageSlot,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ethstatesimulation.SimulateMsgUpdateStorageSlot(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteStorageSlot,
			defaultWeightMsgDeleteStorageSlot,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ethstatesimulation.SimulateMsgDeleteStorageSlot(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
