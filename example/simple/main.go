package main

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

type GithubAPIResponse struct {
	CurrentUserURL string `json:"current_user_url,omitempty"`
}

func main() {
	c := simpleresty.New()

	var result *GithubAPIResponse
	response, getErr := c.Get("https://api.github.com", &result, nil)
	if getErr != nil {
		panic(getErr)
	}

	fmt.Println(response.StatusCode)   // Returns 200
	fmt.Println(result.CurrentUserURL) // Returns 'https://api.github.com/user'
}
