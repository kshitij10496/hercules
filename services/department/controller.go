package department

import (
	"context"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func (sd *serviceDepartment) departmentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	conn, err := sd.GetDBConnection(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error connecting to DB:", err)
	}

	departments, err := GetDepartments(conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	common.RespondWithJSON(w, r, http.StatusOK, departments)
}
