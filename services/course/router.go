package course

import (
	"github.com/kshitij10496/hercules/common"
)

var Routes = common.Routes{
	common.Route{
		Name:        "Course Info",
		Method:      "GET",
		Pattern:     "/info/all",
		HandlerFunc: handlerCourseInfoAll,
		PathPrefix:  common.VERSION + "/course",
	},
	common.Route{
		Name:        "Course Info",
		Method:      "POST",
		Pattern:     "/info",
		HandlerFunc: handlerCourseInfo,
		PathPrefix:  common.VERSION + "/course",
	},
}

type ServiceCourse = common.Service
