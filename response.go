package simpleresty

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
)

// Response represents the response after executing a HTTP request.
type Response struct {
	RequestURL    string
	RequestMethod string
	Status        string
	StatusCode    int
	Body          string
	RawHTTP       *http.Response
}

func checkResponse(resp *resty.Response) (*Response, error) {
	path, _ := url.QueryUnescape(resp.Request.URL)
	r := &Response{Status: resp.Status(), StatusCode: resp.StatusCode(),
		Body: string(resp.Body()), RawHTTP: resp.RawResponse, RequestURL: path,
		RequestMethod: resp.Request.Method}

	// If response is any of the below, return early.
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return r, nil
	}

	// Otherwise, return the response along with the error.
	return r, fmt.Errorf("%s %s: %d %s", r.RequestMethod, r.RequestURL, r.StatusCode, r.Body)
}
