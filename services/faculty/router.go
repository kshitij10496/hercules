package faculty

import "github.com/kshitij10496/hercules/common"

var Routes = common.Routes{
	common.Route{
		Name:        "Faculty Info",
		Method:      "GET",
		Pattern:     "/info",
		HandlerFunc: facultyHandler,
		PathPrefix:  "/faculty",
	},
	common.Route{
		Name:        "Faculty Timetable",
		Method:      "POST",
		Pattern:     "/timetable",
		HandlerFunc: facultyTimetableHandler,
		PathPrefix:  "/faculty",
	},
}

type ServiceFaculty common.Service
