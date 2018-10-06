package course

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server *httptest.Server

func setup() {
	server = httptest.NewServer(ServiceCourse.Router)
}

func teardown() {
	server.Close()
}

func Test_CoursesInfo(t *testing.T) {
	setup()
	defer teardown()

	tests := []struct {
		name             string
		method           string
		path             string
		body             io.Reader
		expectedResponse string
		expectedStatus   int
	}{
		{
			name:           "Valid Course Info",
			method:         "GET",
			path:           "/course/info",
			body:           strings.NewReader(`{"code":"NA61001"}`),
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		request, err := http.NewRequest(test.method, server.URL+test.path, test.body)
		assert.NoError(t, err)
		log.Println("Successfully created the request")
		log.Println("URL: ", server.URL+test.path)

		response, err := http.DefaultClient.Do(request)
		assert.NoError(t, err)

		actualResponse, err := ioutil.ReadAll(response.Body)
		assert.NoError(t, err)

		assert.Equal(t, test.expectedResponse, actualResponse, "Response")
		assert.Equal(t, test.expectedStatus, response.StatusCode, "Status Code")
	}

}
