package simpleresty

import (
	"os"
	"regexp"
	"strings"
)

var (
	proxyVars   = []string{"HTTPS_PROXY", "https_proxy", "HTTP_PROXY", "http_proxy"}
	noProxyVars = []string{"NO_PROXY", "no_proxy"}
)

func parseProxyForServer(proxyURL string) string {
	var serverRaw string

	// Split proxyURL by '@' to account for username/password in the URL
	proxyURLSplitted := strings.Split(proxyURL, "@")

	if len(proxyURLSplitted) == 1 {
		// If no username/password in URL, return proxyURLSplitted's zero index
		serverRaw = strings.ToLower(proxyURLSplitted[0])
	} else {
		// Take the 1st index value
		serverRaw = proxyURLSplitted[1]
	}

	// Strip out the protocol
	regex := regexp.MustCompile(`http[s]?://`)
	serverRaw = regex.ReplaceAllString(serverRaw, "")

	// Split by colon to separate server from PORT and get the zero index
	server := strings.Split(serverRaw, ":")[0]

	return strings.ToLower(server)
}

// getNoProxyDomains fetches no proxy variables from the environment and parses each variable value for domains.
//
// Returns a String array of domain names (default empty) and a Boolean for if there are any no proxy domains.
func getNoProxyDomains() ([]string, bool) {
	noProxyDomains := make([]string, 0)

	for _, v := range noProxyVars {
		noProxyDomainString, isSet := os.LookupEnv(v)
		if !isSet || noProxyDomainString == "" {
			continue
		}

		// Split by comma
		noProxyDomainsRaw := strings.Split(noProxyDomainString, ",")

		// Iterate through each URL and format properly
		for _, domainRaw := range noProxyDomainsRaw {
			// Remove leading and trailing whitespaces
			domainRaw = strings.TrimSpace(domainRaw)

			// Strip out any wildcard notation, `*.`
			regexWC1 := regexp.MustCompile(`\*\.`)
			domainRaw = regexWC1.ReplaceAllString(domainRaw, "")

			// Strip out any wildcard notation, `.*`
			regexWC2 := regexp.MustCompile(`\.\*`)
			domainRaw = regexWC2.ReplaceAllString(domainRaw, "")

			// Make sure noProxyURLRaw is a valid domain, such as example.info|com|net|etc...
			validDomainFormatRegex := regexp.MustCompile(`\S+\.\S+`)
			isValidDomain := validDomainFormatRegex.MatchString(domainRaw)

			if isValidDomain {
				noProxyDomains = append(noProxyDomains, strings.ToLower(domainRaw))
			}
		}
	}

	return noProxyDomains, len(noProxyDomains) > 0
}

func getProxyURL() *string {
	for _, v := range proxyVars {
		proxyURL, isVarSet := os.LookupEnv(v)
		if !isVarSet {
			continue
		}

		if proxyURL != "" {
			url := strings.TrimSpace(proxyURL)
			return &url
		}
	}

	return nil
}
