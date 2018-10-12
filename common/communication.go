package common

import (
	"fmt"
	"net/http"
	"time"
)

func SendToService(service, method, endpoint string, body interface{}) (*http.Response, error) {
	// Create a new HTTP client
	client := http.Client{
		Timeout: time.Second * 5,
	}

	url := fmt.Sprintf("http://localhost:8080%s/%s%s", VERSION, service, endpoint)
	fmt.Println("URL:", url)
	switch method {
	case "GET":
		return client.Get(url)
	default:
		return nil, ErrNotImplemented
	}
}
