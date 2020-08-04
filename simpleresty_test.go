package simpleresty

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDetermineSetProxy_HttpsBasic(t *testing.T) {
	c := &Client{Client: resty.New()}

	setErr := os.Setenv("https_proxy", "some.url:8080")
	assert.Nil(t, setErr)

	c.determineSetProxy()

	assert.Equal(t, true, c.IsProxySet())

	_ = os.Setenv("https_proxy", "")
}

func TestDetermineSetProxy_HttpsBasicWithNoProxy(t *testing.T) {
	c := &Client{Client: resty.New()}

	c.SetBaseURL("somedirecturl.com")

	setErr := os.Setenv("https_proxy", "some.url:8080")
	setErr = os.Setenv("no_proxy", "somedirecturl.com")

	assert.Nil(t, setErr)

	c.determineSetProxy()

	assert.Equal(t, true, c.IsProxySet())

	_ = os.Unsetenv("https_proxy")
	_ = os.Unsetenv("no_proxy")

}

func TestDetermineSetProxy_HttpBasic(t *testing.T) {
	c := &Client{Client: resty.New()}

	setErr := os.Setenv("http_proxy", "some.url:8080")
	assert.Nil(t, setErr)

	c.determineSetProxy()

	assert.Equal(t, true, c.IsProxySet())

	_ = os.Setenv("http_proxy", "")
}

func TestDetermineSetProxy_NoneSet(t *testing.T) {
	c := &Client{Client: resty.New()}

	setErr := os.Setenv("https_proxy123", "some.url:8080")
	assert.Nil(t, setErr)

	c.determineSetProxy()

	assert.Equal(t, false, c.IsProxySet())

	_ = os.Setenv("https_proxy123", "")
}
