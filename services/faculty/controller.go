package faculty

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
)

func (sf *serviceFaculty) handlerFacultyAll(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	// conn, err := sf.GetDBConnection(ctx)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal("Error connecting to DB:", err)
	// }
	faculty, err := GetFaculty(sf.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	common.RespondWithJSON(w, r, http.StatusOK, faculty)
}

func (sf *serviceFaculty) handlerFacultyDepartment(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	// conn, err := sf.GetDBConnection(ctx)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal("Error connecting to DB:", err)
	// }
	deptCode, found := mux.Vars(r)["code"]
	if !found {
		http.Error(w, "[required]: Department code in URL parameter", http.StatusBadRequest)
		return
	}
	faculty, err := GetFacultyDepartment(sf.DB, deptCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	common.RespondWithJSON(w, r, http.StatusOK, faculty)
}

func (sf *serviceFaculty) handlerFacultyTimetable(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	values := r.URL.Query()
	names, found := values["name"]
	if !found || len(names) != 1 {
		http.Error(w, "[required]: name as a query parameter", http.StatusBadRequest)
		log.Println("Bad Request: No faculty name provided")
		return
	}
	deptCodes, found := values["dept"]
	if !found || len(deptCodes) != 1 {
		http.Error(w, "[required]: dept as query parameter", http.StatusBadRequest)
		log.Println("Bad Request: No faculty department provided")
		return
	}

	facultyMember := common.FacultyMember{
		Name: names[0],
		Department: common.Department{
			Code: deptCodes[0],
		},
	}
	// conn, err := sf.GetDBConnection(ctx)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal("Error connecting to DB:", err)
	// }

	timetable, err := GetTimetable(sf.DB, facultyMember)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := timetable
	common.RespondWithJSON(w, r, http.StatusOK, response)
}
