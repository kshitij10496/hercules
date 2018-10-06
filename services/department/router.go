package department

import "github.com/kshitij10496/hercules/common"

var Routes = common.Routes{
	common.Route{
		Name:        "Department Info",
		Method:      "GET",
		Pattern:     "/info",
		HandlerFunc: departmentsHandler,
		PathPrefix:  "/department",
	},
}

type ServiceDepartment common.Service
