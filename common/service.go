package common

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	Name   string
	URL    string
	Router *mux.Router
}

func NewService(name string, url string, routes Routes) *Service {
	return &Service{
		Name:   name,
		URL:    url,
		Router: NewSubRouter(routes),
	}
}

// This makes the Service a http.Handler which can be directly passed to the central router.
// How to track all the services and use them in `server.go`
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	// Possibly a check on the base URL?
	log.Printf("%s is running: %v\n", s.Name, s.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("%s is stopping: %v\n", s.Name, s.URL)

}
