package main

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"golang.org/x/oauth2/github"
	"log"
)

func main() {
	t, err := simpleresty.OAuth("",
		"", github.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
}