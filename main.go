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
	}

	for _, l := range links {
		checkLink(l)
	}
}

func checkLink(l string) {
	_, e := http.Get(l)
	if e != nil {
		fmt.Println(l, "might be down.")
		return
	}

	fmt.Println(l, "is up.")
}
