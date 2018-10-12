package migration

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/kshitij10496/hercules/common"
)

var facultyDirectoryFile = "/Users/kshitij10496/Software/go/src/github.com/kshitij10496/hercules/data/faculty_directory.json"
var coursesDetailsFile = "/Users/kshitij10496/Software/go/src/github.com/kshitij10496/hercules/data/courses_details.json"
var slotsFile = "/Users/kshitij10496/Software/go/src/github.com/kshitij10496/hercules/data/slots.csv"

// serviceMigration implements the server interface
//
type serviceMigration common.Service

func (s *serviceMigration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: What should go in here?
	log.Printf("[request initiate] %s - %v\n", s.Name, r.URL)
	s.Router.ServeHTTP(w, r)
	log.Printf("[request end] %s - %v\n", s.Name, r.URL)
}

func (s *serviceMigration) GetDBConnection(ctx context.Context) (*sql.Conn, error) {
	return s.DB.Conn(ctx)
}

func (s *serviceMigration) GetName() string {
	return s.Name
}

func (s *serviceMigration) GetURL() string {
	return s.URL
}

func migrations(db *sql.DB) (err error) {
	// TODO: Find an effective way to perform data migration.

	// Migrate `departments`
	// err = readFromJSONDepartments(db, facultyDirectoryFile)
	// if err != nil {
	// 	return err
	// }

	// Migrate `slots` and `time_slots`
	// err = readFromTimeSlots(db, slotsFile)
	// if err != nil {
	// 	return err
	// }

	// Migrate `faculty_designations` and `faculty`
	// err = readFromJSONFaculty(db, facultyDirectoryFile)
	// if err != nil {
	// 	return err
	// }

	// Migrate `rooms`, `courses`, `courses_faculty`, `courses_slots` and `courses_rooms`
	// err = readFromCourses(db, coursesDetailsFile)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// SetDB sets the service to use the given DB.
// Note that this function overwrites the current value.
//
func (s *serviceMigration) ConnectDB(url string) error {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	s.DB = db
	return migrations(db)
}

func (s *serviceMigration) CloseDB() error {
	return s.DB.Close()
}

// serviceMigration represents the course service.
var ServiceMigration serviceMigration

func init() {
	ServiceMigration = serviceMigration{
		Name:   "service-migration",
		URL:    "/migration",
		Router: common.NewSubRouter(nil),
	}
}
