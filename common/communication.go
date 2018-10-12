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

	u := fmt.Sprintf("https://hercules-10496.herokuapp.com%s/%s%s", VERSION, service, endpoint)
	if query != nil {
		u = fmt.Sprintf("%s?%s", u, query.Encode())
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
