package faculty

import (
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func facultyHandler(w http.ResponseWriter, r *http.Request) {
	faculty, err := GetFaculty()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	common.RespondWithJSON(w, r, http.StatusOK, faculty)
}

func facultyTimetableHandler(w http.ResponseWriter, r *http.Request) {
	faculty, err := ReadFaculty(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	timetable, err := GetTimetable(faculty.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	response := common.TimeTableResponse{Timetable: *timetable}
	common.RespondWithJSON(w, r, http.StatusOK, response)
}
