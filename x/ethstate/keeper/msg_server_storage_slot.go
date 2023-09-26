package keeper

import (
	"context"
	"fmt"

	"ethState/x/ethstate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStorageSlot(goCtx context.Context, msg *types.MsgCreateStorageSlot) (*types.MsgCreateStorageSlotResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var storageSlot = types.StorageSlot{
		Creator: msg.Creator,
		Address: msg.Address,
		Storage: msg.Storage,
	}

	id := k.AppendStorageSlot(
		ctx,
		storageSlot,
	)

	return &types.MsgCreateStorageSlotResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateStorageSlot(goCtx context.Context, msg *types.MsgUpdateStorageSlot) (*types.MsgUpdateStorageSlotResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var storageSlot = types.StorageSlot{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
		Storage: msg.Storage,
	}

	// Checks that the element exists
	val, found := k.GetStorageSlot(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetStorageSlot(ctx, storageSlot)

	return &types.MsgUpdateStorageSlotResponse{}, nil
}

func (k msgServer) DeleteStorageSlot(goCtx context.Context, msg *types.MsgDeleteStorageSlot) (*types.MsgDeleteStorageSlotResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetStorageSlot(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStorageSlot(ctx, msg.Id)

	return &types.MsgDeleteStorageSlotResponse{}, nil
}
