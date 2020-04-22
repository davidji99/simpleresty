package simpleresty

import "github.com/go-resty/resty/v2"

// Request represents a HTTP request.
//
// This struct embeds resty.Request
type Request struct {
	*resty.Request
}

func (c *Client) NewRequest() *Request {
	return &Request{Request: c.R()}
}