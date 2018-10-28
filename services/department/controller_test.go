package department

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kshitij10496/hercules/common"
	"github.com/stretchr/testify/assert"
)

func newFakeServiceDepartment(hasRoutes bool) *serviceDepartment {
	testDepartmentService := &serviceDepartment{
		Service: common.Service{
			Name: "service-department",
			URL:  "/department",
		},
		DB: newFakeDataSouce(),
	}
	if hasRoutes {
		testDepartmentService.Router = initRoutes(testDepartmentService)
	}
	return testDepartmentService
}

func setup(sd *serviceDepartment) (*httptest.Server, error) {
	err := sd.ConnectDB("dummy_url")
	if err != nil {
		return nil, nil
	}
	testServer := httptest.NewServer(sd)
	return testServer, nil
}

func teardown(sd *serviceDepartment) error {
	return sd.CloseDB()
}

func Test_Handler_GetDepartments(t *testing.T) {
	testDepartmentService := newFakeServiceDepartment(false)
	defer teardown(testDepartmentService)

	tt := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   common.Departments
	}{
		{
			name:           "Valid request",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody: common.Departments{
				common.Department{
					Name: "Mathematics",
					Code: "MA",
				},
				common.Department{
					Name: "Computer Science",
					Code: "CS",
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

			handler := http.HandlerFunc(testDepartmentService.handlerDepartments)
			handler.ServeHTTP(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)

			defer res.Body.Close()
			assert.NotNil(t, res.Body)

			var responseBody common.Departments
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, responseBody)
		})
	}
}
