package migration

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/kshitij10496/hercules/common"
)

func containsDesignations(designations []common.FacultyDesignation, d common.FacultyDesignation) bool {
	for _, designation := range designations {
		if d == designation {
			return true
		}
	}
	return false
}

func containsDepartments(departments []readDepartment, d readDepartment) bool {
	for _, department := range departments {
		if d == department {
			return true
		}
	}
	return false
}

func createSetDepartments(departments []readDepartment) []readDepartment {
	setDepartments := []readDepartment{}
	for _, d := range departments {
		if !containsDepartments(setDepartments, d) {
			setDepartments = append(setDepartments, d)
		}
	}
	return setDepartments
}

func createSetDesignations(designations []common.FacultyDesignation) []common.FacultyDesignation {
	setDesignations := []common.FacultyDesignation{}
	for _, d := range designations {
		if !containsDesignations(setDesignations, d) {
			setDesignations = append(setDesignations, d)
		}
	}
	return setDesignations
}

type readDepartment struct {
	Name string `json:"department"`
	Code string `json:"code"`
}

func readFromJSONDepartments(db *sql.DB, filename string) error {
	// Open JSON file
	departmentsFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer departmentsFile.Close()

	// Load data into a "Departments" value
	decoder := json.NewDecoder(departmentsFile)
	var departments []readDepartment
	err = decoder.Decode(&departments)
	if err != nil {
		return err
	}

	// Insert data into db
	// TODO: Implement a SQL transaction here.
	departmentsSet := createSetDepartments(departments)
	log.Println("Department Set:", departmentsSet)
	for _, department := range departmentsSet {
		_, err := db.Exec(common.TableInsertionDepartments, department.Code, department.Name)
		if err != nil {
			log.Println("[insertion] departments:", department, err)
		}
	}

	return nil
}

func readFromJSONFacultyDesignations(db *sql.DB, filename string) error {
	designationsFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer designationsFile.Close()

	decoder := json.NewDecoder(designationsFile)
	var designations []common.FacultyDesignation
	err = decoder.Decode(&designations)
	if err != nil {
		return err
	}

	designationsSet := createSetDesignations(designations)
	log.Println("DESIGNATION SET:", designationsSet)
	for _, designation := range designationsSet {
		_, err := db.Exec(common.TableInsertionFacultyDesignations, designation)
		if err != nil {
			log.Println("[insertion] faculty_designations:", designation, err)
		}
	}
	return nil
}

