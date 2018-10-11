package course

import (
	"database/sql"
	"log"

	"github.com/kshitij10496/hercules/common"
)

// responseCourse refers to the JSON response body of a course.
// TODO: Explore the use of struct embedding here
type responseCourse struct {
	Code    string               `json:"code"`
	Name    string               `json:"name"`
	Credits int                  `json:"credits"`
	Faculty common.FacultyMember `json:"faculty"`
}

type responseCourses []responseCourse

// GetCourseFromID populates the course with all the relevant information given the course code.
// If no such course exists, an ErrCourseNotFound error is returned.
func GetCourseFromID(db *sql.DB, courseID int) (common.Course, error) {
	return common.Course{}, nil
}

// GetCoursesFromDepartment returns all the courses offered by a given department.
//
func getCoursesFromDepartment(db *sql.DB, department common.Department) (responseCourses, error) {
	query := `SELECT c.code, c.name, c.credits, f.name, d.code, d.name, fd.designation 
				FROM courses c, faculty f, departments d, faculty_designations fd 
				WHERE c.department=$1 
					AND f.id = c.faculty 
					AND f.department=d.id 
					AND f.designation=fd.id`
	rows, err := db.Query(query, department.ID)
	if err != nil {
		return nil, err
	}
	courses := responseCourses{}
	for rows.Next() {
		var newCourse responseCourse
		err = rows.Scan(&newCourse.Code, &newCourse.Name, &newCourse.Credits,
			&newCourse.Faculty.Name, &newCourse.Faculty.Department.Code,
			&newCourse.Faculty.Department.Name, &newCourse.Faculty.Designation)
		if err != nil {
			log.Printf("Error while scanning department courses: %v, err: %v\n", department.Code, err)
		}
		courses = append(courses, newCourse)
	}
	return courses, nil
}

// GetCoursesFromFaculty returns all the courses offered by the given faculty member.
//
func GetCoursesFromFaculty(db *sql.DB, facultyMember common.FacultyMember) (common.Courses, error) {

	return nil, nil
}
