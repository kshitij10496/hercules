package department

import (
	"github.com/kshitij10496/hercules/common"
)

var Routes = common.Routes{
	common.Route{
		Name:        "Information for all the departments",
		Method:      "GET",
		Pattern:     "/info/all",
		HandlerFunc: ServiceDepartment.handlerDepartments,
		PathPrefix:  common.VERSION + "/department",
	},
}
