package simpleresty

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
)

type Response struct {
	URL        string
	Method     string
	Status     string
	StatusCode int
	Body       string
}

func checkResponse(resp *resty.Response) (*Response, error) {
	path, _ := url.QueryUnescape(resp.Request.URL)
	r := &Response{Status: resp.Status(), StatusCode: resp.StatusCode(),
		Body: string(resp.Body()), URL: path, Method: resp.Request.Method}

	// If response is the below, return.
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return r, nil
	}

	// Otherwise, return an error
	return r, fmt.Errorf("%s %s: %d %s", r.Method, r.URL, r.StatusCode, r.Body)
}
