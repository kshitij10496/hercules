package course

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
			"Valid Department MA",
			"GET",
			"MA",
			http.StatusOK,
			responseCourses{
				responseCourse{
					Code:    "MA10496",
					Name:    "MA Course 1",
					Credits: 10,
				},
			},
		},
		{
			"Invalid Department",
			"GET",
			"ABCD",
			http.StatusBadRequest,
			responseCourses{
				responseCourse{
					Code:    "CS10496",
					Name:    "CS Course 1",
					Credits: 10,
				},
			},
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
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
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
