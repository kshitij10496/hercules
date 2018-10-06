package common

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const VERSION = "v1"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	PathPrefix  string
}

type Routes []Route

func NewSubRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// subrouter := server.Router.Subrouter()
	for _, route := range routes {
		path := "/" + VERSION + route.PathPrefix + route.Pattern
		router.
			Methods(route.Method).
			Path(path).
			Name(route.Name).
			Handler(route.HandlerFunc)

		log.Println("created route:", route.Method, path)
	}
	return router
}
