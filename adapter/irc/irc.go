package irc

import (
	"crypto/tls"
	"fmt"
	"github.com/danryan/env"
	"github.com/danryan/hal"
	irc "github.com/thoj/go-ircevent"
	"strings"
)

func init() {
	hal.RegisterAdapter("irc", New)
}

// adapter struct
type adapter struct {
	hal.BasicAdapter
	user     string
	nick     string
	password string
	server   string
	port     int
	mode     string
	channels []string //[]string
	useTLS   bool
	conn     *irc.Connection
}

type config struct {
	User     string `env:"required key=HAL_IRC_USER"`
	Nick     string `env:"required key=HAL_IRC_NICK"`
	Password string `env:"key=HAL_IRC_PASSWORD"`
	Server   string `env:"required key=HAL_IRC_SERVER"`
	Port     int    `env:"key=HAL_IRC_PORT default=6667"`
	Channels string `env:"required key=HAL_IRC_CHANNELS"`
	UseTLS   bool   `env:"key=HAL_IRC_USE_TLS default=false"`
}

// New returns an initialized adapter
func New(robot *hal.Robot) (hal.Adapter, error) {
	c := &config{}
	env.MustProcess(c)

	a := &adapter{
		user:     c.User,
		nick:     c.Nick,
		password: c.Password,
		server:   c.Server,
		port:     c.Port,
		channels: func() []string { return strings.Split(c.Channels, ",") }(),
		useTLS:   c.UseTLS,
	}
	// Set the robot name to the IRC nick so respond commands will work
	a.SetRobot(robot)
	a.Robot.SetName(a.nick)
	return a, nil
}

// Send sends a regular response
func (a *adapter) Send(res *hal.Response, strings ...string) error {
	a.Logger.Debug("irc - sending IRC response")
	for _, str := range strings {
		s := &ircPayload{
			Channel: res.Message.Room,
			Text:    str,
		}
		a.conn.Privmsg(s.Channel, s.Text)
	}
	a.Logger.Debug("irc - sent IRC response")
	return nil
}

// Reply sends a direct response
func (a *adapter) Reply(res *hal.Response, strings ...string) error {
	newStrings := make([]string, len(strings))
	for _, str := range strings {
		newStrings = append(newStrings, res.UserID()+`: `+str)
	}

	a.Send(res, newStrings...)

	return nil
}

// Emote is not implemented.
func (a *adapter) Emote(res *hal.Response, strings ...string) error {
	return nil
}

// Topic sets the topic
func (a *adapter) Topic(res *hal.Response, strings ...string) error {
	for _, str := range strings {
		a.conn.SendRawf("TOPIC %s %s", res.Room(), str)
	}
	return nil
}

// Play is not implemented.
func (a *adapter) Play(res *hal.Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *adapter) Receive(msg *hal.Message) error {
	a.Logger.Debug("irc - adapter received message")
	a.Robot.Receive(msg)
	a.Logger.Debug("irc - adapter sent message to robot")

	return nil
}

// Run starts the adapter
func (a *adapter) Run() error {
	// set up a connection to the IRC gateway
	a.Logger.Debug("irc - starting IRC connection")
	go a.startIRCConnection()
	a.Logger.Debug("irc - started IRC connection")

	return nil
}

// Stop shuts down the adapter
func (a *adapter) Stop() error {
	a.Logger.Debug("irc - stopping IRC connection")
	a.stopIRCConnection()
	a.Logger.Debug("irc - stopped IRC connection")

	return nil
}

func (a *adapter) Name() string {
	return "irc"
}

func (a *adapter) newMessage(req *irc.Event) *hal.Message {
	return &hal.Message{
		User: &hal.User{
			ID: req.Nick,
		},
		Room: req.Arguments[0],
		Text: req.Message(),
	}
}

type ircPayload struct {
	Channel  string
	Username string
	Text     string
}

func (a *adapter) startIRCConnection() {
	if a.nick == "" {
		a.nick = a.user
	}

	conn := irc.IRC(a.nick, a.user)
	if a.useTLS {
		conn.UseTLS = true
		conn.TLSConfig = &tls.Config{ServerName: a.server}
	}
	conn.Password = a.password
	conn.Debug = (a.Logger.Level() == 10)

	err := conn.Connect(a.connectionString())
	if err != nil {
		panic("failed to connect to" + err.Error())
	}

	conn.AddCallback("001", func(e *irc.Event) {
		for _, channel := range a.channels {
			conn.Join(channel)
			a.Logger.Debug("irc - joined " + channel)
		}
	})

	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		message := a.newMessage(e)
		a.Receive(message)
	})

	a.conn = conn
	a.Logger.Debug("irc - waiting for server acknowledgement")
	conn.Loop()
}

func (a *adapter) stopIRCConnection() {
	a.Logger.Debug("Stopping irc IRC connection")
	a.conn.Quit()
	a.Logger.Debug("Stopped irc IRC connection")
}

func (a *adapter) connectionString() string {
	return fmt.Sprintf("%s:%d", a.server, a.port)
}
