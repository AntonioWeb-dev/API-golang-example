package routes

import (
	"api/src/controllers"
	"net/http"
)

var routerLogin = Route{
	URI:            "/login",
	Method:         http.MethodPost,
	Controller:     controllers.Login,
	Authentication: false,
}
