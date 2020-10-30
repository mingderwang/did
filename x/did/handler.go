package did

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mingderwang/did/x/did/keeper"
	"github.com/mingderwang/did/x/did/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateIdentifier:
			return handleMsgCreateIdentifier(ctx, k, msg)
		case types.MsgSetIdentifier:
			return handleMsgSetIdentifier(ctx, k, msg)
		case types.MsgDeleteIdentifier:
			return handleMsgDeleteIdentifier(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
