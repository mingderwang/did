package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Identifier struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Diddoc string `json:"diddoc" yaml:"diddoc"`
}