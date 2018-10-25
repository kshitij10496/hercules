package faculty

import (
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

// serviceFaculty implements the server interface
//
type serviceFaculty struct {
	common.Service
	DB facultyDataSource
}

func (s *serviceFaculty) GetName() string {
	return s.Name
}

func (s *serviceFaculty) GetURL() string {
	return s.URL
}

func (s *serviceFaculty) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

// SetDB sets the service to use the given DB.
// Note that this function overwrites the current value.
//
func (s *serviceFaculty) ConnectDB(url string) error {
	return s.DB.ConnectDS(url)
}

func (s *serviceFaculty) CloseDB() error {
	return s.DB.CloseDS()
}

func NewServiceFaculty() *serviceFaculty {
	ServiceFaculty := &serviceFaculty{
		common.Service{
			Name: "service-faculty",
			URL:  "/faculty",
		},
		NewRealFacultyDataSource(),
	}
	ServiceFaculty.Router = initRoutes(ServiceFaculty)
	return ServiceFaculty
}
