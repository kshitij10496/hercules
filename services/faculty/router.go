package faculty

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

var Routes = common.Routes{
	common.Route{
		Name:        "Faculty Info",
		Method:      "GET",
		Pattern:     "/info",
		HandlerFunc: ServiceFaculty.facultyHandler,
		PathPrefix:  common.VERSION + "/faculty",
	},
	common.Route{
		Name:        "Faculty Timetable",
		Method:      "POST",
		Pattern:     "/timetable",
		HandlerFunc: ServiceFaculty.facultyTimetableHandler,
		PathPrefix:  common.VERSION + "/faculty",
	},
}

// serviceFaculty implements the server interface
//
type serviceFaculty common.Service

func (s serviceFaculty) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

func (s serviceFaculty) GetDBConnection(ctx context.Context) (*sql.Conn, error) {
	return s.DB.Conn(ctx)
}

func (s serviceFaculty) GetURL() string {
	return s.URL
}

func (s serviceFaculty) SetDB(db *sql.DB) common.Server {
	s.DB = db
	return s
}

// ServiceFaculty represents the course service.
var ServiceFaculty serviceFaculty

// Initialise the service with no DB.
func init() {
	ServiceFaculty = serviceFaculty{
		Name:   "service-course",
		URL:    "/course",
		DB:     nil,
		Router: common.NewSubRouter(Routes),
	}
}
