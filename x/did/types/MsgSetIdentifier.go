package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetIdentifier{}

type MsgSetIdentifier struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Diddoc string `json:"diddoc" yaml:"diddoc"`
}

func NewMsgSetIdentifier(creator sdk.AccAddress, id string, diddoc string) MsgSetIdentifier {
  return MsgSetIdentifier{
    ID: id,
		Creator: creator,
    Diddoc: diddoc,
	}
}

func (msg MsgSetIdentifier) Route() string {
  return RouterKey
}

func (msg MsgSetIdentifier) Type() string {
  return "SetIdentifier"
}

func (msg MsgSetIdentifier) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetIdentifier) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetIdentifier) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}