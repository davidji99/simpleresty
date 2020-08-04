package simpleresty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestNewWithBaseURL(t *testing.T) {
	c := NewWithBaseURL("https://base.url/api/v3")

	assert.NotNil(t, c)
	assert.Equal(t, "https://base.url/api/v3", c.baseURL)
}
