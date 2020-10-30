package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mingderwang/did/x/did/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateIdentifier creates a identifier
func (k Keeper) CreateIdentifier(ctx sdk.Context, identifier types.Identifier) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.IdentifierPrefix + identifier.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(identifier)
	store.Set(key, value)
}

// GetIdentifier returns the identifier information
func (k Keeper) GetIdentifier(ctx sdk.Context, key string) (types.Identifier, error) {
	store := ctx.KVStore(k.storeKey)
	var identifier types.Identifier
	byteKey := []byte(types.IdentifierPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &identifier)
	if err != nil {
		return identifier, err
	}
	return identifier, nil
}

// SetIdentifier sets a identifier
func (k Keeper) SetIdentifier(ctx sdk.Context, identifier types.Identifier) {
	identifierKey := identifier.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(identifier)
	key := []byte(types.IdentifierPrefix + identifierKey)
	store.Set(key, bz)
}

// DeleteIdentifier deletes a identifier
func (k Keeper) DeleteIdentifier(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.IdentifierPrefix + key))
}

//
// Functions used by querier
//

func listIdentifier(ctx sdk.Context, k Keeper) ([]byte, error) {
	var identifierList []types.Identifier
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.IdentifierPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var identifier types.Identifier
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &identifier)
		identifierList = append(identifierList, identifier)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, identifierList)
	return res, nil
}

func getIdentifier(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	identifier, err := k.GetIdentifier(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, identifier)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetIdentifierOwner(ctx sdk.Context, key string) sdk.AccAddress {
	identifier, err := k.GetIdentifier(ctx, key)
	if err != nil {
		return nil
	}
	return identifier.Creator
}


// Check if the key exists in the store
func (k Keeper) IdentifierExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.IdentifierPrefix + key))
}
