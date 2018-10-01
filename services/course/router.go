package course

import "github.com/kshitij10496/hercules/common"

var Routes = common.Routes{
	common.Route{
		Name:        "Course Info",
		Method:      "GET",
		Pattern:     "/info",
		HandlerFunc: coursesHandler,
	},
}

var CourseService = common.NewService("course-service", "/course", common.NewRouter("/course", Routes))
