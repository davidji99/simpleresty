package simpleresty

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseProxyForServer_Creds(t *testing.T) {
	testURL := "http://BOB.BUILDER:ACOMPLICATEDPASSWORD@company.com:8080"
	expected := "company.com"

	assert.Equal(t, expected, parseProxyURLForDomain(testURL))
}

func TestParseProxyForServer_NoCreds(t *testing.T) {
	testURL := "http://company.com:8080"
	expected := "company.com"

	assert.Equal(t, expected, parseProxyURLForDomain(testURL))
}

func TestParseProxyForServer_NoCredsHTTPS(t *testing.T) {
	testURL := "https://company.com:8080"
	expected := "company.com"

	assert.Equal(t, expected, parseProxyURLForDomain(testURL))
}

func TestGetNoProxyDomains_NoneSet(t *testing.T) {
	domains, isPresent := getNoProxyDomains()

	assert.Equal(t, 0, len(domains))
	assert.Equal(t, false, isPresent)
}

func TestGetNoProxyDomains_SingleSet_NO_PROXY(t *testing.T) {
	_ = os.Setenv("NO_PROXY", "somedomain.com")
	defer os.Unsetenv("NO_PROXY")

	domains, isPresent := getNoProxyDomains()

	assert.Equal(t, []string{"somedomain.com"}, domains)
	assert.Equal(t, true, isPresent)
}

func TestGetNoProxyDomains_SingleSet_no_proxy(t *testing.T) {
	_ = os.Setenv("no_proxy", "somedomain.com")
	defer os.Unsetenv("no_proxy")

	domains, isPresent := getNoProxyDomains()

	assert.Equal(t, []string{"somedomain.com"}, domains)
	assert.Equal(t, true, isPresent)
}

func TestGetNoProxyDomains_MultipleSet(t *testing.T) {
	_ = os.Setenv("NO_PROXY", "somedomain.com, somedomain2.com")
	defer os.Unsetenv("NO_PROXY")

	domains, isPresent := getNoProxyDomains()

	assert.Equal(t, []string{"somedomain.com", "somedomain2.com"}, domains)
	assert.Equal(t, true, isPresent)
}

func TestGetNoProxyDomains_WildcardMultipleSet(t *testing.T) {
	_ = os.Setenv("NO_PROXY", "*.somedomain.com, somedomain2.com")
	defer os.Unsetenv("NO_PROXY")

	domains, isPresent := getNoProxyDomains()

	assert.Equal(t, []string{"somedomain.com", "somedomain2.com"}, domains)
	assert.Equal(t, true, isPresent)
}

func TestGetNoProxyDomains_WildcardMultipleSet2(t *testing.T) {
	_ = os.Setenv("NO_PROXY", "*.somedomain.*, somedomain2.com")
	defer os.Unsetenv("NO_PROXY")

	domains, isPresent := getNoProxyDomains()

	assert.Equal(t, []string{"somedomain2.com"}, domains)
	assert.Equal(t, true, isPresent)
}

func TestGetNoProxyDomains_MultipleSetWithOneInvalid(t *testing.T) {
	_ = os.Setenv("NO_PROXY", "somedomain, test.somedomain2.com")
	defer os.Unsetenv("NO_PROXY")

	domains, isPresent := getNoProxyDomains()

	assert.Equal(t, []string{"test.somedomain2.com"}, domains)
	assert.Equal(t, true, isPresent)
}
