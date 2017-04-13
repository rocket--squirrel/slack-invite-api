package routes

import "net/http"

type Route struct {
	Name           string
	Method         string
	Pattern        string
	HandlerFunc    http.HandlerFunc
	Authenitcation bool
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
		true,
	},
	Route{
		"PostIndex",
		"POST",
		"/",
		PostIndex,
		true,
	},
	Route{
		"PostInvite",
		"POST",
		"/invite",
		PostInvite,
		true,
	},
}
