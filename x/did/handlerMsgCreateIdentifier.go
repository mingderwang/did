package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mingderwang/did/x/did/types"
	"github.com/mingderwang/did/x/did/keeper"
)

func handleMsgCreateIdentifier(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateIdentifier) (*sdk.Result, error) {
	var identifier = types.Identifier{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Diddoc: msg.Diddoc,
	}
	k.CreateIdentifier(ctx, identifier)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
