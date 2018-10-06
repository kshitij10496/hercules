package course

import (
	"context"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

func (sc *serviceCourse) handlerCourseInfoAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	conn, err := sc.GetDBConnection(ctx)
	if err != nil {
		// Respond with http.StatusInternalServerError
	}
	// TODO: Catch the error returned while closing connection
	defer conn.Close()

	courses, err := GetCourses(conn)
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

func (sc *serviceCourse) handlerCourseInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var course *common.Course
	err := common.DecodeFromJSON(r, course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request:", err)
	}

	conn, err := sc.GetDBConnection(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	// TODO: Catch the error returned while closing connection
	defer conn.Close()

	err = GetCourse(conn, course)
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
