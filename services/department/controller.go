package department

import (
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func (sd *serviceDepartment) handlerDepartments(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	// conn, err := sd.GetDBConnection(ctx)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal("Error connecting to DB:", err)
	// }
	if sd.DB == nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("cannot setup env")
		return
	}
	departments, err := sd.DB.GetDepartments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	common.RespondWithJSON(w, r, http.StatusOK, departments)
}
