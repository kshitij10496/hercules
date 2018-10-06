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
		PathPrefix:  "/course",
	},
	common.Route{
		Name:        "Course Info",
		Method:      "POST",
		Pattern:     "/info",
		HandlerFunc: handlerCourseInfo,
		PathPrefix:  "/course",
	},
}

type ServiceCourse common.Service

// func init() {
// 	sc := ServiceCourse(*common.NewService("service-course", "/course", Routes))
// 	log.Println("service-course created")
// 	log.Println(sc.Name, sc.URL)
// }
