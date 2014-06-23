package hal

import (
	_ "fmt"
	"log"
	"net/http"
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
	a.preRun()
	// set up handlers
	a.robot.Router.HandleFunc("/hal/slack-webhook", slackHandler)
	a.postRun()

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
}

func slackHandler(w http.ResponseWriter, r *http.Request) {
	// send an empty response
	w.Write([]byte("yay"))
}

// Stop shuts down the adapter
func (a *SlackAdapter) Stop() error {
	a.stop()
	return nil
}

func (a *SlackAdapter) Name() string {
	return "slack"
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
