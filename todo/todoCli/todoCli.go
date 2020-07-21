package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Get    bool `cli:"g,get" usage:"set to do get"`
	Post   bool `cli:"p,post" usage:"set to do post"`
	Put    bool `cli:"u,put" usage:"set to do put"`
	Delete bool `cli:"d,delete" usage:"set to do delete"`

	Title       string `cli:"t,title" usage:"the title of the item in question"`
	Description string `cli:"e,description" usage:"the description of the item in question"`
	Complete    bool   `cli:"c,complete" usage:"the complete status of the item in question"`
}

const (
	GET     = 1
	POST    = 2
	PUT     = 3
	DELETE  = 4
	UNKNOWN = 5
)

func validateOperationExclusive(ctx *cli.Context) error {
	opCount := 0
	argv := ctx.Argv().(*argT)
	if argv.Get {
		opCount++
	}
	if argv.Post {
		opCount++
	}
	if argv.Put {
		opCount++
	}
	if argv.Delete {
		opCount++
	}
	if opCount == 1 {
		return nil
	}
	return errors.New("Exactly one operation is required")
}

func getOperation(argv argT) int {
	if argv.Get {
		return GET
	}
	if argv.Post {
		return POST
	}
	if argv.Put {
		return PUT
	}
	if argv.Delete {
		return DELETE
	}
	return UNKNOWN
}

func performGet(argv argT) {
	fmt.Printf("Doing the get")
}

func performPost(argv argT) {
	fmt.Printf("Doing the Post")
}

func performPut(argv argT) {
	fmt.Printf("Doing the Put")
}

func performDelete(argv argT) {
	fmt.Printf("Doing the Delete")
}

func performOperation(argv argT) {
	switch op := getOperation(argv); op {
	case GET:
		performGet(argv)
	case POST:
		performPost(argv)
	case PUT:
		performPut(argv)
	case DELETE:
		performDelete(argv)
	}
}

// Validate implements cli.Validator interface
func (argv *argT) Validate(ctx *cli.Context) error {
	return validateOperationExclusive(ctx)
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		performOperation(*argv)
		return nil
	}))
}
