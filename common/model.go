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
	Name string `json:"name"`
	Code string `json:"code"`
}

// Departments represents the response returned by the DepartmentsHandler.
type Departments []Department

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

// GetCourse returns all the information related to a course given the course code.
func (c *Course) GetCourseInfo(db *sql.DB) error {
	// TODO: Replace this with a DB lookup
	if c.Code == "NA61001" {
		c.Name = "COASTAL ENGINEERING"
		c.Credits = 3
		return nil
	}

	return ErrCourseNotFound
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

// SlotTime represents the daily timing slots.
type SlotTime int

const (
	AM8 = iota + 1
	AM9
	AM10
	AM11
	PM12
	PM2
	PM3
	PM4
	PM5
)

type Room string

const (
	// Nalanda Rooms
	NR121 = Room("NR121")
	NR122 = Room("NR122")
	NR123 = Room("NR123")
	NR124 = Room("NR124")

	NR221 = Room("NR221")
	NR222 = Room("NR222")
	NR223 = Room("NR223")
	NR224 = Room("NR224")

	NR321 = Room("NR321")
	NR322 = Room("NR322")
	NR323 = Room("NR323")
	NR324 = Room("NR324")

	NR421 = Room("NR421")
	NR422 = Room("NR422")
	NR423 = Room("NR423")
	NR424 = Room("NR424")

	NC141 = Room("NC141")
	NC142 = Room("NC142")
	NC143 = Room("NC143")
	NC144 = Room("NC144")

	NC241 = Room("NC241")
	NC242 = Room("NC242")
	NC243 = Room("NC243")
	NC244 = Room("NC244")

	NC341 = Room("NC341")
	NC342 = Room("NC342")
	NC343 = Room("NC343")
	NC344 = Room("NC344")

	NC441 = Room("NC441")
	NC442 = Room("NC442")
	NC443 = Room("NC443")
	NC444 = Room("NC444")

	NC131 = Room("NC131")
	NC132 = Room("NC132")
	NC133 = Room("NC133")
	NC134 = Room("NC134")

	NC231 = Room("NC231")
	NC232 = Room("NC232")
	NC233 = Room("NC233")
	NC234 = Room("NC234")

	NC331 = Room("NC331")
	NC332 = Room("NC332")
	NC333 = Room("NC333")
	NC334 = Room("NC334")

	NC431 = Room("NC431")
	NC432 = Room("NC432")
	NC433 = Room("NC433")
	NC434 = Room("NC434")
)

type Slot struct {
	Course Course   `json:"course"`
	Timing SlotTime `json:"slot"`
	Room   Room     `json:"room"`
}

type Timetable struct {
	Monday    []Slot `json:"Monday"`
	Tuesday   []Slot `json:"Tuesday"`
	Wednesday []Slot `json:"Wednesday"`
	Thursday  []Slot `json:"Thursday"`
	Friday    []Slot `json:"Friday"`
}

type TimeTableResponse struct {
	Timetable Timetable `json:"timetable"`
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
