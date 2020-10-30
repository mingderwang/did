package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mingderwang/did/x/did/types"
	"github.com/spf13/cobra"
)

func GetCmdListIdentifier(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-identifier",
		Short: "list all identifier",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListIdentifier, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Identifier\n%s\n", err.Error())
				return nil
			}
			var out []types.Identifier
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetIdentifier(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-identifier [key]",
		Short: "Query a identifier by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetIdentifier, key), nil)
			if err != nil {
				fmt.Printf("could not resolve identifier %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Identifier
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
