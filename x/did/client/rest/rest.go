package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers did-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/did/identifier", createIdentifierHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/did/identifier", listIdentifierHandler(cliCtx, "did")).Methods("GET")
		r.HandleFunc("/did/identifier/{key}", getIdentifierHandler(cliCtx, "did")).Methods("GET")
		r.HandleFunc("/did/identifier", setIdentifierHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/did/identifier", deleteIdentifierHandler(cliCtx)).Methods("DELETE")

		
}
