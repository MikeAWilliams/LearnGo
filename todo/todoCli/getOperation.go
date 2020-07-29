package main

const (
	GET     = 1
	POST    = 2
	PUT     = 3
	DELETE  = 4
	UNKNOWN = 5
)

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
