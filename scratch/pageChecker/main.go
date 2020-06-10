package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func checkPage(url string, c chan string) {
	resp, err := http.Get(url)
	if nil != err {
		c <- "Error reading " + url + err.Error()
		return
	}
	defer resp.Body.Close()
	if 200 != resp.StatusCode {
		c <- "Error reading " + url + " response was " + strconv.Itoa(resp.StatusCode)
		return
	}
	c <- url + " is ok"
}

func main() {
	pages := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, page := range pages {
		go checkPage(page, c)
	}
	fmt.Println(<-c)
}
