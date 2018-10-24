package course

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/kshitij10496/hercules/common"
)

type courseDataSource interface {
	ConnectDS(string) error
	CloseDS() error

	GetDepartmentInfo(*common.Department) error
	GetCourseInfo(*common.Course) error

	GetCoursesFromDepartment(common.Department) (responseCourses, error)
	GetCoursesFromFaculty(common.FacultyMember) (responseCourses, error)
	GetCourseTimetable(common.Course) (*common.Timetable, error)
}

type realDataSource struct {
	db *sql.DB
}

func NewRealDataSource() *realDataSource {
	log.Println("creating a new real datasource...")
	return &realDataSource{db: nil}
}

func (ds *realDataSource) ConnectDS(url string) error {
	db, err := sql.Open("postgres", url)
	if err == nil {
		ds.db = db
	}
	return err
}

func (ds *realDataSource) CloseDS() error {
	return ds.db.Close()
}

func (ds *realDataSource) GetDepartmentInfo(d *common.Department) error {
	return d.GetInfo(ds.db)
}

func (ds *realDataSource) GetCourseInfo(c *common.Course) error {
	return c.GetInfo(ds.db)
}

func (ds *realDataSource) GetCoursesFromDepartment(d common.Department) (responseCourses, error) {
	return getCoursesFromDepartment(ds.db, d)
}

func (ds *realDataSource) GetCoursesFromFaculty(f common.FacultyMember) (responseCourses, error) {
	return getCoursesFromFaculty(ds.db, f)
}

func (ds *realDataSource) GetCourseTimetable(c common.Course) (*common.Timetable, error) {
	return getCourseTimetable(ds.db, c)
}

// fakeDataSource implements the courseDatasource interface.
// This helps mock the DB; primarily used for testing.
type fakeDataSource struct {
	db string
}

func NewFakeDataSouce() *fakeDataSource {
	log.Println("Creating a new fake courseDataSource")
	return &fakeDataSource{"dummy"}
}
func (f *fakeDataSource) ConnectDS(url string) error {
	log.Printf("Connecting to fake courseDataSource: %v\n", url)
	return nil
}

func (f *fakeDataSource) CloseDS() error {
	log.Println("Closing connection to fake courseDataSource")
	return nil
}

func (f *fakeDataSource) GetDepartmentInfo(d *common.Department) error {
	// TODO: Enter mock data for testing
	switch d.Code {
	case "MA":
		d.ID = "1"
	case "CS":
		d.ID = "2"
	default:
		return errors.New("not a valid department")
	}
	return nil
}

func (f *fakeDataSource) GetCourseInfo(c *common.Course) error {
	switch c.Code {
	case "MA10496":
		c.Name = "MATHEMATICS DUMMY COURSE"
		c.Credits = 10
	default:
		return errors.New("not a valid course")
	}
	return nil
}

func (f *fakeDataSource) GetCoursesFromDepartment(d common.Department) (data responseCourses, err error) {
	switch d.ID {
	case "1":
		data = responseCourses{
			responseCourse{
				Code:    "MA10496",
				Name:    "MATHEMATICS DUMMY COURSE",
				Credits: 10,
			},
		}
	case "2":
		data = responseCourses{
			responseCourse{
				Code:    "CS10496",
				Name:    "CS DUMMY COURSE",
				Credits: 10,
			},
		}
	default:
		return nil, fmt.Errorf("Cannot find courses for dept:%+v", d.Code)
	}
	return data, nil
}

func (f *fakeDataSource) GetCoursesFromFaculty(fm common.FacultyMember) (data responseCourses, err error) {
	switch fm.Name {
	case "DUMMY FACULTY MEMBER":
		data = responseCourses{
			responseCourse{
				Code:    "MA10496",
				Name:    "MATHEMATICS DUMMY COURSE",
				Credits: 10,
				Department: &common.Department{
					Name: "Mathematics",
					Code: "MA",
				},
			},
			responseCourse{
				Code:    "CS10496",
				Name:    "CS DUMMY COURSE",
				Credits: 10,
				Department: &common.Department{
					Name: "Computer Science",
					Code: "CS",
				},
			},
		}
	default:
		return nil, fmt.Errorf("No a valid faculty member")
	}
	return data, nil
}

func (f *fakeDataSource) GetCourseTimetable(c common.Course) (data *common.Timetable, err error) {
	switch c.Code {
	case "MA10496":
		data = &common.Timetable{
			Monday: common.TimetableSlots{
				common.TimetableSlot{
					common.Course{
						Name:    "MATHEMATICS DUMMY COURSE",
						Code:    "MA10496",
						Credits: 10,
					},
					common.TimeSlot{
						common.Time{
							Day:  "Monday",
							Time: "12 PM",
						},
						common.Slot("DUMMY SLOT"),
					},
					common.Rooms{
						common.Room("DUMMY ROOM"),
					},
				},
			},
			Tuesday: common.TimetableSlots{
				common.TimetableSlot{
					common.Course{
						Name:    "MATHEMATICS DUMMY COURSE",
						Code:    "MA10496",
						Credits: 10,
					},
					common.TimeSlot{
						common.Time{
							Day:  "Tuesday",
							Time: "10 AM",
						},
						common.Slot("DUMMY SLOT"),
					},
					common.Rooms{
						common.Room("DUMMY ROOM"),
					},
				},
				common.TimetableSlot{
					common.Course{
						Name:    "MATHEMATICS DUMMY COURSE",
						Code:    "MA10496",
						Credits: 10,
					},
					common.TimeSlot{
						common.Time{
							Day:  "Tuesday",
							Time: "11 AM",
						},
						common.Slot("DUMMY SLOT"),
					},
					common.Rooms{
						common.Room("DUMMY ROOM"),
					},
				},
			},
		}
	default:
		return nil, fmt.Errorf("No timetable for course: %+v", c.Code)
	}
	return data, nil
}
