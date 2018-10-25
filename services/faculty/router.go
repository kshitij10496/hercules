package faculty

import (
	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
)

func initRoutes(sf *serviceFaculty) *mux.Router {
	routes := common.Routes{
		common.Route{
			Name:        "Faculty Info All",
			Method:      "GET",
			Pattern:     "/info/all",
			HandlerFunc: sf.handlerFacultyAll,
			PathPrefix:  common.VERSION + "/faculty",
		},
		common.Route{
			Name:        "Faculty Info Department",
			Method:      "GET",
			Pattern:     "/info/{code}",
			HandlerFunc: sf.handlerFacultyDepartment,
			PathPrefix:  common.VERSION + "/faculty",
		},
		common.Route{
			Name:        "Faculty Timetable",
			Method:      "GET",
			Pattern:     "/timetable",
			HandlerFunc: sf.handlerFacultyTimetable,
			PathPrefix:  common.VERSION + "/faculty",
		},
	}
	return common.NewSubRouter(routes)
}
