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

	determineSetProxy(c)

	assert.Equal(t, true, c.IsProxySet())

	_ = os.Setenv("https_proxy", "")
}

func TestDetermineSetProxy_HttpBasic(t *testing.T) {
	c := &Client{Client: resty.New()}

	setErr := os.Setenv("http_proxy", "some.url:8080")
	assert.Nil(t, setErr)

	determineSetProxy(c)

	assert.Equal(t, true, c.IsProxySet())

	_ = os.Setenv("http_proxy", "")
}

func TestDetermineSetProxy_NoneSet(t *testing.T) {
	c := &Client{Client: resty.New()}

	setErr := os.Setenv("https_proxy123", "some.url:8080")
	assert.Nil(t, setErr)

	determineSetProxy(c)

	assert.Equal(t, false, c.IsProxySet())

	_ = os.Setenv("https_proxy123", "")
}
