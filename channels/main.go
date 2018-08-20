package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://google.com",
		"http://github.com",
		"http://golang.org",
		"http://not_an_website_com",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string)  {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " is down!")
		return
	}

	fmt.Println(link, " is up!")
}
