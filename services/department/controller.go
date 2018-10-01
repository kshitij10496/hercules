package department

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func departmentsHandler(w http.ResponseWriter, r *http.Request) {
	// response := DepartmentsResponse{}
	departments, err := GetDepartments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	response := common.DepartmentsResponse(departments)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}
