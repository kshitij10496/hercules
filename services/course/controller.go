package course

import (
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func handlerCourseInfoAll(w http.ResponseWriter, r *http.Request) {
	courses, err := GetCourses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	common.RespondWithJSON(w, r, http.StatusOK, courses)
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(response)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal(err)
	// }
}

func handlerCourseInfo(w http.ResponseWriter, r *http.Request) {
	var course *common.Course
	common.DecodeFromJSON(r, course)
	err := GetCourse(course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	common.RespondWithJSON(w, r, http.StatusOK, *course)
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(response)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal(err)
	// }
}
