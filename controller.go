package main

import (
	"encoding/json"
	"net/http"
)

// GetDepartments returns the list of departments in IITKGP
func GetDepartments() (data *[]Department, err error) {
	// TODO: Fetch data from ERP. Use a JSON as backup.
	departments := &[]Department{
		Department{
			Name: "Mathematics",
			Code: "MA",
		},
		Department{
			Name: "Computer Science",
			Code: "CS",
		},
		Department{
			Name: "Civil Engineering",
			Code: "CE",
		},
	}
	return departments, nil
}

// GetCourses returns the list of courses in IITKGP
func GetCourses() (data *[]Course, err error) {
	courses := &[]Course{
		Course{
			Name:    "VLSI TECHNOLOGY",
			Credits: 3,
			Code:    "EC60289",
		},
		Course{
			Name:    "OCEAN CIRCULATION",
			Credits: 3,
			Code:    "NA61002",
		},
		Course{
			Name:    "COASTAL ENGINEERING",
			Credits: 3,
			Code:    "NA61001",
		},
	}
	return courses, nil
}

// GetFaculty returns the list of faculty members in IITKGP.
func GetFaculty() (data *[]Faculty, err error) {
	// TODO: Fetch data from http://www.iitkgp.ac.in/facultylist
	faculty := &[]Faculty{
		Faculty{
			Name: "Geetanjali Panda",
			Department: Department{
				Name: "Mathematics",
				Code: "MA",
			},
			Designation: Professor,
		},
		Faculty{
			Name: "Pratima Panigrahi",
			Department: Department{
				Name: "Mathematics",
				Code: "MA",
			},
			Designation: Professor,
		},
		Faculty{
			Name: "Somesh Kumar",
			Department: Department{
				Name: "Mathematics",
				Code: "MA",
			},
			Designation: Professor,
		},
	}
	return faculty, nil
}

func ReadFaculty(r *http.Request) (faculty Faculty, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&faculty)
	return faculty, err
}

func GetTimetable(name string) (data *Timetable, err error) {
	timetable := &Timetable{
		Monday: []Slot{
			Slot{
				Course: Course{
					Name:    "OCEAN CIRCULATION",
					Credits: 3,
					Code:    "NA61002",
				},
				Timing: AM10,
				Room:   NC231,
			},
		},
		Tuesday: []Slot{
			Slot{
				Course: Course{
					Name:    "COASTAL ENGINEERING",
					Credits: 3,
					Code:    "NA61001",
				},
				Timing: PM5,
				Room:   NC142,
			},
		},
		Wednesday: []Slot{
			Slot{
				Course: Course{
					Name:    "OCEAN CIRCULATION",
					Credits: 3,
					Code:    "NA61002",
				},
				Timing: PM12,
				Room:   NC231,
			},

			Slot{
				Course: Course{
					Name:    "COASTAL ENGINEERING",
					Credits: 3,
					Code:    "NA61001",
				},
				Timing: PM5,
				Room:   NC142,
			},
		},
		Thursday: []Slot{
			Slot{
				Course: Course{
					Name:    "OCEAN CIRCULATION",
					Credits: 3,
					Code:    "NA61002",
				},
				Timing: PM12,
				Room:   NC231,
			},
			Slot{
				Course: Course{
					Name:    "COASTAL ENGINEERING",
					Credits: 3,
					Code:    "NA61001",
				},
				Timing: PM5,
				Room:   NC142,
			},
		},
	}
	return timetable, nil
}
