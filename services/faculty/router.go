package faculty

import "github.com/kshitij10496/hercules/common"

var Routes = common.Routes{
	common.Route{
		Name:        "Faculty Info",
		Method:      "GET",
		Pattern:     "/info",
		HandlerFunc: facultyHandler,
	},
	common.Route{
		Name:        "Faculty Timetable",
		Method:      "POST",
		Pattern:     "/timetable",
		HandlerFunc: facultyTimetableHandler,
	},
}

var FacultyService = common.NewService("faculty-service", "/faculty", common.NewRouter("/faculty", Routes))
