package department

import (
	"database/sql"

	"github.com/kshitij10496/hercules/common"
)

// GetDepartments returns the list of departments in IITKGP
func GetDepartments(db *sql.DB) (data common.Departments, err error) {
	// TODO: Fetch data from ERP. Use a JSON as backup.
	rows, err := db.Query("SELECT code, name FROM department")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	departments := common.Departments{}
	for rows.Next() {
		var newDepartment common.Department
		err := rows.Scan(&newDepartment.Code, &newDepartment.Name)
		if err != nil {
			return nil, err
		}
		departments = append(departments, newDepartment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}
