package common

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents the fundamental components of the API.
// Every service needs to implement this interface.
//
type Server interface {
	http.Handler

	HerculesDB

	GetName() string

	GetURL() string
}

type HerculesDB interface {
	ConnectDB(databaseURL string) error

	CloseDB() error
}

// Service is the implementation of the Server interface.
//
type Service struct {
	Name   string
	URL    string
	Router *mux.Router
}

// SAMPLE SERVER
//
// // Service is a http.Handler which can be directly passed to the central router.
// // How to track all the services and use them in `server.go`
// func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// TODO: What should go in here?
// 	// Possibly a check on the base URL?
// 	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
// 	s.Router.ServeHTTP(w, r)
// 	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
// }
//
// func (s *Service) GetName() string {
// 	return s.Name
// }
//
// // GetURL returns the URL of the service.
// //
// func (s *Service) GetURL() string {
// 	return s.URL
// }
//
// // SetDB sets the service to use the given DB.
// // Note that this function overwrites the current value.
// //
// func (s *Service) ConnectDB(url string) error {
// 	db, err := sql.Open("postgres", url)
// 	if err == nil {
// 		s.DB = db
// 	}
// 	return err
// }
//
// func (s *Service) CloseDB() error {
// 	return s.DB.Close()
// }
//
// // GetDBConnection creates a connection to the DB given a context.
// // The context is generated for each request and cleaned up after response.
// //
// func (s *Service) GetDBConnection(ctx context.Context) (*sql.Conn, error) {
// 	return s.DB.Conn(ctx)
// }
