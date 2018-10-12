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
		Name:        "Faculty Info All",
		Method:      "GET",
		Pattern:     "/info/all",
		HandlerFunc: ServiceFaculty.handlerFacultyAll,
		PathPrefix:  common.VERSION + "/faculty",
	},
	common.Route{
		Name:        "Faculty Info Department",
		Method:      "GET",
		Pattern:     "/info/{code}",
		HandlerFunc: ServiceFaculty.handlerFacultyDepartment,
		PathPrefix:  common.VERSION + "/faculty",
	},
	common.Route{
		Name:        "Faculty Timetable",
		Method:      "GET",
		Pattern:     "/timetable",
		HandlerFunc: ServiceFaculty.handlerFacultyTimetable,
		PathPrefix:  common.VERSION + "/faculty",
	},
}

// serviceFaculty implements the server interface
//
type serviceFaculty common.Service

func (s *serviceFaculty) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

func (s *serviceFaculty) GetDBConnection(ctx context.Context) (*sql.Conn, error) {
	return s.DB.Conn(ctx)
}

func (s *serviceFaculty) GetName() string {
	return s.Name
}

func (s *serviceFaculty) GetURL() string {
	return s.URL
}

// SetDB sets the service to use the given DB.
// Note that this function overwrites the current value.
//
func (s *serviceFaculty) ConnectDB(url string) error {
	db, err := sql.Open("postgres", url)
	if err == nil {
		s.DB = db
	}
	return err
}

func (s *serviceFaculty) CloseDB() error {
	return s.DB.Close()
}

// ServiceFaculty represents the course service.
var ServiceFaculty serviceFaculty

func init() {
	ServiceFaculty = serviceFaculty{
		Name:   "service-faculty",
		URL:    "/faculty",
		Router: common.NewSubRouter(Routes),
	}
}
