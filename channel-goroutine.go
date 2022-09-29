package main

import (
	"fmt"
	"net/http"
	"time"
)

func testChannel() {
	links := []string{
		"https://google.com",
		"https://golang.org",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 2)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // network (blocking) call
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	c <- link
	fmt.Println(link, " is up")
}
