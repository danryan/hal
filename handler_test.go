package hal_test

import (
	"github.com/danryan/hal"
)

func ExampleBasicHandler_hear() {
	res := hal.Response{
		Match: []string{},
	}
	h := &hal.BasicHandler{
		Method:  hal.HEAR,
		Pattern: `echo (.+)`,
		Usage:   "echo <string> - repeats <string> back",
		Run: func(res *hal.Response) error {
			res.Send(res.Match[1])
		},
	}
	// output:
	// > echo foo bar baz
	// foo bar baz
}

func ExampleBasicHandler_respond() {
	&BasicHandler{
		Method:  hal.RESPOND,
		Pattern: `(?i)ping`, // (?i) is a flag that makes the match case insensitive
		Usage:   `hal ping - replies with "PONG"`,
		Run: func(res *hal.Response) error {
			res.Send("PONG")
		},
	}

}
