package main

// GetDepartments returns the list of departments in IITKGP
func GetDepartments() (data *[]Department, err error) {
	department := &[]Department{
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
	return department, nil
}
