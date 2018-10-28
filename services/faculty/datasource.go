package faculty

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kshitij10496/hercules/common"
)

type facultyDataSource interface {
	ConnectDS(string) error
	CloseDS() error

	GetFaculty() (common.Faculty, error)
	GetFacultyDepartment(string) (common.Faculty, error)
	GetTimetable(common.FacultyMember) (*common.Timetable, error)
}

type realFacultyDataSource struct {
	db *sql.DB
}

func newRealFacultyDataSource() *realFacultyDataSource {
	log.Println("creating a new real facultyDataSource...")
	return &realFacultyDataSource{db: nil}
}

func (ds *realFacultyDataSource) ConnectDS(url string) error {
	db, err := sql.Open("postgres", url)
	if err == nil {
		ds.db = db
	}
	return err
}

func (ds *realFacultyDataSource) CloseDS() error {
	return ds.db.Close()
}

func (ds *realFacultyDataSource) GetFaculty() (common.Faculty, error) {
	return GetFaculty(ds.db)
}

func (ds *realFacultyDataSource) GetFacultyDepartment(code string) (common.Faculty, error) {
	return GetFacultyDepartment(ds.db, code)
}

func (ds *realFacultyDataSource) GetTimetable(f common.FacultyMember) (*common.Timetable, error) {
	return GetTimetable(ds.db, f)
}

type fakeFacultyDataSource struct {
	db string
}

func newFakeFacultyDataSource() *fakeFacultyDataSource {
	log.Println("Creating a new fake facultyDataSource")
	return &fakeFacultyDataSource{"dummy"}
}

func (f *fakeFacultyDataSource) ConnectDS(url string) error {
	log.Printf("Connecting to fake facultyDataSource: %v\n", url)
	return nil
}

func (f *fakeFacultyDataSource) CloseDS() error {
	log.Println("Closing connection to fake facultyDataSource")
	return nil
}

func (f *fakeFacultyDataSource) GetFaculty() (common.Faculty, error) {
	data := common.Faculty{
		common.FacultyMember{
			Name: "Dummy Prof MA",
			Department: common.Department{
				Name: "Mathematics",
				Code: "MA",
			},
			Designation: common.FacultyDesignation("Professor"),
		},
		common.FacultyMember{
			Name: "Dummy Assistant Prof CS",
			Department: common.Department{
				Name: "Computer Science",
				Code: "CS",
			},
			Designation: common.FacultyDesignation("Assistant Professor"),
		},
	}
	return data, nil
}
func (f *fakeFacultyDataSource) GetFacultyDepartment(code string) (data common.Faculty, err error) {
	switch code {
	case "MA":
		data = common.Faculty{
			common.FacultyMember{
				Name: "Dummy Prof MA",
				Department: common.Department{
					Name: "Mathematics",
					Code: "MA",
				},
				Designation: common.FacultyDesignation("Professor"),
			},
		}
	default:
		return nil, fmt.Errorf("invalid department: %v", code)
	}
	return data, nil
}
func (f *fakeFacultyDataSource) GetTimetable(fm common.FacultyMember) (data *common.Timetable, err error) {
	// TODO: Enter mock data
	mockFacultyMember := common.FacultyMember{
		Name:       "DUMMY FACULTY MEMBER",
		Department: common.Department{Code: "MA"},
	}
	switch fm {
	case mockFacultyMember:
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
		return nil, fmt.Errorf("invalid faculty member: %+v", fm)
	}
	return data, nil
}
