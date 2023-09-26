package keeper

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"

	"bytes"
	"fmt"
	"net/http"

	"ethState/x/ethstate/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type JSONResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

// GetStorageSlotCount get the total number of storageSlot
func (k Keeper) GetStorageSlotCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.StorageSlotCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetStorageSlotCount set the total number of storageSlot
func (k Keeper) SetStorageSlotCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.StorageSlotCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetStorageData returns the data stored at the given address and slot
func GetStorageData(address, slot string) string {

	url := "https://mainnet.infura.io/v3/23e42bffc6a04ed79aa2c9fa01dd2fa9"

	payload := []byte(`{
		"method": "eth_getStorageAt",
		"params": [` + address + `,` + slot + `,"latest"],
		"id": 1,
		"jsonrpc": "2.0"
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	// Read and print the response body
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println("Response Body:", buf.String())

	// Define a struct to unmarshal the JSON response
	var jsonResponse JSONResponse

	// Unmarshal the JSON response
	err2 := json.Unmarshal([]byte(buf.String()), &jsonResponse)
	if err2 != nil {
		fmt.Println("Error:", err)
	}

	// Access the "result" field
	return jsonResponse.Result

}

// AppendStorageSlot appends a storageSlot in the store with a new id and update the count
func (k Keeper) AppendStorageSlot(
	ctx sdk.Context,
	storageSlot types.StorageSlot,
) uint64 {
	// Get the data stored at the given address and slot
	data := GetStorageData(storageSlot.Address, storageSlot.Storage)

	// Remove the "0x" prefix if present
	dataHexString := data[2:]

	// Convert the hexadecimal string to bytes
	storageBytes, err := hex.DecodeString(dataHexString)
	if err != nil {
		fmt.Println("Error decoding hexadecimal string:", err)
	}

	// Create the storageSlot
	count := k.GetStorageSlotCount(ctx)

	// Create a StorageSlot instance with the storageBytes
	newStorageSlot := types.StorageSlot{
		Id:      count,                // Set the ID as needed
		Address: storageSlot.Address,  // Set the address as needed
		Storage: string(storageBytes), // Convert bytes to a string
		Creator: storageSlot.Creator,  // Set the creator as needed
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageSlotKey))
	appendedValue := k.cdc.MustMarshal(&newStorageSlot)
	store.Set(GetStorageSlotIDBytes(storageSlot.Id), appendedValue)

	// Update storageSlot count
	k.SetStorageSlotCount(ctx, count+1)

	return count
}

// SetStorageSlot set a specific storageSlot in the store
func (k Keeper) SetStorageSlot(ctx sdk.Context, storageSlot types.StorageSlot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageSlotKey))
	b := k.cdc.MustMarshal(&storageSlot)
	store.Set(GetStorageSlotIDBytes(storageSlot.Id), b)
}

// GetStorageSlot returns a storageSlot from its id
func (k Keeper) GetStorageSlot(ctx sdk.Context, id uint64) (val types.StorageSlot, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageSlotKey))
	b := store.Get(GetStorageSlotIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStorageSlot removes a storageSlot from the store
func (k Keeper) RemoveStorageSlot(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageSlotKey))
	store.Delete(GetStorageSlotIDBytes(id))
}

// GetAllStorageSlot returns all storageSlot
func (k Keeper) GetAllStorageSlot(ctx sdk.Context) (list []types.StorageSlot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageSlotKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StorageSlot
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStorageSlotIDBytes returns the byte representation of the ID
func GetStorageSlotIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetStorageSlotIDFromBytes returns ID in uint64 format from a byte array
func GetStorageSlotIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
