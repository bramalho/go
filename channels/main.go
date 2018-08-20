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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<- c)
}

func checkLink(link string, c chan string)  {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " is down!")
		c <- "Down!"
		return
	}

	fmt.Println(link, " is up!")
	c <- "Up!"
	return
}
