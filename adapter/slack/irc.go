package slack

import (
	"crypto/tls"

	"github.com/danryan/hal"
	"github.com/thoj/go-ircevent"
)

func (a *adapter) startIRCConnection() {
	con := irc.IRC(a.botname, a.botname)
	con.UseTLS = true
	// con.Debug = true
	con.Password = a.ircPassword
	con.TLSConfig = &tls.Config{ServerName: "*.irc.slack.com"}
	err := con.Connect(a.ircServer())
	if err != nil {
		panic("failed to connect to" + err.Error())
	}

	con.AddCallback("001", func(e *irc.Event) {
		for _, channel := range a.channels {
			hal.Logger.Debugf("slack - joined channel %v", channel)
		}
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		hal.Logger.Debug("slack - IRC handler received message")

		message := a.newMessageFromIRC(e)
		a.Receive(message)
	})

	a.ircConnection = con
	con.Loop()
}

func (a *adapter) stopIRCConnection() {
	hal.Logger.Debug("Stopping slack IRC connection")
	a.ircConnection.Quit()
	hal.Logger.Debug("Stopped slack IRC connection")
}

func (a *adapter) newMessageFromIRC(req *irc.Event) *hal.Message {
	return &hal.Message{
		User: hal.User{
			ID:   req.Nick,
			Name: req.Nick,
		},
		Room: req.Arguments[0],
		Text: req.Message(),
	}
}

func (a *adapter) ircServer() string {
	return a.team + `.irc.slack.com:6667`
}

func (a *adapter) sendIRC(res *hal.Response, strings ...string) error {
	hal.Logger.Debug("slack - sending IRC response")
	for _, str := range strings {
		s := &slackPayload{
			Channel: res.Message.Room,
			Text:    str,
		}
		a.ircConnection.Privmsg(s.Channel, s.Text)
	}

	return nil
}
