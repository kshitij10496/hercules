package faculty

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/kshitij10496/hercules/common"
)

// GetFaculty returns the list of faculty members in IITKGP.
func GetFaculty(db *sql.DB) (data common.Faculty, err error) {
	query := `SELECT f.name, fd.designation, d.code, d.name 
			FROM departments d, faculty_designations fd, faculty f 
			WHERE f.designation=fd.id AND f.department=d.id;`
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

type responseCourse struct {
	Code       string             `json:"code"`
	Name       string             `json:"name"`
	Credits    int                `json:"credits"`
	Department *common.Department `json:"department,omitempty"`
}

type responseCourses []responseCourse

func GetTimetable(db *sql.DB, facultyMember common.FacultyMember) (*common.Timetable, error) {
	// Get courses for the faculty member from service-course
	query := url.Values{
		"name": []string{facultyMember.Name},
		"dept": []string{facultyMember.Department.Code},
	}
	response, err := common.SendToService("course", "GET", "/info/faculty", query, nil)
	if response.StatusCode != 200 {
		log.Println("Unable to fetch all courses from faculty:", response.StatusCode)
		return nil, fmt.Errorf("Unable to fetch all courses from faculty: %v", response.StatusCode)
	}
	defer response.Body.Close()

	var courses responseCourses
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&courses)
	if err != nil {
		log.Println("Error decoding courses:", err)
		return nil, err
	}

	var facultyTimetable common.Timetable
	for _, course := range courses {
		// Get time slots for the course
		// Get the rooms for the course
		endpoint := fmt.Sprintf("/timetable/%s", course.Code)
		response, err := common.SendToService("course", "GET", endpoint, nil, nil)
		if response.StatusCode != 200 {
			log.Println("Unable to fetch course's timetable:", course.Code)
			continue
		}
		defer response.Body.Close()

		var courseTimetable common.Timetable
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&courseTimetable)
		if err != nil {
			log.Println("Error decoding timetable for course:", course.Code, err)
			continue
		}

		facultyTimetable.Monday = append(facultyTimetable.Monday, courseTimetable.Monday...)
		facultyTimetable.Tuesday = append(facultyTimetable.Tuesday, courseTimetable.Tuesday...)
		facultyTimetable.Wednesday = append(facultyTimetable.Wednesday, courseTimetable.Wednesday...)
		facultyTimetable.Thursday = append(facultyTimetable.Thursday, courseTimetable.Thursday...)
		facultyTimetable.Friday = append(facultyTimetable.Friday, courseTimetable.Friday...)
	}

	// Build the timetable
	return &facultyTimetable, nil
}
