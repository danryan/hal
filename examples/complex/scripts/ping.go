package scripts

import "github.com/danryan/hal"

var PingHandler = hal.Hear(`ping`, func(res *hal.Response) error {
	return res.Send("PONG")
})
