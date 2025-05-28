package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Represent routes structure
type Rote struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Pull all routes into the router (r)
func Configure(r *mux.Router) *mux.Router {
	routes := userRouters

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
