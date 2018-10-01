package common

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(serviceURL string, routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		log.Println(route.Name)
		router.
			Methods(route.Method).
			Path(VERSION + serviceURL + route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
