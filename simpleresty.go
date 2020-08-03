package simpleresty

import (
	"os"
)

var (
	proxyVars = []string{"HTTPS_PROXY", "https_proxy", "HTTP_PROXY", "http_proxy"}
)

// Client represents a simpleresty client.
type Client struct {
	*HttpClient
}

// New function creates a new SimpleResty client.
func New() *Client {
	c := &Client{HttpClient: &HttpClient{}}

	determineSetProxy(c)

	return c
}

// determineSetProxy checks if any proxy variables are defined in the environment.
// If so, set the first occurrence and exit the loop.
func determineSetProxy(c *Client) {
	for _, v := range proxyVars {
		proxyUrl := os.Getenv(v)
		if proxyUrl != "" {
			c.SetProxy(proxyUrl)
			break
		}
	}
}
