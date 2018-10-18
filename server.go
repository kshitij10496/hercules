package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"

	"github.com/kshitij10496/hercules/common"
	"github.com/kshitij10496/hercules/services/course"
	"github.com/kshitij10496/hercules/services/department"
	"github.com/kshitij10496/hercules/services/faculty"
	"github.com/kshitij10496/hercules/services/migration"
)

type Specification struct {
	Port        string `required:"true"`
	DatabaseUrl string `required:"true"`
}

func main() {
	// Get $PORT and $DATABASE_URL from env
	var s Specification
	err := envconfig.Process("hercules", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	port := s.Port
	databaseURL := s.DatabaseUrl

	log.Print("From envconfig:")
	log.Print(s.Port, s.DatabaseUrl)

	// Create a new router
	mainRouter := mux.NewRouter()

	// Create the subrouter which handles all the API calls
	servicesRouter := mainRouter.PathPrefix(common.VERSION).Subrouter()

	// List all the services
	servers := map[string]common.Server{
		"service-course":     &course.ServiceCourse,
		"service-department": &department.ServiceDepartment,
		"service-faculty":    &faculty.ServiceFaculty,
		"service-migration":  &migration.ServiceMigration,
	}

	// Connect each service with the DB and add them to the subrouters
	for name, server := range servers {
		log.Printf("%s creating...\n", name)

		err := server.ConnectDB(databaseURL)
		if err != nil {
			log.Fatalf("Error connecting with DB for %s: %v\n", name, err)
		}

		servicesRouter.PathPrefix(server.GetURL()).Handler(server)

		log.Printf("%s created!\n", name)
	}
	// TODO: Handle services page and home page

	log.Printf("Server starting on %v\n", port)
	if err := http.ListenAndServe(":"+port, mainRouter); err != nil {
		for name, server := range servers {
			log.Printf("%s closing...\n", name)

			err := server.CloseDB()
			if err != nil {
				log.Fatalf("Error closing DB for %s: %v\n", name, err)
			}

			log.Printf("%s closed!\n", name)
		}
		log.Printf("Server cannot be started!\n")
		log.Fatal(err)
	}
}
