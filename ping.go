package main

import (
	"gobot/gobot"
)

// type pingHandler struct{}

// func (h *pingHandler) Pattern() string {
// 	return `/ping$/i`
// }

// func (h *pingHandler) Handle(msg *gobot.Message) error {
// 	msg.Send("PONG")
// 	return nil
// }

// var pingCommand = func(msg *gobot.Message) {
// 	// command := gobot.Command(`/ping$/i`)
// 	msg.Send("PONG")
// }

// func pingCmd(msg *gobot.Message) (string, gobot.Handler) {
// 	return `/ping$/i`, func(msg *gobot.Message) {
// 		msg.Send("PONG")
// 	}(msg)
// }

var pingCmd = &gobot.Command{"respond", `/ping$/i`, func(msg *gobot.Message) {
	msg.Send("PONG")
}}

var synCmd = gobot.Respond(`/syn$/i`, func(msg *gobot.Message) {
	msg.Send("ACK")
})
