package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ifErrorlogAndDie(err error) {
	if nil != err {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args
	if 2 != len(args) {
		log.Fatal("Usage: fileToStdout fileName")
	}

	fileName := args[1]
	fmt.Println("Printing " + fileName)
	fileBytes, err := ioutil.ReadFile(fileName)
	ifErrorlogAndDie(err)

	fmt.Print(string(fileBytes))
}
