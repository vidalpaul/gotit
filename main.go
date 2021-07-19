package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://paulovidal.com",
		"http://golang.org",
		"http://github.com",
		"http://stackoverflow.com",
		"http://npmjs.com",
		"http://docs.python.org",
		"https://rust-lang.org/",
		"https://julialang.org/",
		"https://dart.dev/",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	fmt.Println(<-c)
}

func checkLink(l string, c chan string) {
	_, e := http.Get(l)
	if e != nil {
		fmt.Println(l, "might be down.")
		c <- "Might be down"
		return
	}

	fmt.Println(l, "is up.")
	c <- "Is up"
}
