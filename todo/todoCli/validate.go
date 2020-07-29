package main

import (
	"errors"

	"github.com/mkideal/cli"
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
