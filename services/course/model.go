package course

import "github.com/kshitij10496/hercules/common"

// GetCourses returns the list of courses in IITKGP
func GetCourses() (data []common.Course, err error) {
	courses := []common.Course{
		common.Course{
			Name:    "VLSI TECHNOLOGY",
			Credits: 3,
			Code:    "EC60289",
		},
		common.Course{
			Name:    "OCEAN CIRCULATION",
			Credits: 3,
			Code:    "NA61002",
		},
		common.Course{
			Name:    "COASTAL ENGINEERING",
			Credits: 3,
			Code:    "NA61001",
		},
	}
	return courses, nil
}
