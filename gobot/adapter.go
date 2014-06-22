package gobot

import (
	"bufio"
	"errors"
	"os"
)

// Adapter interface
type Adapter interface {
	Run() error
	Stop() error

	Send(*Response, ...string) error
	Emote(*Response, ...string) error
	Reply(*Response, ...string) error
	Topic(*Response, ...string) error
	Play(*Response, ...string) error

	Receive(*Message) error

	Robot() *Robot
	SetRobot(*Robot)
}

func NewAdapter(name string) (Adapter, error) {
	switch name {
	case "shell":
		return &ShellAdapter{
			out: bufio.NewWriter(os.Stdout),
			in:  bufio.NewReader(os.Stdin),
		}, nil
	default:
		return nil, errors.New("invalid adapter name")
	}

}
