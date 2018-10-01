package department

import "github.com/kshitij10496/hercules/common"

var Routes = common.Routes{
	common.Route{
		Name:        "Department Info",
		Method:      "GET",
		Pattern:     "/info",
		HandlerFunc: departmentsHandler,
	},
}

var DepartmentService = common.NewService("department-service", "/department", common.NewRouter("/department", Routes))
