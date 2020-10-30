package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mingderwang/did/x/did/types"
	"github.com/mingderwang/did/x/did/keeper"
)

// Handle a message to delete name
func handleMsgDeleteIdentifier(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteIdentifier) (*sdk.Result, error) {
	if !k.IdentifierExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetIdentifierOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteIdentifier(ctx, msg.ID)
	return &sdk.Result{}, nil
}
