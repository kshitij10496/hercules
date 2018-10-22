package department

import (
	"net/http"
	"testing"

	"github.com/kshitij10496/hercules/common"
	"github.com/stretchr/testify/assert"
)

func Test_Routing_GetDepartments(t *testing.T) {
	// Setup tests
	testDepartmentService := newFakeServiceDepartment(true)
	testServer, err := setup(testDepartmentService)
	assert.NoError(t, err)

	// Teardown tests
	defer testServer.Close()
	defer teardown(testDepartmentService)
	// TODO: Handler error during teardown

	tt := []struct {
		name           string
		method         string
		endpoint       string
		expectedStatus int
	}{
		{
			name:           "Valid request",
			method:         "GET",
			endpoint:       "/info/all",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid request endpoint",
			method:         "GET",
			endpoint:       "/info/al",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Invalid request method",
			method:         "POST",
			endpoint:       "info/all",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := testServer.URL + common.VERSION + testDepartmentService.URL + tc.endpoint
			var req *http.Request
			req, err = http.NewRequest(tc.method, url, nil)
			if err != nil {
				t.Fatal("cannot create request:", err)
			}
			client := testServer.Client()
			res, err := client.Do(req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, res.StatusCode)
		})
	}
}
