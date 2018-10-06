package faculty

import (
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func (sf *serviceFaculty) facultyHandler(w http.ResponseWriter, r *http.Request) {
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

func (sf *serviceFaculty) facultyTimetableHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	faculty, err := ReadFaculty(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	// conn, err := sf.GetDBConnection(ctx)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal("Error connecting to DB:", err)
	// }

	timetable, err := GetTimetable(sf.DB, faculty.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	response := common.TimeTableResponse{Timetable: *timetable}
	common.RespondWithJSON(w, r, http.StatusOK, response)
}
