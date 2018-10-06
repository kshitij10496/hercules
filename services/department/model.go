package department

import (
	"database/sql"

	"github.com/kshitij10496/hercules/common"
)

// GetDepartments returns the list of departments in IITKGP
func GetDepartments(conn *sql.Conn) (data common.Departments, err error) {
	// TODO: Fetch data from ERP. Use a JSON as backup.
	departments := common.Departments{
		common.Department{
			Name: "Mathematics",
			Code: "MA",
		},
		common.Department{
			Name: "Computer Science",
			Code: "CS",
		},
		common.Department{
			Name: "Civil Engineering",
			Code: "CE",
		},
	}
	return departments, nil
}
