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
	String() string
	Name() string

	Send(*Response, ...string) error
	Emote(*Response, ...string) error
	Reply(*Response, ...string) error
	Topic(*Response, ...string) error
	Play(*Response, ...string) error

	Receive(*Message) error
}

// NewAdapter returns a new Adapter object
func NewAdapter(robot *Robot) (Adapter, error) {
	switch robot.AdapterName {
	case "shell":
		return newShellAdapter(robot)
	case "slack":
		return newSlackAdapter(robot)
	default:
		return nil, errors.New("invalid adapter name")
	}
}

func newSlackAdapter(robot *Robot) (Adapter, error) {
	slack := &SlackAdapter{
		token:          os.Getenv("HAL_SLACK_TOKEN"),
		team:           os.Getenv("HAL_SLACK_TEAM"),
		channels:       os.Getenv("HAL_SLACK_CHANNELS"),
		mode:           GetenvDefault("HAL_SLACK_CHANNELMODE", "blacklist"),
		botname:        GetenvDefault("HAL_SLACK_BOTNAME", robot.Name),
		iconEmoji:      os.Getenv("HAL_SLACK_ICON_EMOJI"),
		ircEnabled:     GetenvDefaultBool("HAL_SLACK_IRC_ENABLED", false),
		ircPassword:    os.Getenv("HAL_SLACK_IRC_PASSWORD"),
		responseMethod: GetenvDefault("HAL_SLACK_RESPONSE_METHOD", "http"),
	}
	slack.SetRobot(robot)
	return slack, nil
}

func newShellAdapter(robot *Robot) (Adapter, error) {
	shell := &ShellAdapter{
		out:  bufio.NewWriter(os.Stdout),
		in:   bufio.NewReader(os.Stdin),
		quit: make(chan bool),
	}
	shell.SetRobot(robot)
	return shell, nil
}
