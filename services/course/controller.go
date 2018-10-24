package course

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
)

func (sc *serviceCourse) handlerCourseTimetable(w http.ResponseWriter, r *http.Request) {
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
	courseCode, found := mux.Vars(r)["code"]
	if !found {
		http.Error(w, "[required]: Course Code in URL Parameter", http.StatusBadRequest)
		log.Println("Bad Request: No course code provided")
		return
	}

	course := common.Course{Code: courseCode}
	err := sc.DB.GetCourseInfo(&course)
	if err != nil {
		http.Error(w, "[invalid]: Invalid Department Code in URL Parameter", http.StatusBadRequest)
		log.Println("Bad Request: Invalid department code provided", err)
		return
	}

	timetable, err := sc.DB.GetCourseTimetable(course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Bad Request: Invalid course code provided %s: %s\n", course.Code, err)
		return
	}

	common.RespondWithJSON(w, r, http.StatusOK, timetable)
}

func (sc *serviceCourse) handlerCoursesFromDepartment(w http.ResponseWriter, r *http.Request) {
	// 1. Get Department code from requests URL param
	// 2. DB lookup to fetch course information
	// 3. Return the data
	deptCode, found := mux.Vars(r)["code"]
	if !found || deptCode == "" {
		http.Error(w, "[required]: Department Code in URL Parameter", http.StatusBadRequest)
		log.Println("Bad Request: No department code provided")
		return
	}

	department := common.Department{Code: deptCode}
	err := sc.DB.GetDepartmentInfo(&department)
	if err != nil {
		// TODO: There could be 2 possible reasons for error here:
		// 		1. Invalid Department code
		//		2. Network issue with DB connection
		http.Error(w, "[invalid]: Invalid Department Code in URL Parameter", http.StatusBadRequest)
		log.Println("Bad Request: Invalid department code provided", err)
		return
	}

	courses, err := sc.DB.GetCoursesFromDepartment(department)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Server Error: Cannot fetch courses for department: %+v, err: %v\n", department, err)
		return
	}

	common.RespondWithJSON(w, r, http.StatusOK, courses)
}

func (sc *serviceCourse) handlerCoursesFromFaculty(w http.ResponseWriter, r *http.Request) {
	// 1. Get faculty member's name and their department from URL query parameter
	// 2. Query DB
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

	courses, err := sc.DB.GetCoursesFromFaculty(facultyMember)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Server Error: Cannot fetch courses for faculty: %+v, err: %v\n", facultyMember, err)
		return
	}

	common.RespondWithJSON(w, r, http.StatusOK, courses)
}
