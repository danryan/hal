package hal

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

// NewAdapter returns a new Adapter object
func NewAdapter(name string) (Adapter, error) {
	switch name {
	case "shell":
		return &ShellAdapter{
			out: bufio.NewWriter(os.Stdout),
			in:  bufio.NewReader(os.Stdin),
		}, nil
	case "slack":
		return &SlackAdapter{
			token:    os.Getenv("HAL_SLACK_TOKEN"),
			team:     os.Getenv("HAL_SLACK_TEAM"),
			channels: os.Getenv("HAL_SLACK_CHANNELS"),
			mode: func() string {
				mode := os.Getenv("HAL_SLACK_MODE")
				if mode == "" {
					return "blacklist"
				}
				return mode
			}(),
		}, nil
	default:
		return nil, errors.New("invalid adapter name")
	}

}
