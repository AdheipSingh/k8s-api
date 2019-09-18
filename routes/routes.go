package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8s-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Deploy",
		"POST",
		"/deploy/prometheus",
		handlers.CreatePromOperators,
	},
	Route{
		"Deploy",
		"POST",
		"/deploy/anchore",
		handlers.CreateAnchoreOperators,
	},
	Route{
		"PutSvc",
		"GET",
		"/getsvc/prometheus",
		handlers.GetSvcProm,
	},
	Route{
		"PutSvc",
		"GET",
		"/getsvc/anchore",
		handlers.GetSvcAnc,
	},
}
