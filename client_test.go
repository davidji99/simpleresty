package simpleresty

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_RequestURL_BaseUrlSet(t *testing.T) {
	c := &Client{Client: resty.New(), baseURL: "https://sass.com/api/v1"}
	assert.Equal(t, "https://sass.com/api/v1/apps", c.RequestURL("/apps"))
}

func TestClient_RequestURL_BaseUrlSetWithArgs(t *testing.T) {
	c := &Client{Client: resty.New(), baseURL: "https://sass.com/api/v1"}
	assert.Equal(t, "https://sass.com/api/v1/apps/123", c.RequestURL("/apps/%s", "123"))
}

func TestClient_RequestURL_BaseUrlNotSet(t *testing.T) {
	c := &Client{Client: resty.New()}
	assert.Panics(t, func() { c.RequestURL("/apps") })
}

func TestClient_SetBaseURL(t *testing.T) {
	c := &Client{Client: resty.New()}
	c.SetBaseURL("https://www.google.com/api/v1")

	assert.Equal(t, "https://www.google.com/api/v1", c.baseURL)
}

func TestClient_RequestURLWithQueryParams(t *testing.T) {
	c := &Client{Client: resty.New(), baseURL: "https://sass.com/api/v1"}
	queryParams := struct {
		Since string `url:"since,omitempty"`
		Page  int    `url:"page,omitempty"`
	}{
		Since: "3days",
		Page:  4,
	}

	url, err := c.RequestURLWithQueryParams("/apps", queryParams)
	assert.Nil(t, err)
	assert.Equal(t, "https://sass.com/api/v1/apps?page=4&since=3days", url)
}

func TestClient_RequestURLWithNoQueryParams(t *testing.T) {
	c := &Client{Client: resty.New(), baseURL: "https://sass.com/api/v1"}
	url, err := c.RequestURLWithQueryParams("/apps")
	assert.Nil(t, err)
	assert.Equal(t, "https://sass.com/api/v1/apps", url)
}

func TestDetermineSetProxy_HttpsBasic(t *testing.T) {
	proxyURL := "some.url:8080"
	c := &Client{Client: resty.New()}
	c.proxyURL = &proxyURL

	setErr := os.Setenv("https_proxy", "some.url:8080")
	assert.Nil(t, setErr)

	c.determineSetProxy()

	assert.Equal(t, true, c.IsProxySet())

	_ = os.Setenv("https_proxy", "")
}

func TestDetermineSetProxy_WithNoProxySet(t *testing.T) {
	proxyURL := "some.url:8080"
	c := &Client{Client: resty.New()}
	c.proxyURL = &proxyURL
	c.noProxyDomains = []string{"somedirect0.url", "somedirecturl.com"}

	c.SetBaseURL("https://somedirecturl.com/api/1")

	c.determineSetProxy()

	assert.Equal(t, false, c.IsProxySet())
}

func TestDetermineSetProxy_NoneSet(t *testing.T) {
	c := &Client{Client: resty.New()}

	c.determineSetProxy()

	assert.Equal(t, false, c.IsProxySet())
}

func TestDetermineSetProxy_ProxyAlreadySet(t *testing.T) {
	c := &Client{Client: resty.New()}
	c.SetProxy("some.url:8080")

	c.determineSetProxy()

	assert.Equal(t, true, c.IsProxySet())
}
