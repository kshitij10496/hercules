package common

import "github.com/gorilla/mux"

type Service struct {
	Name   string
	URL    string
	Router *mux.Router
}

func NewService(name string, url string, router *mux.Router) *Service {
	return &Service{
		Name:   name,
		URL:    url,
		Router: router,
	}
}
