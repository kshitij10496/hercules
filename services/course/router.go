package course

import (
	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
)

func initRoutes(sc *serviceCourse) *mux.Router {
	routes := common.Routes{
		common.Route{
			Name:        "Courses Timetable",
			Method:      "GET",
			Pattern:     "/timetable/{code}",
			HandlerFunc: sc.handlerCourseTimetable,
			PathPrefix:  common.VERSION + "/course",
		},
		common.Route{
			Name:        "Courses From Department",
			Method:      "GET",
			Pattern:     "/info/department/{code}",
			HandlerFunc: sc.handlerCoursesFromDepartment,
			PathPrefix:  common.VERSION + "/course",
		},
		common.Route{
			Name:        "Courses From Faculty Member",
			Method:      "GET",
			Pattern:     "/info/faculty",
			HandlerFunc: sc.handlerCoursesFromFaculty,
			PathPrefix:  common.VERSION + "/course",
		},
	}
	return common.NewSubRouter(routes)
}
