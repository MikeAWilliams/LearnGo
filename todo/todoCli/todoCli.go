package main

import (
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
