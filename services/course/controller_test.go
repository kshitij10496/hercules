package course

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"

	"github.com/kshitij10496/hercules/common"
)

var testServer *httptest.Server
var testServiceCourse serviceCourse

func setup() error {

	testServiceCourse = serviceCourse{
		Name:   "service-course",
		URL:    "/course",
		Router: common.NewSubRouter(Routes),
	}

	// Grab $HERCULES_DATABASE from env
	databaseURL := os.Getenv("HERCULES_DATABASE")
	if databaseURL == "" {
		log.Fatal("Missing: HERCULES_DATABASE environment variable")
	}
	testServer = httptest.NewServer(&testServiceCourse)
	return testServiceCourse.ConnectDB(databaseURL)

}

func teardown() {
	testServiceCourse.CloseDB()
	testServer.Close()
}

func Test_handlerCoursesFromDepartment(t *testing.T) {
	err := setup()
	defer teardown()

	if !assert.NoError(t, err) {
		t.Fatalf("unable to setup test: %v\n", err)
	}

	endpoint := "/info/department"

	tt := []struct {
		name           string
		method         string
		code           string
		expectedStatus int
	}{
		{
			name:           "Valid Department MA",
			method:         "GET",
			code:           "MA",
			expectedStatus: http.StatusOK,
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := testServer.URL + common.VERSION + testServiceCourse.URL + endpoint
			var req *http.Request
			req, err = http.NewRequest(tc.method, url, nil)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			if tc.code != "" {
				req = mux.SetURLVars(req, map[string]string{"code": tc.code})
				log.Println("UPDATED req:", req.URL.String())
			}
			rec := httptest.NewRecorder()

			testServiceCourse.handlerCoursesFromDepartment(rec, req)
			res := rec.Result()
			assert.Equal(t, res.StatusCode, tc.expectedStatus)
			defer res.Body.Close()

			var course responseCourses
			decoder := json.NewDecoder(res.Body)
			err = decoder.Decode(&course)

			if res.StatusCode == http.StatusOK {
				assert.Equal(t, err, nil)
			}
		})
	}
}
