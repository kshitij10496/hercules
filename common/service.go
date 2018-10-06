package common

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server interface {
	http.Handler
	// GetDBConnection creates a connection to the DB given a context.
	// The context is generated for each request and cleaned up after response.
	//
	GetDBConnection(ctx context.Context) (*sql.Conn, error)

	GetURL() string

	SetDB(db *sql.DB) Server
}

type Service struct {
	Name   string
	URL    string
	DB     *sql.DB
	Router *mux.Router
}

// GetDBConnection creates a connection to the DB given a context.
// The context is generated for each request and cleaned up after response.
//
func (s *Service) GetDBConnection(ctx context.Context) (*sql.Conn, error) {
	return s.DB.Conn(ctx)
}

// GetURL returns the URL of the service.
//
func (s *Service) GetURL() string {
	return s.URL
}

// SetDB sets the service to use the given DB.
// Note that this function overwrites the current value.
//
func (s *Service) SetDB(db *sql.DB) Server {
	s.DB = db
	return s
}

// This makes the Service a http.Handler which can be directly passed to the central router.
// How to track all the services and use them in `server.go`
func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	// Possibly a check on the base URL?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)

}
