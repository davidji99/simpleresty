package simpleresty

import (
	"github.com/go-resty/resty/v2"
)

// New function creates a new simpleresty client with base url set to empty string.
func New() *Client {
	c := &Client{Client: resty.New(), baseURL: ""}

	// Set proxy if applicable
	c.determineSetProxy()

	return c
}

// NewWithBaseURL creates a new simpleresty client with base url set.
func NewWithBaseURL(url string) *Client {
	c := &Client{Client: resty.New(), baseURL: url}

	// Set proxy if applicable
	c.determineSetProxy()

	return c
}
