package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func dealWithError(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}

func printHttpResponse(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if dealWithError(err) {
		return
	}
	pretty := &bytes.Buffer{}
	indentErr := json.Indent(pretty, body, "", "  ")
	if dealWithError(indentErr) {
		return
	}
	fmt.Println(pretty.String())
}

func getURI(title string) string {
	uri := "http://localhost:8000/api/v1/items"
	if len(title) > 0 {
		uri += "/" + title
	}
	return uri
}

func performGet(argv argT) {
	uri := getURI(argv.Title)
	fmt.Printf("Doing the get on %v\n", uri)

	resp, err := http.Get(uri)

	if dealWithError(err) {
		return
	}
	printHttpResponse(resp)
}

func performPost(argv argT) {
	uri := getURI(argv.Title)
	fmt.Printf("Doing the post on %v\n", uri)

	bodyMap := make(map[string]string)
	bodyMap["Description"] = argv.Description
	bodyJson, marshalErr := json.Marshal(bodyMap)
	if dealWithError(marshalErr) {
		return
	}

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(bodyJson))
	if dealWithError(err) {
		return
	}
	printHttpResponse(resp)
}

func performPut(argv argT) {
	uri := getURI(argv.Title)
	fmt.Printf("Doing the put on %v\n", uri)
}

func performDelete(argv argT) {
	uri := getURI(argv.Title)
	fmt.Printf("Doing the delete on %v\n", uri)
}
