package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteIdentifier{}

type MsgDeleteIdentifier struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteIdentifier(id string, creator sdk.AccAddress) MsgDeleteIdentifier {
  return MsgDeleteIdentifier{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteIdentifier) Route() string {
  return RouterKey
}

func (msg MsgDeleteIdentifier) Type() string {
  return "DeleteIdentifier"
}

func (msg MsgDeleteIdentifier) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteIdentifier) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteIdentifier) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}