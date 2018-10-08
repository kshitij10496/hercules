package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/kshitij10496/hercules/common"
	"github.com/kshitij10496/hercules/services/course"
	"github.com/kshitij10496/hercules/services/department"
	"github.com/kshitij10496/hercules/services/faculty"
	"github.com/kshitij10496/hercules/services/migration"
)

func main() {
	// Grab $PORT from env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing: PORT environment variable")
	}

	// Grab $DATABASE_URL from env
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("Missing: DATABASE_URL environment variable")
	}

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
