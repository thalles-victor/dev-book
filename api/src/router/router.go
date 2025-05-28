package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate and return a router with route configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
