package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/kshitij10496/hercules/common"
	"github.com/kshitij10496/hercules/services/department"
	_ "github.com/lib/pq"
)

func main() {
	var config common.Config
	if err := envconfig.Process("hercules", &config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		envconfig.Usage("hercules", &config)
		os.Exit(1)
	}

	// Create a new router
	mainRouter := mux.NewRouter()

	// Create the subrouter which handles all the API calls
	servicesRouter := mainRouter.PathPrefix(common.VERSION).Subrouter()

	// List all the services
	servers := map[string]common.Server{
		// "service-course":     &course.ServiceCourse,
		"service-department": &department.ServiceDepartment,
		// "service-faculty":    &faculty.ServiceFaculty,
		// "service-migration":  &migration.ServiceMigration,
	}

	// Connect each service with the DB and add them to the subrouters
	for name, server := range servers {
		log.Printf("%s creating...\n", name)

		err := server.ConnectDB(config.Database)
		if err != nil {
			log.Fatalf("Error connecting with DB for %s: %v\n", name, err)
		}

		servicesRouter.PathPrefix(server.GetURL()).Handler(server)

		log.Printf("%s created!\n", name)
	}
	// TODO: Handle services page and home page

	log.Printf("Server starting on %v\n", config.Port)
	if err := http.ListenAndServe(":"+config.Port, mainRouter); err != nil {
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
