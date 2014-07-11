package shell

import (
	"bufio"
	"fmt"
	"github.com/danryan/hal"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	hal.RegisterAdapter("shell", New)
}

type adapter struct {
	hal.BasicAdapter
	in   *bufio.Reader
	out  *bufio.Writer
	quit chan bool
}

// New returns an initialized adapter
func New(r *hal.Robot) (hal.Adapter, error) {
	a := &adapter{
		out:  bufio.NewWriter(os.Stdout),
		in:   bufio.NewReader(os.Stdin),
		quit: make(chan bool),
	}
	a.SetRobot(r)
	return a, nil
}

// Send sends a regular response
func (a *adapter) Send(res *hal.Response, strings ...string) error {
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
func (a *adapter) Reply(res *hal.Response, strings ...string) error {
	for _, str := range strings {
		s := res.UserName() + `: ` + str
		err := a.writeString(s)
		if err != nil {
			log.Println("error: ", err)
			return err
		}
	}

	return nil
}

// Emote performs an emote
func (a *adapter) Emote(res *hal.Response, strings ...string) error {
	return nil
}

// Topic sets the topic
func (a *adapter) Topic(res *hal.Response, strings ...string) error {
	return nil
}

// Play plays a sound
func (a *adapter) Play(res *hal.Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *adapter) Receive(msg *hal.Message) error {
	a.Robot.Receive(msg)
	return nil
}

// Run executes the adapter run loop
func (a *adapter) Run() error {
	prompt()

	go func() {
		for {
			line, _, err := a.in.ReadLine()
			message := a.newMessage(string(line))

			if err != nil {
				if err == io.EOF {
					break
					// a.Robot.signalChan <- syscall.SIGTERM
				}
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
func (a *adapter) Stop() error {
	a.quit <- true
	return nil
}

func prompt() {
	fmt.Print("> ")
}

// func newMessage(text string) *Message {
func (a *adapter) newMessage(text string) *hal.Message {
	return &hal.Message{
		ID:   "local-message",
		User: hal.User{ID: "1", Name: "shell"},
		Room: "shell",
		Text: text,
	}
}

func (a *adapter) writeString(str string) error {
	msg := fmt.Sprintf("%s\n", strings.TrimSpace(str))

	if _, err := a.out.WriteString(msg); err != nil {
		return err
	}

	if err := a.out.Flush(); err != nil {
		return err
	}

	return nil
}
