package faculty

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

// GetFaculty returns the list of faculty members in IITKGP.
func GetFaculty(db *sql.DB) (data common.Faculty, err error) {
	// TODO: Fetch data from http://www.iitkgp.ac.in/facultylist
	query := "SELECT f.name, fd.designation, d.code, d.name FROM departments d, faculty_designations fd, faculty f WHERE f.designation=fd.id AND f.department=d.id;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faculty common.Faculty
	for rows.Next() {
		var name string
		var department common.Department
		var designation common.FacultyDesignation

		err := rows.Scan(&name, &designation, &department.Code, &department.Name)
		if err != nil {
			return nil, err
		}

		newFacultyMember := common.FacultyMember{
			Name:        name,
			Designation: designation,
			Department:  department,
		}

		faculty = append(faculty, newFacultyMember)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return faculty, nil
}

func GetFacultyDepartment(db *sql.DB, deptCode string) (common.Faculty, error) {
	query := `SELECT f.name, fd.designation, d.code, d.name FROM faculty f, departments d, faculty_designations fd WHERE d.code=$1 AND f.department=d.id AND f.designation=fd.id`
	rows, err := db.Query(query, deptCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faculty common.Faculty
	for rows.Next() {
		var newFacultyMember common.FacultyMember
		err := rows.Scan(&newFacultyMember.Name, &newFacultyMember.Designation, &newFacultyMember.Department.Code, &newFacultyMember.Department.Name)
		if err != nil {
			log.Printf("Error fetching department faculty: %v, err: %v", deptCode, err)
			continue
		}

		faculty = append(faculty, newFacultyMember)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return faculty, nil
}

func ReadFaculty(r *http.Request) (facultyMember common.FacultyMember, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&facultyMember)
	return facultyMember, err
}

func GetTimetable(db *sql.DB, name string) (data *common.Timetable, err error) {
	// timetable := &common.Timetable{
	// 	Monday: []common.Slot{
	// 		common.Slot{
	// 			Course: common.Course{
	// 				Name:    "OCEAN CIRCULATION",
	// 				Credits: 3,
	// 				Code:    "NA61002",
	// 			},
	// 			Timing: common.AM10,
	// 			Room:   common.NC231,
	// 		},
	// 	},
	// 	Tuesday: []common.Slot{
	// 		common.Slot{
	// 			Course: common.Course{
	// 				Name:    "COASTAL ENGINEERING",
	// 				Credits: 3,
	// 				Code:    "NA61001",
	// 			},
	// 			Timing: common.PM5,
	// 			Room:   common.NC142,
	// 		},
	// 	},
	// 	Wednesday: []common.Slot{
	// 		common.Slot{
	// 			Course: common.Course{
	// 				Name:    "OCEAN CIRCULATION",
	// 				Credits: 3,
	// 				Code:    "NA61002",
	// 			},
	// 			Timing: common.PM12,
	// 			Room:   common.NC231,
	// 		},

	// 		common.Slot{
	// 			Course: common.Course{
	// 				Name:    "COASTAL ENGINEERING",
	// 				Credits: 3,
	// 				Code:    "NA61001",
	// 			},
	// 			Timing: common.PM5,
	// 			Room:   common.NC142,
	// 		},
	// 	},
	// 	Thursday: []common.Slot{
	// 		common.Slot{
	// 			Course: common.Course{
	// 				Name:    "OCEAN CIRCULATION",
	// 				Credits: 3,
	// 				Code:    "NA61002",
	// 			},
	// 			Timing: common.PM12,
	// 			Room:   common.NC231,
	// 		},
	// 		common.Slot{
	// 			Course: common.Course{
	// 				Name:    "COASTAL ENGINEERING",
	// 				Credits: 3,
	// 				Code:    "NA61001",
	// 			},
	// 			Timing: common.PM5,
	// 			Room:   common.NC142,
	// 		},
	// 	},
	// }
	return nil, nil
}
