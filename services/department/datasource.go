package department

import (
	"database/sql"
	"log"

	"github.com/kshitij10496/hercules/common"
)

type departmentsDatasource interface {
	ConnectDS(string) error
	CloseDS() error

	GetDepartments() (common.Departments, error)
}

type realDataSource struct {
	db *sql.DB
}

func newRealDataSource() *realDataSource {
	log.Println("creating a new real departmentsDatasource...")
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

func (ds *realDataSource) GetDepartments() (common.Departments, error) {
	return GetDepartments(ds.db)
}

// fakeDataSource implements the departmentsDatasource interface.
// This helps mock the DB; primarily used for testing.
type fakeDataSource struct {
	db string
}

func newFakeDataSouce() *fakeDataSource {
	log.Println("Creating a new fake departmentsDatasource")
	return &fakeDataSource{"dummy"}
}
func (f *fakeDataSource) ConnectDS(url string) error {
	log.Printf("Connecting to fake departmentsDatasource: %v\n", url)
	return nil
}

func (f *fakeDataSource) CloseDS() error {
	log.Println("Closing connection to fake departmentsDatasource")
	return nil
}

func (f *fakeDataSource) GetDepartments() (common.Departments, error) {
	depts := []common.Department{
		{
			Name: "Mathematics",
			Code: "MA",
		},
		{
			Name: "Computer Science",
			Code: "CS",
		},
	}
	return common.Departments(depts), nil
}
