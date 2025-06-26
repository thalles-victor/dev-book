package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Rote{
	URI:         "/login",
	Method:      http.MethodPost,
	Function:    controllers.Login,
	RequireAuth: false,
}
