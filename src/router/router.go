package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Create return the routes
func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRouters(r)
}
