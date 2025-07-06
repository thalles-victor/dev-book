package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Rote{
	{
		URI:         "/publications",
		Method:      http.MethodPost,
		Function:    controllers.CreatePublication,
		RequireAuth: true,
	},

	{
		URI:         "/publications/{pubID}",
		Method:      http.MethodGet,
		Function:    controllers.SearchPublications,
		RequireAuth: false,
	},

	{
		URI:         "/publications/{pubID}",
		Method:      http.MethodPost,
		Function:    controllers.SearchPublication,
		RequireAuth: false,
	},

	{
		URI:         "/publications/{pubID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePublication,
		RequireAuth: true,
	},

	{
		URI:         "/publications/{pubID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePublication,
		RequireAuth: true,
	},
}
