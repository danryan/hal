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
	in   *bufio.Reader
	out  *bufio.Writer
	quit chan bool
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
		s := res.UserID() + `: ` + str
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
	a.Robot.Receive(msg)
	return nil
}

// Run executes the adapter run loop
func (a *ShellAdapter) Run() error {
	a.run()
	prompt()

	go func() {
		for {
			line, _, err := a.in.ReadLine()
			message := a.newMessage(string(line))
			if err != nil {
				fmt.Println("error:", err)
			}
			a.Receive(message)
			prompt()
		}
	}()

	<-a.quit
	return nil
}

// Stop the adapter
func (a *ShellAdapter) Stop() error {
	a.preStop()
	a.quit <- true
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

	if _, err := a.out.WriteString(msg); err != nil {
		return err
	}

	if err := a.out.Flush(); err != nil {
		return err
	}

	return nil
}

func (a *ShellAdapter) Name() string {
	return "shell"
}