type readFacultyMember struct {
	Name        string `json:"faculty"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
	DeptCode    string `json:"code"`
}

type readFaculty []readFacultyMember

func readFromJSONFaculty(db *sql.DB, filename string) error {
	facultyFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer facultyFile.Close()

	decoder := json.NewDecoder(facultyFile)
	var faculty readFaculty
	err = decoder.Decode(&faculty)
	if err != nil {
		return err
	}

	for _, member := range faculty {
		var designationID, departmentID int

		err := db.QueryRow(common.TableReadDesignation, member.Designation).Scan(&designationID)
		if err != nil {
			err = db.QueryRow(common.TableInsertionFacultyDesignations, member.Designation).Scan(&designationID)
			if err != nil {
				log.Println("[insertion] faculty_designations:", member.Designation, err)
				continue
			}
		}

		row := db.QueryRow(common.TableReadDepartment, member.DeptCode)
		if err := row.Scan(&departmentID); err != nil {
			log.Println("[read] departments:", member.DeptCode, err)
			continue
		}

		_, err = db.Exec(common.TableInsertionFaculty, member.Name, designationID, departmentID)
		if err != nil {
			log.Printf("[insertion] faculty: %v, %v, %v, err=%v\n", member, designationID, departmentID, err)
			continue
		}
	}
	return nil
}

type readCourse struct {
	Code    string   `json:"code"`
	Name    string   `json:"name"`
	Profs   []string `json:"profs"`
	Credits int      `json:"credits"`
	Slots   []string `json:"slots"`
	Rooms   []string `json:"rooms"`
}

type readCourses []readCourse

type readDepartmentCourse struct {
	Department string      `json:"dept"`
	Courses    readCourses `json:"courses"`
}

type readDepartmentCourses []readDepartmentCourse

func readFromCourses(db *sql.DB, filename string) error {
	coursesFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer coursesFile.Close()

	decoder := json.NewDecoder(coursesFile)

	var departmentCourses readDepartmentCourses
	err = decoder.Decode(&departmentCourses)
	if err != nil {
		return err
	}

	for _, deptCourses := range departmentCourses {
		var deptID int

		row := db.QueryRow(common.TableReadDepartment, deptCourses.Department)
		if err := row.Scan(&deptID); err != nil {
			log.Println("[read] department:", deptCourses.Department, err)
			continue
		}

		for _, course := range deptCourses.Courses {
			// Add the course to the `courses` table
			var courseID int
			err := db.QueryRow(common.TableInsertionCourses, course.Code,
				course.Name, course.Credits, deptID).Scan(&courseID)
			if err != nil {
				log.Println("[insertion] courses:", course, err)
			}

			// Insert course-faculty mapping in `course_faculty` table
			for _, prof := range course.Profs {
				// Find the professor's unique ID and add it to DB
				var profID int

				row := db.QueryRow(common.TableReadFaculty, prof)
				if err := row.Scan(&profID); err != nil {
					log.Println("[read] faculty:", prof, err)
					continue
				}

				_, err = db.Exec(common.TableInsertionCourseFaculty, profID, courseID)
				if err != nil {
					log.Printf("[insertion] course_faculty: %v, %v, err = %v\n", profID, courseID, err)
				}
			}

			// Add all the course-slots mapping to `course_slots` table
			for _, slot := range course.Slots {
				// Fetch the slotID from `slots` table
				var slotID int

				err := db.QueryRow(common.TableReadSlots, slot).Scan(&slotID)
				if err != nil {
					log.Println("[read] slots:", slot, err)
					err = db.QueryRow(common.TableInsertionSlots, slot).Scan(&slotID)
					if err != nil {
						log.Printf("[insertion] slots: %v, err=%v\n", slot, err)
						continue
					}
				}

				// Find the slot id from the "slots" table
				_, err = db.Exec(common.TableInsertionCourseSlots, slotID, courseID)
				if err != nil {
					log.Printf("[insertion] course_slots: %v, %v, err = %v\n", slotID, courseID, err)
				}
			}

			// Add all the rooms to `rooms` table and `course_rooms` table
			for _, room := range course.Rooms {
				// Insert room into `rooms` if it doesn't already exists.
				var roomID int
				err := db.QueryRow(common.TableReadRooms, room).Scan(&roomID)
				if err != nil {
					err := db.QueryRow(common.TableInsertionRooms, room).Scan(&roomID)
					if err != nil {
						log.Println("[insertion] rooms:", room, err)
						continue
					}
				}

				// Add room to `course_rooms` mapping table
				_, err = db.Exec(common.TableInsertionCourseRooms, roomID, courseID)
				if err != nil {
					log.Printf("[insertion] course_rooms: %v, %v, err = %v\n", roomID, courseID, err)
				}
			}
		}
	}
	return nil
}

func readFromTimeSlots(db *sql.DB, filename string) error {
	timeSlotsFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer timeSlotsFile.Close()

	csvReader := csv.NewReader(timeSlotsFile)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading csv file: %v\n", err)
			return err
		}

		slot, times := record[0], record[1:]
		var slotID int
		err = db.QueryRow(common.TableInsertionSlots, slot).Scan(&slotID)
		if err != nil {
			log.Printf("[insertion] slots: %v, err=%v\n", slot, err)
			continue
		}

		for _, time := range times {
			// Compute time
			t, err := strconv.Atoi(time)
			if err != nil {
				log.Println("Cannot convert time to int:", err)
				continue
			}

			// Convert slot times to DB time ids
			switch {
			case 00 <= t && t < 10:
				t = t + 1
			case 10 <= t && t < 20:
				t = t
			case 20 <= t && t < 30:
				t = t - 1
			case 30 <= t && t < 40:
				t = t - 2
			case 40 <= t && t < 50:
				t = t - 3
			default:
				log.Println("Invalid value of time", time)
				continue
			}
			// Possible formula t = t - ((t - 10) / 10)

			_, err = db.Exec(common.TableInsertionTimeSlots, t, slotID)
			if err != nil {
				log.Printf("[insertion] time slots: %v, %v, err=%v\n", t, slotID, err)
			}
		}
	}
	return nil
}
