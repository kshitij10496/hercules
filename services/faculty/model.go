package faculty

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

// GetFaculty returns the list of faculty members in IITKGP.
func GetFaculty(conn *sql.Conn) (data common.Faculty, err error) {
	// TODO: Fetch data from http://www.iitkgp.ac.in/facultylist
	faculty := common.Faculty{
		common.FacultyMember{
			Name: "Geetanjali Panda",
			Department: common.Department{
				Name: "Mathematics",
				Code: "MA",
			},
			Designation: common.Professor,
		},
		common.FacultyMember{
			Name: "Pratima Panigrahi",
			Department: common.Department{
				Name: "Mathematics",
				Code: "MA",
			},
			Designation: common.Professor,
		},
		common.FacultyMember{
			Name: "Somesh Kumar",
			Department: common.Department{
				Name: "Mathematics",
				Code: "MA",
			},
			Designation: common.Professor,
		},
	}
	return faculty, nil
}

func ReadFaculty(r *http.Request) (facultyMember common.FacultyMember, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&facultyMember)
	return facultyMember, err
}

func GetTimetable(conn *sql.Conn, name string) (data *common.Timetable, err error) {
	timetable := &common.Timetable{
		Monday: []common.Slot{
			common.Slot{
				Course: common.Course{
					Name:    "OCEAN CIRCULATION",
					Credits: 3,
					Code:    "NA61002",
				},
				Timing: common.AM10,
				Room:   common.NC231,
			},
		},
		Tuesday: []common.Slot{
			common.Slot{
				Course: common.Course{
					Name:    "COASTAL ENGINEERING",
					Credits: 3,
					Code:    "NA61001",
				},
				Timing: common.PM5,
				Room:   common.NC142,
			},
		},
		Wednesday: []common.Slot{
			common.Slot{
				Course: common.Course{
					Name:    "OCEAN CIRCULATION",
					Credits: 3,
					Code:    "NA61002",
				},
				Timing: common.PM12,
				Room:   common.NC231,
			},

			common.Slot{
				Course: common.Course{
					Name:    "COASTAL ENGINEERING",
					Credits: 3,
					Code:    "NA61001",
				},
				Timing: common.PM5,
				Room:   common.NC142,
			},
		},
		Thursday: []common.Slot{
			common.Slot{
				Course: common.Course{
					Name:    "OCEAN CIRCULATION",
					Credits: 3,
					Code:    "NA61002",
				},
				Timing: common.PM12,
				Room:   common.NC231,
			},
			common.Slot{
				Course: common.Course{
					Name:    "COASTAL ENGINEERING",
					Credits: 3,
					Code:    "NA61001",
				},
				Timing: common.PM5,
				Room:   common.NC142,
			},
		},
	}
	return timetable, nil
}
