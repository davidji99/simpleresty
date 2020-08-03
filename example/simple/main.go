package main

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"log"
)

type GithubAPIResponse struct {
	CurrentUserURL string `json:"current_user_url,omitempty"`
}

type Yugioh struct {
	simpleresty.Client

	ProjectToken string
}

func UserAgent(v string) simpleresty.Option {
	return func(i interface{}) error {
		y := i.(Yugioh)
		y.ProjectToken = v
		return nil
	}
}

func y1(opts ...simpleresty.Option) {

	x := Yugioh{}

	_ = simpleresty.ParseOptions(x, opts...)

	log.Println("asdasd")
	log.Printf(x.ProjectToken)
}

func main() {
	c := simpleresty.New()

	y1(UserAgent("asdasdasd!!!"))

	var result *GithubAPIResponse
	response, getErr := c.Get("https://api.github.com", &result, nil)
	if getErr != nil {
		panic(getErr)
	}

	fmt.Println(response.StatusCode)   // Returns 200
	fmt.Println(result.CurrentUserURL) // Returns 'https://api.github.com/user'
}
