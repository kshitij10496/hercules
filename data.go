package main

// Department represents the metadata related to a departments.
type Department struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// DepartmentsResponse represents the response returned by the DepartmentsHandler.
type DepartmentsResponse []Department
