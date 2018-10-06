package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kshitij10496/hercules/common"
	"github.com/kshitij10496/hercules/services/course"
	"github.com/kshitij10496/hercules/services/department"
	"github.com/kshitij10496/hercules/services/faculty"
)

func main() {
	port := 8080

	mainRouter := mux.NewRouter()

	log.Println("service-course creating...")
	serviceCourse := course.ServiceCourse(*common.NewService("service-course", "/course", course.Routes))
	log.Println("service-course created")
	log.Println(serviceCourse.Name, serviceCourse.URL)
	log.Println()

	log.Println("service-department creating...")
	serviceDepartment := department.ServiceDepartment(*common.NewService("service-department", "/department", department.Routes))
	log.Println("service-department created")
	log.Println(serviceDepartment.Name, serviceDepartment.URL)
	log.Println()

	log.Println("service-faculty creating...")
	serviceFaculty := faculty.ServiceFaculty(*common.NewService("service-faculty", "/faculty", faculty.Routes))
	log.Println("service-faculty created")
	log.Println(serviceFaculty.Name, serviceFaculty.URL)
	log.Println()
	mainRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})

	mainRouter.Handle("/", serviceCourse.Router)
	mainRouter.Handle("/", serviceDepartment.Router)
	mainRouter.Handle("/", serviceFaculty.Router)

	log.Println("After adding subrouters")
	mainRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})

	log.Printf("Go to http://127.0.0.1:%v\n", port)

	http.Handle("/", mainRouter)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mainRouter)
	if err != nil {
		log.Printf("Server cannot be started!\n")
		log.Fatal(err)
	}
}
