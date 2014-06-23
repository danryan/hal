package hal

import (
	_ "fmt"
	_ "log"
	_ "net/http"
	_ "os"
	_ "strings"
)

// SlackAdapter struct
type SlackAdapter struct {
	BasicAdapter // includes robot field
	token        string
	team         string
	mode         string
	channels     string //[]string
	linkNames    []string
}

// Send sends a regular response
func (a *SlackAdapter) Send(res *Response, strings ...string) error {
	for _, str := range strings {
		_ = str
	}

	return nil
}

// Reply sends a direct response
func (a *SlackAdapter) Reply(res *Response, strings ...string) error {
	for _, str := range strings {
		_ = str
	}

	return nil
}

// Emote is not implemented.
func (a *SlackAdapter) Emote(res *Response, strings ...string) error {
	return nil
}

// TODO: implement
// Topic sets the topic
func (a *SlackAdapter) Topic(res *Response, strings ...string) error {
	for _, str := range strings {
		_ = str
	}
	return nil
}

// Play is not implemented.
func (a *SlackAdapter) Play(res *Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *SlackAdapter) Receive(msg *Message) error {
	a.robot.Receive(msg)
	return nil
}

// Run starts the adapter
func (a *SlackAdapter) Run() error {
	a.run()

	return nil
}

// Stop shuts down the adapter
func (a *SlackAdapter) Stop() error {
	a.stop()
	return nil
}

// TODO: implement
// func (a *SlackAdapter) newMessage(text string) *Message {
// 	return &Message{
// 		ID:   "local-message",
// 		User: &User{ID: "shell"},
// 		Room: "shell",
// 		Text: text,
// 	}
// }
