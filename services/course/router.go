package course

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

var Routes = common.Routes{
	common.Route{
		Name:        "Course Info",
		Method:      "GET",
		Pattern:     "/info/all",
		HandlerFunc: ServiceCourse.handlerCourseInfoAll,
		PathPrefix:  common.VERSION + "/course",
	},
	common.Route{
		Name:        "Course Info",
		Method:      "POST",
		Pattern:     "/info",
		HandlerFunc: ServiceCourse.handlerCourseInfo,
		PathPrefix:  common.VERSION + "/course",
	},
}

// serviceCourse implements the server interface
//
type serviceCourse common.Service

func (s *serviceCourse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

func (s *serviceCourse) GetDBConnection(ctx context.Context) (*sql.Conn, error) {
	return s.DB.Conn(ctx)
}

func (s *serviceCourse) GetName() string {
	return s.Name
}

func (s *serviceCourse) GetURL() string {
	return s.URL
}

// SetDB sets the service to use the given DB.
// Note that this function overwrites the current value.
//
func (s *serviceCourse) ConnectDB(url string) error {
	db, err := sql.Open("postgres", url)
	if err == nil {
		s.DB = db
	}
	return err
}

// ServiceCourse represents the course service.
var ServiceCourse serviceCourse

func init() {
	ServiceCourse = serviceCourse{
		Name:   "service-course",
		URL:    "/course",
		Router: common.NewSubRouter(Routes),
	}
}
