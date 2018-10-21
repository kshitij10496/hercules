package department

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kshitij10496/hercules/common"
	"github.com/stretchr/testify/assert"
)

func Test_GetDepartments_ete(t *testing.T) {
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
			expectedStatus: http.StatusOK,
		},
	}

	url := testServer.URL + common.VERSION + testDepartmentService.URL + endpoint
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(url)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, res.StatusCode)
		})
	}
}
