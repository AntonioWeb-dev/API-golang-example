package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct to build routers
type Route struct {
	URI            string
	Method         string
	Controller     func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// ConfigRouters - join all routes configs
func ConfigRouters(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, routerLogin)

	for _, router := range routes {
		if router.Authentication {
			r.HandleFunc(router.URI,
				middlewares.Logger(
					middlewares.Authentication(router.Controller),
				)).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI, middlewares.Logger(router.Controller)).Methods(router.Method)
		}
	}
	return r
}
