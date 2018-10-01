package common

import (
	"github.com/gorilla/mux"
)

const VERSION = "/v1"

type Server struct {
	// db *someDatabase
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}
