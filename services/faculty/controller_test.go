package faculty

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
	"github.com/stretchr/testify/assert"
)

func newFakeServiceFaculty(hasRoutes bool) *serviceFaculty {
	ServiceCourse := &serviceFaculty{
		common.Service{
			Name: "service-faculty",
			URL:  "/faculty",
		},
		newFakeFacultyDataSource(),
	}
	if hasRoutes {
		ServiceCourse.Router = initRoutes(ServiceCourse)
	}

	return ServiceCourse
}

func setup(sc *serviceFaculty) (*httptest.Server, error) {
	err := sc.ConnectDB("dummy_url")
	if err != nil {
		return nil, nil
	}
	testServer := httptest.NewServer(sc)
	return testServer, nil
}

func teardown(sc *serviceFaculty) error {
	return sc.CloseDB()
}

func Test_handlerFacultyAll(t *testing.T) {
	testFacultyService := newFakeServiceFaculty(false)
	defer teardown(testFacultyService)

	tt := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   common.Faculty
	}{
		{
			name:           "Valid Request",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody: common.Faculty{
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
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "", nil)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(testFacultyService.handlerFacultyAll)
			handler.ServeHTTP(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)

			defer res.Body.Close()
			if tc.expectedStatus == http.StatusOK {
				assert.NotNil(t, res.Body)

				var responseBody common.Faculty
				err := json.NewDecoder(res.Body).Decode(&responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody, responseBody)
			}
		})
	}
}

func Test_handlerFacultyDepartment(t *testing.T) {
	testFacultyService := newFakeServiceFaculty(false)
	defer teardown(testFacultyService)

	tt := []struct {
		name           string
		method         string
		code           string
		expectedStatus int
		expectedBody   common.Faculty
	}{
		{
			name:           "Valid Request MA",
			method:         "GET",
			code:           "MA",
			expectedStatus: http.StatusOK,
			expectedBody: common.Faculty{
				common.FacultyMember{
					Name: "Dummy Prof MA",
					Department: common.Department{
						Name: "Mathematics",
						Code: "MA",
					},
					Designation: common.FacultyDesignation("Professor"),
				},
			},
		},
		{
			name:           "Invalid Department",
			method:         "GET",
			code:           "ABCD",
			expectedStatus: http.StatusInternalServerError, // TODO: Should be BadRequest
		},
		{
			name:           "No Department",
			method:         "GET",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "", nil)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			if tc.code != "" {
				req = mux.SetURLVars(req, map[string]string{"code": tc.code})
			}

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(testFacultyService.handlerFacultyDepartment)
			handler.ServeHTTP(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)

			defer res.Body.Close()
			if tc.expectedStatus == http.StatusOK {
				assert.NotNil(t, res.Body)

				var responseBody common.Faculty
				err := json.NewDecoder(res.Body).Decode(&responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody, responseBody)
			}
		})
	}
}

func Test_handlerFacultyTimetable(t *testing.T) {
	testFacultyService := newFakeServiceFaculty(false)
	defer teardown(testFacultyService)

	tt := []struct {
		name           string
		method         string
		params         url.Values
		expectedStatus int
		expectedBody   common.Timetable
	}{
		{
			name:   "Valid Request",
			method: "GET",
			params: url.Values{
				"name": {"DUMMY FACULTY MEMBER"},
				"dept": {"MA"},
			},
			expectedStatus: http.StatusOK,
			expectedBody: common.Timetable{
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
			},
		},
		{
			name:   "Valid Request - invalid params",
			method: "GET",
			params: url.Values{
				"name": {"Non-existent FACULTY MEMBER"},
				"dept": {"dummy dept"},
			},
			expectedStatus: http.StatusInternalServerError, // TODO: Should be BadRequest
		},
		{
			name:   "Invalid Request - No name",
			method: "GET",
			params: url.Values{
				"dept": {"MA"},
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "Invalid Request - No department code",
			method: "GET",
			params: url.Values{
				"name": {"DUMMY FACULTY MEMBER"},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "", nil)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			req.URL.RawQuery = tc.params.Encode()

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(testFacultyService.handlerFacultyTimetable)
			handler.ServeHTTP(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)

			defer res.Body.Close()
			if tc.expectedStatus == http.StatusOK {
				assert.NotNil(t, res.Body)

				var responseBody common.Timetable
				err := json.NewDecoder(res.Body).Decode(&responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody, responseBody)
			}
		})
	}
}
