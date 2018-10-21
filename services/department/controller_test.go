package department

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kshitij10496/hercules/common"
	"github.com/stretchr/testify/assert"
)

var testDepartmentService *serviceDepartment
var testServer *httptest.Server

func newFakeRouter() *mux.Router {
	var fakeRoutes = common.Routes{
		common.Route{
			Name:        "Information for all the departments",
			Method:      "GET",
			Pattern:     "/info/all",
			HandlerFunc: testDepartmentService.handlerDepartments,
			PathPrefix:  common.VERSION + "/department",
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range fakeRoutes {
		path := route.PathPrefix + route.Pattern
		router.
			Methods(route.Method).
			Path(path).
			Name(route.Name).
			Handler(route.HandlerFunc)

		log.Println("created route:", route.Method, path)
	}
	return router
}

func setup() error {
	testDepartmentService = &serviceDepartment{
		Service: common.Service{
			Name: "service-department",
			URL:  "/department",
		},
		DB: NewFakeDataSouce(),
	}

	testDepartmentService.Router = newFakeRouter()
	testServer = httptest.NewServer(testDepartmentService)
	return testDepartmentService.ConnectDB("dummy_url")
}

func teardown() error {
	testServer.Close()
	return testDepartmentService.CloseDB()
}

func Test_GetDepartments(t *testing.T) {
	err := setup()
	assert.NoError(t, err)
	// TODO: Handler error during teardown
	defer teardown()

	fmt.Printf("testDepartmentService: %+v\n", testDepartmentService)
	endpoint := "/info/all"

	tt := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "Valid request",
			method:         "GET",
			expectedStatus: http.StatusOK,
		},
	}

	url := testServer.URL + common.VERSION + testDepartmentService.URL + endpoint
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			req, err = http.NewRequest(tc.method, url, nil)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			rec := httptest.NewRecorder()

			fmt.Printf("testDepartmentService: %+v\n", testDepartmentService)

			testDepartmentService.handlerDepartments(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatus, res.StatusCode)
		})
	}
}
