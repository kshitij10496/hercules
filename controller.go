package main

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
