package rest

import (
	"github.com/gorilla/mux"

	"github.com/amaurymartiny/step3/x/step3/types"
	"github.com/cosmos/cosmos-sdk/client"
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers step3-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("custom/step3/"+types.QueryListComment, listCommentHandler(clientCtx)).Methods("GET")

	r.HandleFunc("custom/step3/"+types.QueryListPost, listPostHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/step3/comment", createCommentHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/step3/post", createPostHandler(clientCtx)).Methods("POST")

}
