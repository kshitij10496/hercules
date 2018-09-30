package main

// Department represents the metadata related to a department.
type Department struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// DepartmentsResponse represents the response returned by the DepartmentsHandler.
type DepartmentsResponse []Department

// Course denotes the information related to each course.
type Course struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Credits int    `json:"credits"`
	// TODO: Add syllabus
	// TODO: Add prerequisites
}

// CoursesResponse represents the reponse by the CoursesHandler.
type CoursesResponse []Course

// FacultyDesignation represents the designation of a Faculty member.
type FacultyDesignation string

// Professor represents the faculty designation of a Professor.
const Professor = FacultyDesignation("Professor")

// AssociateProfessor represents the faculty designation of an Associate Professor.
const AssociateProfessor = FacultyDesignation("Associate Professor")

// VisitingFaculty represents the faculty designation of a Visiting Faculty.
const VisitingFaculty = FacultyDesignation("Visiting Faculty")

// Faculty represents the information related to a faculty member at IIT KGP.
type Faculty struct {
	Name        string             `json:"name"`
	Department  Department         `json:"department"`
	Designation FacultyDesignation `json:"designation"`
	// TODO: [mcmp] Add research interests
}

// FacultyResponse represents the response returned by the FacultyHandler.
type FacultyResponse []Faculty
