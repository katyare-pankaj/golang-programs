// pkg/server/server.go
package server

import (
	"go-programs/RLHF/golang_random/29-11-24/389143/turn1/modelA/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(port string) {
	r := mux.NewRouter()
	handlers.RegisterRoutes(r)

	http.ListenAndServe(":"+port, r)
}
