package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	mux := http.DefaultServeMux
	mux.HandleFunc("/departments", departmentsHandler)
	mux.HandleFunc("/courses", coursesHandler)
	mux.HandleFunc("/faculty", facultyHandler)
	mux.HandleFunc("/faculty/timetable", facultyTimetableHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Printf("Go to http://127.0.0.1:%v\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
	if err != nil {
		log.Printf("Server cannot be started!\n")
		log.Fatal(err)
	}
}
