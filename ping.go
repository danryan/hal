package main

import (
	"gobot/gobot"
)

var respondHandler = gobot.Respond(`(?i)respond$`, func(res *gobot.Response) error {
	res.Reply("responding")
	return nil
})

var hearHandler = gobot.Hear(`(?i)hear$`, func(res *gobot.Response) error {
	res.Send("hearing")
	return nil
})

var enterHandler = gobot.Enter(func(res *gobot.Response) error {
	res.Send("ACK")
	return nil
})
