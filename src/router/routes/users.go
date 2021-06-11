package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []DefaultRoute{
	{
		URI:             "/users",
		Method:          http.MethodPost,
		Function:        controllers.AddUser,
		IsAuthenticated: false,
	},
}
