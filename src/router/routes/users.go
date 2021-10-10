package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Controller:     controllers.CreateUser,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Controller:     controllers.GetUsers,
		Authentication: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Controller:     controllers.GetUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Controller:     controllers.DeleteUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Controller:     controllers.UpdateUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}/follow",
		Method:         http.MethodPost,
		Controller:     controllers.FollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}/unfollow",
		Method:         http.MethodDelete,
		Controller:     controllers.UnFollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}/followers",
		Method:         http.MethodGet,
		Controller:     controllers.GetFollowers,
		Authentication: true,
	},
	{
		URI:            "/users/{id}/following",
		Method:         http.MethodGet,
		Controller:     controllers.GetFollowing,
		Authentication: true,
	},
}
