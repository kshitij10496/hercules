package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func departmentsHandler(w http.ResponseWriter, r *http.Request) {
	// response := DepartmentsResponse{}
	departments, err := GetDepartments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	response := DepartmentsResponse(*departments)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func main() {
	port := 8080

	mux := http.DefaultServeMux
	mux.HandleFunc("/departments", departmentsHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Printf("Go to http://127.0.0.1:%v\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
	if err != nil {
		log.Printf("Server cannot be started!\n")
		log.Fatal(err)
	}
}
