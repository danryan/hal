package hal_test

import (
	"github.com/danryan/hal"
)

func ExampleHandler_hear() {
	res := hal.Response{
		Match: []string{},
	}
	h := &hal.Handler{
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

func ExampleHandler_respond() {
	&Handler{
		Method:  hal.RESPOND,
		Pattern: `(?i)ping`, // (?i) is a flag that makes the match case insensitive
		Usage:   `hal ping - replies with "PONG"`,
		Run: func(res *hal.Response) error {
			res.Send("PONG")
		},
	}

}
