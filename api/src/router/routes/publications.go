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
		URI:         "/publications",
		Method:      http.MethodGet,
		Function:    controllers.SearchPublication,
		RequireAuth: true,
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

	{
		URI:         "/users/{userID}/publications",
		Method:      http.MethodGet,
		Function:    controllers.SearchPubsByUsers,
		RequireAuth: true,
	},

	{
		URI:         "/publications/{pubID}/like",
		Method:      http.MethodPost,
		Function:    controllers.Like,
		RequireAuth: true,
	},
}
