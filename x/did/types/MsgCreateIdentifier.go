package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateIdentifier{}

type MsgCreateIdentifier struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Diddoc  string         `json:"diddoc" yaml:"diddoc"`
}

func NewMsgCreateIdentifier(creator sdk.AccAddress, diddoc string) MsgCreateIdentifier {
	return MsgCreateIdentifier{
		ID:      fmt.Sprintf("did:cosm:%s", creator),
		Creator: creator,
		Diddoc:  diddoc,
	}
}

func (msg MsgCreateIdentifier) Route() string {
	return RouterKey
}

func (msg MsgCreateIdentifier) Type() string {
	return "CreateIdentifier"
}

func (msg MsgCreateIdentifier) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateIdentifier) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateIdentifier) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
