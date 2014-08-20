package handler

import (
	"github.com/danryan/hal"
)

// Echo is an example of a simple handler.
var Echo = hal.Respond(`echo (.+)`, func(res *hal.Response) error {
	return res.Reply(res.Match[1])
})
