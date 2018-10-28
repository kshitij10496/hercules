package course

import (
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

// serviceCourse implements the server interface
type serviceCourse struct {
	common.Service
	DB courseDataSource
}

func (s *serviceCourse) GetName() string {
	return s.Name
}

func (s *serviceCourse) GetURL() string {
	return s.URL
}

func (s *serviceCourse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

func (s *serviceCourse) ConnectDB(url string) error {
	return s.DB.ConnectDS(url)
}

func (s *serviceCourse) CloseDB() error {
	return s.DB.CloseDS()
}

// NewServiceCourse instantiates a new course-service with a real database connection.
func NewServiceCourse() *serviceCourse {
	ServiceCourse := &serviceCourse{
		common.Service{
			Name: "service-course",
			URL:  "/course",
		},
		newRealDataSource(),
	}
	ServiceCourse.Router = initRoutes(ServiceCourse)
	return ServiceCourse
}
