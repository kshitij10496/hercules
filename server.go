package main

import (
	"database/sql"
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

	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Error connecting to the DB:", err)
	// }
	// defer db.Close()

	// log.Println("service-course creating...")
	// course.ServiceCourse.DB = db
	// log.Println("service-course created")
	// fmt.Println()

	// log.Println("service-department creating...")
	// department.ServiceDepartment.DB = db
	// log.Println("service-department created")
	// fmt.Println()

	// log.Println("service-faculty creating...")
	// faculty.ServiceFaculty.DB = db
	// log.Println("service-faculty created")
	// fmt.Println()

	mainRouter := mux.NewRouter()
	servicesRouter := mainRouter.PathPrefix(common.VERSION).Subrouter()
	log.Println("After adding subrouters")
	servers := []common.Server{
		course.ServiceCourse,
		department.ServiceDepartment,
		faculty.ServiceFaculty,
	}
	for _, server := range servers {
		log.Printf("%s creating...\n", server.GetURL())
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal("Error connecting to the DB:", err)
		}
		defer db.Close()

		server.SetDB(db)
		servicesRouter.PathPrefix(server.GetURL()).Handler(server)
		log.Println("%s created\n", server.GetURL())
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
