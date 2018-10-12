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
	Faculty    *common.Faculty    `json:"faculty,omitempty"`
	Slots      *common.Slots      `json:"slots,omitempty"`
	Rooms      *common.Rooms      `json:"rooms,omitempty"`
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
		newCourse := responseCourse{Department: &common.Department{}}
		err = rows.Scan(&newCourse.Code, &newCourse.Name, &newCourse.Credits,
			&newCourse.Department.Code, &newCourse.Department.Name)
		if err != nil {
			log.Printf("Error while scanning faculty courses: %+v, err: %v\n", facultyMember, err)
			continue
		}
		courses = append(courses, newCourse)
	}

	return courses, nil
}

func getCourseTimetable(db *sql.DB, course common.Course) (*common.Timetable, error) {
	var courseTimetable common.Timetable

	var courseID int
	query := `SELECT id, name, credits FROM courses WHERE code=$1`

	err := db.QueryRow(query, course.Code).Scan(&courseID, &course.Name, &course.Credits)
	if err != nil {
		log.Println("[read] courses:", err)
		return nil, err
	}

	query = `SELECT room FROM course_rooms WHERE course=$1;`
	rows, err := db.Query(query, courseID)
	if err != nil {
		log.Println("[read] course_rooms:", err)
	}

	var rooms common.Rooms
	for rows.Next() {
		var roomID int
		err = rows.Scan(&roomID)
		if err != nil {
			log.Println("Error scanning course_rooms:", err)
			continue
		}

		var newRoom common.Room
		query = `SELECT room FROM rooms WHERE id=$1;`
		err = db.QueryRow(query, roomID).Scan(&newRoom)
		if err != nil {
			log.Println("Error scanning rooms:", roomID, err)
			continue
		}

		rooms = append(rooms, newRoom)
	}

	query = `SELECT slot FROM course_slots WHERE course=$1;`
	rows, err = db.Query(query, courseID)
	if err != nil {
		log.Println("[read] course_slots:", err)
	}

	for rows.Next() {
		var slotID int

		err = rows.Scan(&slotID)
		if err != nil {
			log.Println("Error scanning course_slots:", err)
			continue
		}

		query = `SELECT s.slot, t.day, t.time
				FROM slots s, time t, time_slots ts 
				WHERE s.id=$1 AND s.id=ts.slot AND t.id=ts.time;`
		rows, err := db.Query(query, slotID)
		if err != nil {
			log.Println("Error querying time_slots:", slotID, err)
			continue
		}

		for rows.Next() {
			newTimeSlot := common.TimetableSlot{
				Course: course,
				Rooms:  rooms,
			}
			err = rows.Scan(&newTimeSlot.Slot, &newTimeSlot.Time.Day, &newTimeSlot.Time.Time)
			if err != nil {
				log.Println("error scanning new timetable slot:", slotID, err)
				continue
			}

			switch newTimeSlot.Day {
			case "Monday":
				courseTimetable.Monday = append(courseTimetable.Monday, newTimeSlot)
			case "Tuesday":
				courseTimetable.Tuesday = append(courseTimetable.Tuesday, newTimeSlot)
			case "Wednesday":
				courseTimetable.Wednesday = append(courseTimetable.Wednesday, newTimeSlot)
			case "Thursday":
				courseTimetable.Thursday = append(courseTimetable.Thursday, newTimeSlot)
			case "Friday":
				courseTimetable.Friday = append(courseTimetable.Friday, newTimeSlot)
			default:
				log.Println("Invalid day:", newTimeSlot.Day)
			}
		}
	}
	return &courseTimetable, nil
}
