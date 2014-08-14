package handler

import (
	"github.com/danryan/hal"
)

// Echo exports our echo handler
var Echo = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `echo (.+)`,
	Usage:   "echo STRING - echoes STRING",
	Run: func(res *hal.Response) error {
		return res.Reply(res.Match[1])
	},
}
