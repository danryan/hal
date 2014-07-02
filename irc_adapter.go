package hal

import (
	"crypto/tls"
	"fmt"
	irc "github.com/thoj/go-ircevent"
)

// IRCAdapter struct
type IRCAdapter struct {
	BasicAdapter
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

// Send sends a regular response
func (a *IRCAdapter) Send(res *Response, strings ...string) error {
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
func (a *IRCAdapter) Reply(res *Response, strings ...string) error {
	newStrings := make([]string, len(strings))
	for _, str := range strings {
		newStrings = append(newStrings, res.UserID()+`: `+str)
	}

	a.Send(res, newStrings...)

	return nil
}

// Emote is not implemented.
func (a *IRCAdapter) Emote(res *Response, strings ...string) error {
	return nil
}

// Topic sets the topic
func (a *IRCAdapter) Topic(res *Response, strings ...string) error {
	for _, str := range strings {
		a.conn.SendRawf("TOPIC %s %s", res.Room(), str)
	}
	return nil
}

// Play is not implemented.
func (a *IRCAdapter) Play(res *Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *IRCAdapter) Receive(msg *Message) error {
	a.Logger.Debug("irc - adapter received message")
	a.Robot.Receive(msg)
	a.Logger.Debug("irc - adapter sent message to robot")

	return nil
}

// Run starts the adapter
func (a *IRCAdapter) Run() error {
	a.preRun()
	// set up a connection to the IRC gateway
	a.Logger.Debug("irc - starting IRC connection")
	go a.startIRCConnection()
	a.Logger.Debug("irc - started IRC connection")
	a.postRun()

	return nil
}

// Stop shuts down the adapter
func (a *IRCAdapter) Stop() error {
	a.stop()
	a.Logger.Debug("irc - stopping IRC connection")
	a.stopIRCConnection()
	a.Logger.Debug("irc - stopped IRC connection")
	return nil
}

func (a *IRCAdapter) Name() string {
	return "irc"
}

func (a *IRCAdapter) newMessage(req *irc.Event) *Message {
	return &Message{
		User: &User{
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

func (a *IRCAdapter) startIRCConnection() {
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

func (a *IRCAdapter) stopIRCConnection() {
	a.Logger.Debug("Stopping irc IRC connection")
	a.conn.Quit()
	a.Logger.Debug("Stopped irc IRC connection")
}

func (a *IRCAdapter) connectionString() string {
	return fmt.Sprintf("%s:%d", a.server, a.port)
}
