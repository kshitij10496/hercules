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

func (s serviceCourse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

func (s serviceCourse) GetDBConnection(ctx context.Context) (*sql.Conn, error) {
	return s.DB.Conn(ctx)
}

func (s serviceCourse) GetURL() string {
	return s.URL
}

func (s serviceCourse) SetDB(db *sql.DB) common.Server {
	s.DB = db
	return s
}

// ServiceCourse represents the course service.
var ServiceCourse = serviceCourse{
	Name: "service-course",
	URL:  "/course",
}
