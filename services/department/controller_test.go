package department

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kshitij10496/hercules/common"
	"github.com/stretchr/testify/assert"
)

var testDepartmentService *serviceDepartment
var testServer *httptest.Server

func newFakeServiceDepartment(hasRoutes bool) *serviceDepartment {
	testDepartmentService = &serviceDepartment{
		Service: common.Service{
			Name: "service-department",
			URL:  "/department",
		},
		DB: NewFakeDataSouce(),
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
	testServer = httptest.NewServer(sd)
	return testServer, nil
}

func teardown(sd *serviceDepartment) error {
	return testDepartmentService.CloseDB()
}

func Test_Handler_GetDepartments(t *testing.T) {
	// Setup tests
	testDepartmentService := newFakeServiceDepartment(false)
	// testServer, err := setup(testDepartmentService)
	// assert.NoError(t, err)
	// Teardown tests
	// defer testServer.Close()
	defer teardown(testDepartmentService)
	// TODO: Handler error during teardown

	// endpoint := "/info/all"

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

	// url := testServer.URL + common.VERSION + testDepartmentService.URL + endpoint
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "", nil)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			rec := httptest.NewRecorder()

			testDepartmentService.handlerDepartments(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			assert.NotNil(t, res.Body)

			var responseBody common.Departments
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, responseBody)
		})
	}
}
