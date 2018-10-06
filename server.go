package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/kshitij10496/hercules/common"
	"github.com/kshitij10496/hercules/services/course"
	"github.com/kshitij10496/hercules/services/department"
	"github.com/kshitij10496/hercules/services/faculty"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT environment variable")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	// db, err := sql.Open("postgres", )
	// if err != nil {
	// 	log.Fatal("Error connecting to the DB:", err)
	// }
	// defer db.Close()

	// log.Println("service-course creating...")
	// course.ServiceCourse.DB = db
	// course.ServiceCourse.Router = common.NewSubRouter(course.Routes)
	// log.Println("service-course created")

	// log.Println("service-department creating...")
	// department.ServiceDepartment.DB = db

	// log.Println("service-department created")

	// log.Println("service-faculty creating...")
	// faculty.ServiceFaculty.DB = db
	// faculty.ServiceFaculty.Router = common.NewSubRouter(faculty.Routes)

	mainRouter := mux.NewRouter()
	servicesRouter := mainRouter.PathPrefix(common.VERSION).Subrouter()
	log.Println("After adding subrouters")
	servers := []common.Server{
		&course.ServiceCourse,
		&department.ServiceDepartment,
		&faculty.ServiceFaculty,
	}
	for _, server := range servers {
		log.Printf("%s creating...\n", server.GetURL())
		err := server.ConnectDB(databaseURL)
		if err != nil {
			log.Fatal("Error connecting with DB for %s:", server.GetName(), err)
		}
		servicesRouter.PathPrefix(server.GetURL()).Handler(server)
		log.Printf("%s created!\n", server.GetURL())
	}
	// TODO: Handle services page and home page

	mainRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})

	log.Printf("Server starting on %v\n", port)
	if err := http.ListenAndServe(":"+port, mainRouter); err != nil {
		log.Printf("Server cannot be started!\n")
		log.Fatal(err)
	}
}
