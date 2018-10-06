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

	log.Println("service-course creating...")
	serviceCourse := course.ServiceCourse(*common.NewService("service-course", "/course", course.Routes))
	log.Println("service-course created")
	fmt.Println()

	log.Println("service-department creating...")
	serviceDepartment := department.ServiceDepartment(*common.NewService("service-department", "/department", department.Routes))
	log.Println("service-department created")
	fmt.Println()

	log.Println("service-faculty creating...")
	serviceFaculty := faculty.ServiceFaculty(*common.NewService("service-faculty", "/faculty", faculty.Routes))
	log.Println("service-faculty created")
	fmt.Println()

	mainRouter := mux.NewRouter()
	servicesRouter := mainRouter.PathPrefix(common.VERSION).Subrouter()
	log.Println("After adding subrouters")
	services := []common.Service{serviceCourse, serviceDepartment, serviceFaculty}
	for _, service := range services {
		servicesRouter.PathPrefix(service.URL).Handler(service)
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

	log.Printf("Go to http://127.0.0.1:%v\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mainRouter)
	if err != nil {
		log.Printf("Server cannot be started!\n")
		log.Fatal(err)
	}
}
