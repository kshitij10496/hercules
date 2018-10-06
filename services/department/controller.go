package department

import (
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
	common.RespondWithJSON(w, r, http.StatusOK, departments)
}
