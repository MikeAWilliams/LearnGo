package main

import (
	"fmt"
	"net/http"
)

func checkPage(url string) {
	resp, err := http.Get(url)
	if nil != err {
		fmt.Println("Error reading "+url, err)
	}
	defer resp.Body.Close()
	if 200 != resp.StatusCode {
		fmt.Println("Error reading "+url+" response was ", resp.StatusCode)
	}
	fmt.Println(url + " is ok")
}

func main() {
	pages := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, page := range pages {
		checkPage(page)
	}
}
