package hal

import (
	"bufio"
	"errors"
	"os"
	"strings"
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

// var Adapters []Adapter

// func (a *Adapters) Add(adapter Adapter) {
// 	a = append(a, adapter)
// }

// NewAdapter returns a new Adapter object
func NewAdapter(robot *Robot) (Adapter, error) {
	switch robot.AdapterName {
	case "shell":
		return newShellAdapter(robot)
	case "slack":
		return newSlackAdapter(robot)
	case "irc":
		return newIRCAdapter(robot)
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

func newHipchatAdapter(robot *Robot) (Adapter, error) {
	hipchat := &HipchatAdapter{
		jid:      os.Getenv("HAL_HIPCHAT_JID"),
		password: os.Getenv("HAL_HIPCHAT_PASSWORD"),
		rooms:    os.Getenv("HAL_HIPCHAT_ROOMS"),
	}
	hipchat.SetRobot(robot)
	return hipchat, nil
}

func newIRCAdapter(robot *Robot) (Adapter, error) {
	irc := &IRCAdapter{
		user:     os.Getenv("HAL_IRC_USER"),
		nick:     os.Getenv("HAL_IRC_NICK"),
		password: os.Getenv("HAL_IRC_PASSWORD"),
		server:   os.Getenv("HAL_IRC_SERVER"),
		port:     GetenvDefaultInt("HAL_IRC_PORT", 6667),
		channels: func() []string { return strings.Split(os.Getenv("HAL_IRC_CHANNELS"), ",") }(),
		useTLS:   GetenvDefaultBool("HAL_IRC_USE_TLS", false),
	}
	// Set the robot name to the IRC nick so respond commands will work
	irc.SetRobot(robot)
	irc.Robot.SetName(irc.nick)
	return irc, nil
}
