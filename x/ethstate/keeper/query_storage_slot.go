package keeper

import (
	"context"

	"ethState/x/ethstate/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StorageSlotAll(goCtx context.Context, req *types.QueryAllStorageSlotRequest) (*types.QueryAllStorageSlotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storageSlots []types.StorageSlot
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	storageSlotStore := prefix.NewStore(store, types.KeyPrefix(types.StorageSlotKey))

	pageRes, err := query.Paginate(storageSlotStore, req.Pagination, func(key []byte, value []byte) error {
		var storageSlot types.StorageSlot
		if err := k.cdc.Unmarshal(value, &storageSlot); err != nil {
			return err
		}

		storageSlots = append(storageSlots, storageSlot)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStorageSlotResponse{StorageSlot: storageSlots, Pagination: pageRes}, nil
}

func (k Keeper) StorageSlot(goCtx context.Context, req *types.QueryGetStorageSlotRequest) (*types.QueryGetStorageSlotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storageSlot, found := k.GetStorageSlot(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetStorageSlotResponse{StorageSlot: storageSlot}, nil
}
