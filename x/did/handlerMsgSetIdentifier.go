package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mingderwang/did/x/did/types"
	"github.com/mingderwang/did/x/did/keeper"
)

func handleMsgSetIdentifier(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetIdentifier) (*sdk.Result, error) {
	var identifier = types.Identifier{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Diddoc: msg.Diddoc,
	}
	if !msg.Creator.Equals(k.GetIdentifierOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetIdentifier(ctx, identifier)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
