package scripts

import "github.com/danryan/hal"

var SynHandler = hal.Hear(`syn`, func(res *hal.Response) error {
	return res.Send("ACK")
})
