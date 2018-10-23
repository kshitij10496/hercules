package course

import (
	"database/sql"
	"log"

	"github.com/kshitij10496/hercules/common"
)

type courseDataSource interface {
	ConnectDS(string) error
	CloseDS() error

	GetCoursesFromDepartment(common.Department) (responseCourses, error)
	GetDepartmentInfo(*common.Department) error
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
	return nil
}

func (f *fakeDataSource) GetCoursesFromDepartment(d common.Department) (responseCourses, error) {
	// TODO: Enter mock data for testing
	return nil, nil
}

func (f *fakeDataSource) GetCoursesFromFaculty(fm common.FacultyMember) (responseCourses, error) {
	// TODO: Enter mock data for testing
	return nil, nil
}

func (f *fakeDataSource) GetCourseTimetable(c common.Course) (*common.Timetable, error) {
	// TODO: Enter mock data for testing
	return nil, nil
}
