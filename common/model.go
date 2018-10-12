package common

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Department represents the metadata related to a department.
type Department struct {
	ID   string `json:"-"` // DO NOT EXPORT IDs, ever, for now.
	Name string `json:"name"`
	Code string `json:"code"`
}

// Departments represents the response returned by the DepartmentsHandler.
type Departments []Department

// GetInfo method populates the receiver with id and name for the given code.
func (d *Department) GetInfo(db *sql.DB) error {
	// Validate the received department code
	query := "SELECT id, name FROM departments WHERE code=$1"
	row := db.QueryRow(query, d.Code)
	return row.Scan(&d.ID, &d.Name)
}

// FacultyDesignation represents the designation of a Faculty member.
type FacultyDesignation string

// FacultyMember represents the information related to a faculty member at IIT KGP.
type FacultyMember struct {
	Name        string             `json:"name"`
	Department  Department         `json:"department"`
	Designation FacultyDesignation `json:"designation"`
	// TODO: [mcmp] Add research interests
}

// Faculty represents the response returned by the FacultyHandler.
type Faculty []FacultyMember

// Course denotes the information related to each course.
type Course struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Credits int    `json:"credits"`
	// TODO: Add syllabus
	// TODO: Add prerequisites
}

// Courses represents the reponse by the CoursesHandler.
type Courses []Course

// Slot represents a course slot used for allocating a subject.
type Slot string

type Slots []Slot

// Time represents a possible time for scheduling class.
type Time struct {

	// TODO: Use time.Weekday
	Day string `json:"day"`
	// TODO: Use time.Time
	Time string `json:"time"`
}

type TimeSlot struct {
	Time `json:"time"`
	Slot `json:"slot"`
}

type Room string

type Rooms []Room

type TimetableSlot struct {
	Course   `json:"course"`
	TimeSlot `json:"slot"`
	Rooms    `json:"rooms"`
}

type TimetableSlots []TimetableSlot

type Timetable struct {
	Monday    TimetableSlots `json:"Monday"`
	Tuesday   TimetableSlots `json:"Tuesday"`
	Wednesday TimetableSlots `json:"Wednesday"`
	Thursday  TimetableSlots `json:"Thursday"`
	Friday    TimetableSlots `json:"Friday"`
}

// RespondWithJSON is the common function to be used by all the handlers while
// returning JSON data to the caller.
func RespondWithJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(w, ErrDataEncoding.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = io.Copy(w, &buf)
	if err != nil {
		log.Println("RespondWithJSON:", err)
	}
}

// DecodeFromJSON is the common function to be used by all the POST handlers for
// reading JSON data from the request body and performing input validation.
func DecodeFromJSON(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
