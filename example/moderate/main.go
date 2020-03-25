package main

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// GithubGist represents a Github gist.
type GithubGist struct {
	URL         string `json:"url,omitempty"`
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

// GithubUser represents a Github user.
type GithubUser struct {
	Login     string `json:"login,omitempty"`
	ID        int64  `json:"id,omitempty"`
	SiteAdmin bool   `json:"site_admin,omitempty"`
}

// ListGistOpts represents the query parameters available when listing all lists
type ListGistOpts struct {
	Since string `url:"since,omitempty"`
}

func main() {
	c := simpleresty.New()

	// Set base API url so you don't need to prepend it for each request.
	c.SetBaseURL("https://api.github.com")

	// Define query parameters
	opts := ListGistOpts{Since: "2020-03-24T12:37:57+0000"}

	// Construct the first request URL
	urlStr1, urlStr1Err := c.RequestURLWithQueryParams("/gists/public", opts)
	if urlStr1Err != nil {
		panic(urlStr1Err)
	}

	// Execute the request
	var gist []*GithubGist
	response1, getErr1 := c.Get(urlStr1, &gist, nil)
	if getErr1 != nil {
		panic(getErr1)
	}

	fmt.Println("Results of the second API request: ")
	fmt.Println(response1.StatusCode) // Returns 200
	fmt.Println(len(gist))            // should return more than one
	fmt.Println(gist[0].ID)           // Returns a valid gist ID

	fmt.Println("")

	// Construct the second request URL
	urlStr2 := c.RequestURL("/users/%s", "defunct")

	// Execute the request
	var result *GithubUser
	response2, getErr2 := c.Get(urlStr2, &result, nil)
	if getErr2 != nil {
		panic(getErr2)
	}

	fmt.Println("Results of the first API request: ")
	fmt.Println(response2.StatusCode) // Returns 200
	fmt.Println(result.ID)            // Returns a valid user ID
	fmt.Println(result.Login)         // Returns a valid user Login
	fmt.Println(result.SiteAdmin)     // Returns true or false
}
