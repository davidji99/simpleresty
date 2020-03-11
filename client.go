package simpleresty

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty.Client

	// baseURL for the API endpoint. Please include a trailing slash '/'.
	baseURL string
}

// Get executes a HTTP GET request.
func (c *Client) Get(url string, r, body interface{}) (*Response, error) {
	req := c.ConstructRequest(r, body)

	response, getErr := req.Get(url)
	if getErr != nil {
		return nil, getErr
	}

	return checkResponse(response)
}

// Post executes a HTTP POST request.
func (c *Client) Post(url string, r, body interface{}) (*Response, error) {
	req := c.ConstructRequest(r, body)

	response, postErr := req.Post(url)
	if postErr != nil {
		return nil, postErr
	}

	return checkResponse(response)
}

// Put executes a HTTP PUT request.
func (c *Client) Put(url string, r, body interface{}) (*Response, error) {
	req := c.ConstructRequest(r, body)

	response, putErr := req.Put(url)
	if putErr != nil {
		return nil, putErr
	}

	return checkResponse(response)
}

// Patch executes a HTTP PATCH request.
func (c *Client) Patch(url string, r, body interface{}) (*Response, error) {
	req := c.ConstructRequest(r, body)

	response, patchErr := req.Patch(url)
	if patchErr != nil {
		return nil, patchErr
	}

	return checkResponse(response)
}

// Delete executes a HTTP DELETE request.
func (c *Client) Delete(url string, r, body interface{}) (*Response, error) {
	req := c.ConstructRequest(r, body)

	response, deleteErr := req.Delete(url)
	if deleteErr != nil {
		return nil, deleteErr
	}

	return checkResponse(response)
}

// ConstructRequest creates a new request.
func (c *Client) ConstructRequest(r, body interface{}) *resty.Request {
	req := c.R().SetBody(body)

	if r != nil {
		req.SetResult(r)
	}

	return req
}

// RequestURL appends the template argument to the base URL and returns the full request URL.
func (c *Client) RequestURL(template string, args ...interface{}) string {
	if len(args) == 1 && args[0] == "" {
		return c.baseURL + template
	}
	return c.baseURL + fmt.Sprintf(template, args...)
}

// RequestURLWithQueryParams first constructs the request URL and then appends any URL encoded query parameters.
//
// This function operates nearly the same as RequestURL
func (c *Client) RequestURLWithQueryParams(url string, opts ...interface{}) (string, error) {
	u := c.RequestURL(url)
	return AddQueryParams(u, opts)
}

// SetBaseURL sets the base url for the client.
func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
}
