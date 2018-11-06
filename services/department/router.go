package department

import (
	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
)

func initRoutes(sd *serviceDepartment) *mux.Router {
	routes := common.Routes{
		common.Route{
			Name:        "Information for all the departments",
			Method:      "GET",
			Pattern:     "/info/all",
			HandlerFunc: sd.handlerDepartments,
			PathPrefix:  common.VERSION + "/department",
		},
	}
	return common.NewSubRouter(routes)
}
