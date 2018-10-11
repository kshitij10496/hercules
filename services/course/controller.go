package course

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kshitij10496/hercules/common"
)

func (sc *serviceCourse) handlerCourseInfo(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	// var course *common.Course
	// err := common.DecodeFromJSON(r, course)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	log.Println("Bad Request:", err)
	// }

	// // conn, err := sc.GetDBConnection(ctx)
	// // if err != nil {
	// // 	w.WriteHeader(http.StatusInternalServerError)
	// // 	log.Fatal(err)
	// // }

	// // TODO: Catch the error returned while closing connection
	// // defer conn.Close()
	// err = course.GetCourse(sc.DB)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal(err)
	// }
	// common.RespondWithJSON(w, r, http.StatusOK, *course)
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(response)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal(err)
	// }
}

func (sc *serviceCourse) handlerCoursesFromDepartment(w http.ResponseWriter, r *http.Request) {
	// 1. Get Department code from requests URL param
	// 2. DB lookup to fetch course information
	// 3. Return the data
	deptCode, found := mux.Vars(r)["code"]
	if !found {
		http.Error(w, "[required]: Department Code in URL Parameter", http.StatusBadRequest)
		log.Println("Bad Request: No department code provided")
	}

	department := common.Department{Code: deptCode}
	err := department.GetInfo(sc.DB)
	if err != nil {
		// TODO: There could be 2 possible reasons for error here:
		// 		1. Invalid Department code
		//		2. Network issue with DB connection
		http.Error(w, "[invalid]: Invalid Department Code in URL Parameter", http.StatusBadRequest)
		log.Println("Bad Request: Invalid department code provided", err)
	}

	courses, err := getCoursesFromDepartment(sc.DB, department)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Server Error: Cannot fetch courses for department: %+v, err: %v\n", department, err)
	}

	common.RespondWithJSON(w, r, http.StatusOK, courses)
}

func (sc *serviceCourse) handlerCoursesFromFaculty(w http.ResponseWriter, r *http.Request) {

}
