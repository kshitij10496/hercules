package department

import (
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

// serviceDepartment implements the server interface
//
type serviceDepartment struct {
	common.Service
	DB departmentsDatasource
}

func (s *serviceDepartment) GetName() string {
	return s.Name
}

func (s *serviceDepartment) GetURL() string {
	return s.URL
}

func (s *serviceDepartment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

func (s *serviceDepartment) ConnectDB(url string) error {
	return s.DB.ConnectDS(url)
}

func (s *serviceDepartment) CloseDB() error {
	return s.DB.CloseDS()
}

// NewServiceDepartment instantiates a new department-service with a real database connection.
func NewServiceDepartment() *serviceDepartment {
	ServiceDepartment := &serviceDepartment{
		common.Service{
			Name: "service-department",
			URL:  "/department",
		},
		newRealDataSource(),
	}
	ServiceDepartment.Router = initRoutes(ServiceDepartment)
	return ServiceDepartment
}
