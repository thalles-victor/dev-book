package routes

import (
	"api/src/middlewares"
	"fmt"
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
	routes = append(routes, loginRoute)

	for _, route := range routes {
		routeLog := fmt.Sprintf("mapped -> uri: %s  method: %s auth: %t", route.URI, route.Method, route.RequireAuth)
		fmt.Println(routeLog)

		if route.RequireAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authentication(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
