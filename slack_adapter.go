package hal

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	irc "github.com/thoj/go-ircevent"
	"net/http"
	"net/url"
)

// SlackAdapter struct
type SlackAdapter struct {
	BasicAdapter
	token          string
	team           string
	mode           string
	channels       string //[]string
	botname        string
	responseMethod string
	iconEmoji      string
	ircEnabled     bool
	ircPassword    string
	ircConnection  *irc.Connection
	linkNames      int
}

// Send sends a regular response
func (a *SlackAdapter) Send(res *Response, strings ...string) error {
	var err error

	if a.responseMethod == "irc" {
		if !a.ircEnabled {
			return errors.New("slack - IRC response method used but IRC is not enabled")
		}
		a.sendIRC(res, strings...)

	} else {
		err = a.sendHTTP(res, strings...)
		if err != nil {
			return err
		}
	}

	return nil
}

// Reply sends a direct response
func (a *SlackAdapter) Reply(res *Response, strings ...string) error {
	newStrings := make([]string, len(strings))
	for _, str := range strings {
		newStrings = append(newStrings, res.UserID()+`: `+str)
	}

	a.Send(res, newStrings...)

	return nil
}

// Emote is not implemented.
func (a *SlackAdapter) Emote(res *Response, strings ...string) error {
	return nil
}

// Topic sets the topic
func (a *SlackAdapter) Topic(res *Response, strings ...string) error {
	for _, str := range strings {
		_ = str
	}
	return nil
}

// Play is not implemented.
func (a *SlackAdapter) Play(res *Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *SlackAdapter) Receive(msg *Message) error {
	a.Logger.Debug("slack - adapter received message")
	a.Robot.Receive(msg)
	a.Logger.Debug("slack - adapter sent message to robot")

	return nil
}

// Run starts the adapter
func (a *SlackAdapter) Run() error {
	a.preRun()

	if a.ircEnabled {
		// set up a connection to the IRC gateway
		a.Logger.Debug("slack - starting IRC connection")
		go a.startIRCConnection()
		a.Logger.Debug("slack - started IRC connection")
	} else {
		// set up handlers
		a.Logger.Debug("slack - adding HTTP request handlers")
		a.Router.HandleFunc("/hal/slack-webhook", a.slackHandler)
		// Someday we won't need this :D
		a.Router.HandleFunc("/hubot/slack-webhook", a.slackHandler)
		a.Logger.Debug("slack - added HTTP request handlers")
	}
	a.postRun()

	return nil
}

func (a *SlackAdapter) slackHandler(w http.ResponseWriter, r *http.Request) {
	a.Logger.Debug("slack - HTTP handler received message")

	r.ParseForm()
	parsedRequest := a.parseRequest(r.Form)
	message := a.newMessageFromHTTP(parsedRequest)

	// a.Logger.Debug(message)
	a.Receive(message)
	w.Write([]byte(""))
}

// Stop shuts down the adapter
func (a *SlackAdapter) Stop() error {
	a.stop()
	if a.ircEnabled {
		// set up a connection to the IRC gateway
		a.Logger.Debug("slack - stopping IRC connection")
		a.stopIRCConnection()
		a.Logger.Debug("slack - stopped IRC connection")
	}
	return nil
}

func (a *SlackAdapter) Name() string {
	return "slack"
}

// TODO: implement
func (a *SlackAdapter) newMessageFromHTTP(req *slackRequest) *Message {
	return &Message{
		User: &User{
			ID: req.UserName,
		},
		Room: req.ChannelID,
		Text: req.Text,
	}
}

func (a *SlackAdapter) newMessageFromIRC(req *irc.Event) *Message {
	return &Message{
		User: &User{
			ID: req.Nick,
		},
		Room: req.Arguments[0],
		Text: req.Message(),
	}
}

func (a *SlackAdapter) parseRequest(form url.Values) *slackRequest {
	return &slackRequest{
		ChannelID:   form.Get("channel_id"),
		ChannelName: form.Get("channel_name"),
		ServiceID:   form.Get("service_id"),
		TeamID:      form.Get("team_id"),
		TeamDomain:  form.Get("team_domain"),
		Text:        form.Get("text"),
		Timestamp:   form.Get("timestamp"),
		Token:       form.Get("token"),
		UserID:      form.Get("user_id"),
		UserName:    form.Get("user_name"),
	}
}

type slackPayload struct {
	Channel   string `json:"channel,omitempty"`
	Username  string `json:"username,omitempty"`
	Text      string `json:"text,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

// the payload of an inbound request (from Slack to us).
type slackRequest struct {
	ChannelID   string
	ChannelName string
	ServiceID   string
	TeamID      string
	TeamDomain  string
	Text        string
	Timestamp   string
	Token       string
	UserID      string
	UserName    string
}

func (a *SlackAdapter) startIRCConnection() {
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
			a.Logger.Info(channel)
		}
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		a.Logger.Debug("slack - IRC handler received message")

		message := a.newMessageFromIRC(e)
		a.Receive(message)
	})

	a.ircConnection = con
	con.Loop()
}

func (a *SlackAdapter) stopIRCConnection() {
	a.Logger.Debug("Stopping slack IRC connection")
	a.ircConnection.Quit()
	a.Logger.Debug("Stopped slack IRC connection")
}

func (a *SlackAdapter) ircServer() string {
	return a.team + `.irc.slack.com:6667`
}

func (a *SlackAdapter) sendHTTP(res *Response, strings ...string) error {
	a.Logger.Debug("slack - sending HTTP response")
	for _, str := range strings {
		s := &slackPayload{
			Username: a.botname,
			Channel:  res.Message.Room,
			Text:     str,
		}

		u := `https://` + a.team + `.slack.com/services/hooks/hubot?token=` + a.token
		payload, _ := json.Marshal(s)
		data := url.Values{}
		data.Set("payload", string(payload))

		client := http.Client{}
		_, err := client.PostForm(u, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *SlackAdapter) sendIRC(res *Response, strings ...string) error {
	a.Logger.Debug("slack - sending IRC response")
	for _, str := range strings {
		s := &slackPayload{
			Channel: res.Message.Room,
			Text:    str,
		}
		a.ircConnection.Privmsg(s.Channel, s.Text)
	}

	return nil
}
