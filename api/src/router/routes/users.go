package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRouters = []Rote{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.SearchUsers,
		RequireAuth: false,
	}, {
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.SearchUser,
		RequireAuth: false,
	}, {
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/follow",
		Method:      http.MethodPost,
		Function:    controllers.Follow,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnFollowing,
		RequireAuth: true,
	},
}
