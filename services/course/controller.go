package course

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := GetCourses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	response := common.CoursesResponse(courses)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}
