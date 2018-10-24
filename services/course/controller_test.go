package course

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func newFakeServiceCourse(hasRoutes bool) *serviceCourse {
	ServiceCourse := &serviceCourse{
		common.Service{
			Name: "service-course",
			URL:  "/course",
		},
		NewFakeDataSouce(),
	}
	if hasRoutes {
		ServiceCourse.Router = initRoutes(ServiceCourse)
	}

	return ServiceCourse
}

func setup(sc *serviceCourse) (*httptest.Server, error) {
	err := sc.ConnectDB("dummy_url")
	if err != nil {
		return nil, nil
	}
	testServer := httptest.NewServer(sc)
	return testServer, nil
}

func teardown(sc *serviceCourse) error {
	return sc.CloseDB()
}

func Test_handlerCourseTimetable(t *testing.T) {
	testServiceCourse := newFakeServiceCourse(false)
	defer teardown(testServiceCourse)
	// TODO: Handle error returned during teardown

	tt := []struct {
		name           string
		method         string
		code           string
		expectedStatus int
		expectedBody   common.Timetable
	}{
		{
			name:           "Valid course",
			method:         "GET",
			code:           "MA10496",
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
			name:           "Invalid course",
			method:         "GET",
			code:           "ABCD",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "No course code",
			method:         "GET",
			code:           "",
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

			handler := http.HandlerFunc(testServiceCourse.handlerCourseTimetable)
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

func Test_handlerCoursesFromDepartment(t *testing.T) {
	testServiceCourse := newFakeServiceCourse(false)
	defer teardown(testServiceCourse)
	// TODO: Handle error returned during teardown

	tt := []struct {
		name           string
		method         string
		code           string
		expectedStatus int
		expectedBody   responseCourses
	}{
		{
			name:           "Valid Department MA",
			method:         "GET",
			code:           "MA",
			expectedStatus: http.StatusOK,
			expectedBody: responseCourses{
				responseCourse{
					Code:    "MA10496",
					Name:    "MATHEMATICS DUMMY COURSE",
					Credits: 10,
				},
			},
		},
		{
			name:           "Invalid Department",
			method:         "GET",
			code:           "ABCD",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "No Department",
			method:         "GET",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "", nil)
			assert.NoError(t, err)

			if tc.code != "" {
				req = mux.SetURLVars(req, map[string]string{"code": tc.code})
			}
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(testServiceCourse.handlerCoursesFromDepartment)
			handler.ServeHTTP(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)

			defer res.Body.Close()
			if tc.expectedStatus == http.StatusOK {
				assert.NotNil(t, res.Body)

				var responseBody responseCourses
				err := json.NewDecoder(res.Body).Decode(&responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody, responseBody)
			}
		})
	}
}

func Test_handlerCoursesFromFaculty(t *testing.T) {
	testServiceCourse := newFakeServiceCourse(false)
	defer teardown(testServiceCourse)
	// TODO: Handle error returned during teardown

	tt := []struct {
		name           string
		method         string
		params         url.Values
		expectedStatus int
		expectedBody   responseCourses
	}{
		{
			name:   "Valid Faculty",
			method: "GET",
			params: url.Values{
				"name": {"DUMMY FACULTY MEMBER"},
				"dept": {"MA"},
			},
			expectedStatus: http.StatusOK,
			expectedBody: responseCourses{
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
			},
		},
		{
			name:   "No Faculty Name",
			method: "GET",
			params: url.Values{
				"dept": {"MA"},
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "No Faculty Dept",
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
			assert.NoError(t, err)

			req.URL.RawQuery = tc.params.Encode()
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(testServiceCourse.handlerCoursesFromFaculty)
			handler.ServeHTTP(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)

			defer res.Body.Close()
			if tc.expectedStatus == http.StatusOK {
				assert.NotNil(t, res.Body)

				var responseBody responseCourses
				err := json.NewDecoder(res.Body).Decode(&responseBody)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedBody, responseBody)
			}
		})
	}
}
