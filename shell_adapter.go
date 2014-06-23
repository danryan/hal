package hal

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// ShellAdapter struct
type ShellAdapter struct {
	BasicAdapter
	in      *bufio.Reader
	out     *bufio.Writer
	runChan chan bool
}

// Send sends a regular response
func (a *ShellAdapter) Send(res *Response, strings ...string) error {
	for _, str := range strings {
		err := a.writeString(str)
		if err != nil {
			log.Println("error: ", err)
			return err
		}
	}

	return nil
}

// Reply sends a direct response
func (a *ShellAdapter) Reply(res *Response, strings ...string) error {
	for _, str := range strings {
		s := res.Message.User.ID + `: ` + str
		err := a.writeString(s)
		if err != nil {
			log.Println("error: ", err)
			return err
		}
	}

	return nil
}

// Emote performs an emote
func (a *ShellAdapter) Emote(res *Response, strings ...string) error {
	return nil
}

// Topic sets the topic
func (a *ShellAdapter) Topic(res *Response, strings ...string) error {
	return nil
}

// Play plays a sound
func (a *ShellAdapter) Play(res *Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *ShellAdapter) Receive(msg *Message) error {
	a.robot.Receive(msg)
	return nil
}

// Run executes the adapter run loop
func (a *ShellAdapter) Run() error {
	a.run()
	run := true
	prompt()
	for run {
		line, _, err := a.in.ReadLine()
		message := a.newMessage(string(line))

		if err != nil {
			fmt.Println("error:", err)
		}
		a.Receive(message)
		prompt()

		select {
		case b := <-a.runChan:
			switch b {
			// case true:
			case false:
				run = false
			}
		default:
			continue
		}
	}

	return nil
}

// Stop the adapter
func (a *ShellAdapter) Stop() error {
	fmt.Println() // so we don't break up the log formatting :)
	a.preStop()
	a.runChan <- false
	a.postStop()
	return nil
}

func prompt() {
	fmt.Print("> ")
}

// func newMessage(text string) *Message {
func (a *ShellAdapter) newMessage(text string) *Message {
	return &Message{
		ID:   "local-message",
		User: &User{ID: "shell"},
		Room: "shell",
		Text: text,
	}
}

func (a *ShellAdapter) writeString(str string) error {
	msg := fmt.Sprintf("%s\n", strings.TrimSpace(str))

	_, err := a.out.WriteString(msg)
	if err != nil {
		return err
	}

	ferr := a.out.Flush()
	if ferr != nil {
		return err
	}

	return nil
}

func (a *ShellAdapter) Name() string {
	return "shell"
}
