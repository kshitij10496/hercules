package common

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func SendToService(service, method, endpoint string, query url.Values, body interface{}) (*http.Response, error) {
	// Create a new HTTP client
	client := http.Client{
		Timeout: time.Second * 5,
	}

	u := fmt.Sprintf("http://localhost:8080%s/%s%s", VERSION, service, endpoint)
	if query != nil {
		u = fmt.Sprintf("http://localhost:8080%s/%s%s?%s", VERSION, service, endpoint, query.Encode())
	}

	fmt.Println("URL:", u)
	req, err := http.NewRequest(method, u, nil)
	if err != nil {

	}

	switch method {
	case "GET":
		return client.Do(req)
	default:
		return nil, ErrNotImplemented
	}
}
