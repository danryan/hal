package main

import (
	"hal/hal"
)

var respondHandler = hal.Respond(`(?i)respond$`, func(res *hal.Response) error {
	res.Reply("responding")
	return nil
})

var hearHandler = hal.Hear(`(?i)hear$`, func(res *hal.Response) error {
	res.Send("hearing")
	return nil
})

var enterHandler = hal.Enter(func(res *hal.Response) error {
	res.Send("ACK")
	return nil
})
