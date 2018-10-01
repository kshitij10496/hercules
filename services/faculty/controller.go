package faculty

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func facultyHandler(w http.ResponseWriter, r *http.Request) {
	faculty, err := GetFaculty()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	response := common.FacultyResponse(faculty)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func facultyTimetableHandler(w http.ResponseWriter, r *http.Request) {
	faculty, err := ReadFaculty(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	timetable, err := GetTimetable(faculty.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	response := common.TimeTableResponse{Timetable: *timetable}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}
