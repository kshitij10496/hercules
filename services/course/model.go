package course

import (
	"database/sql"
	"log"

	"github.com/kshitij10496/hercules/common"
)

// responseCourse refers to the JSON response body of a course.
// TODO: Explore the use of struct embedding here
type responseCourse struct {
	Code       string             `json:"code"`
	Name       string             `json:"name"`
	Credits    int                `json:"credits"`
	Department *common.Department `json:"department,omitempty"`
}

func newResponseCourse() *responseCourse {
	return &responseCourse{Department: &common.Department{}}
}

type responseCourses []responseCourse

// GetCoursesFromDepartment returns all the courses offered by a given department.
//
func getCoursesFromDepartment(db *sql.DB, department common.Department) (responseCourses, error) {
	query := `SELECT c.code, c.name, c.credits FROM courses c, departments d
				WHERE c.department=$1 AND d.id=c.department`
	rows, err := db.Query(query, department.ID)
	if err != nil {
		return nil, err
	}
	courses := responseCourses{}
	for rows.Next() {
		var newCourse responseCourse
		err = rows.Scan(&newCourse.Code, &newCourse.Name, &newCourse.Credits)
		if err != nil {
			log.Printf("Error while scanning department courses: %v, err: %v\n", department.Code, err)
		}
		courses = append(courses, newCourse)
	}
	return courses, nil
}

// getCoursesFromFaculty returns all the courses offered by the given faculty member.
//
func getCoursesFromFaculty(db *sql.DB, facultyMember common.FacultyMember) (responseCourses, error) {
	// Validate department
	var deptID int
	err := db.QueryRow(common.TableReadDepartment, facultyMember.Department.Code).Scan(&deptID)
	if err != nil {
		return nil, err
	}

	query := `SELECT id FROM faculty WHERE name=$1 AND department=$2`
	// Validate faculty
	var facultyID int
	err = db.QueryRow(query, facultyMember.Name, deptID).Scan(&facultyID)
	if err != nil {
		return nil, err
	}

	query = `SELECT c.code, c.name, c.credits, d.code, d.name 
				FROM course_faculty cf, courses c, departments d 
				WHERE cf.faculty=$1 AND cf.course=c.id AND d.id=c.department`
	rows, err := db.Query(query, facultyID)
	if err != nil {
		return nil, err
	}
	var courses responseCourses
	for rows.Next() {
		newCourse := newResponseCourse()
		err = rows.Scan(&newCourse.Code, &newCourse.Name, &newCourse.Credits,
			&newCourse.Department.Code, &newCourse.Department.Name)
		if err != nil {
			log.Printf("Error while scanning faculty courses: %+v, err: %v\n", facultyMember, err)
		}
		courses = append(courses, *newCourse)
	}

	return courses, nil
}
