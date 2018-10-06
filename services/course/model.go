package course

import (
	"database/sql"

	"github.com/kshitij10496/hercules/common"
)

var courses = common.Courses{
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

// GetCourses returns the list of courses in IITKGP
func GetCourses(db *sql.DB) (data common.Courses, err error) {
	return courses, nil
}

// GetCourse populates the course with all the relevant information given the course code.
// If no such course exists, an ErrCourseNotFound error is returned.
func GetCourse(db *sql.DB, course *common.Course) error {
	return course.GetCourseInfo(db)
}
