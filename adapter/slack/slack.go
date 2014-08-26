package slack

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danryan/env"
	"github.com/danryan/hal"
	irc "github.com/thoj/go-ircevent"
)

func init() {
	hal.RegisterAdapter("slack", New)
}

type adapter struct {
	hal.BasicAdapter
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

type config struct {
	Token          string `env:"key=HAL_SLACK_TOKEN required"`
	Team           string `env:"key=HAL_SLACK_TEAM required"`
	Channels       string `env:"key=HAL_SLACK_CHANNELS"`
	Mode           string `env:"key=HAL_SLACK_MODE"`
	Botname        string `env:"key=HAL_SLACK_BOTNAME default=hal"`
	IconEmoji      string `env:"key=HAL_SLACK_ICON_EMOJI"`
	IrcEnabled     bool   `env:"key=HAL_SLACK_IRC_ENABLED default=false"`
	IrcPassword    string `env:"key=HAL_SLACK_IRC_PASSWORD"`
	ResponseMethod string `env:"key=HAL_SLACK_RESPONSE_METHOD default=http"`
}

// New returns an initialized adapter
func New(r *hal.Robot) (hal.Adapter, error) {
	c := &config{}
	env.MustProcess(c)
	a := &adapter{
		token:          c.Token,
		team:           c.Team,
		channels:       c.Channels,
		mode:           c.Mode,
		botname:        c.Botname,
		iconEmoji:      c.IconEmoji,
		ircEnabled:     c.IrcEnabled,
		ircPassword:    c.IrcPassword,
		responseMethod: c.ResponseMethod,
	}
	a.SetRobot(r)
	return a, nil
}

// Send sends a regular response
func (a *adapter) Send(res *hal.Response, strings ...string) error {
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
func (a *adapter) Reply(res *hal.Response, strings ...string) error {
	newStrings := make([]string, len(strings))
	for _, str := range strings {
		newStrings = append(newStrings, fmt.Sprintf("%s: %s", res.UserName(), str))
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
		_ = str
	}
	return nil
}

// Play is not implemented.
func (a *adapter) Play(res *hal.Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *adapter) Receive(msg *hal.Message) error {
	hal.Logger.Debug("slack - adapter received message")
	a.Robot.Receive(msg)
	hal.Logger.Debug("slack - adapter sent message to robot")

	return nil
}

// Run starts the adapter
func (a *adapter) Run() error {
	if a.ircEnabled {
		// set up a connection to the IRC gateway
		hal.Logger.Debug("slack - starting IRC connection")
		go a.startIRCConnection()
		hal.Logger.Debug("slack - started IRC connection")
	} else {
		// set up handlers
		hal.Logger.Debug("slack - adding HTTP request handlers")
		hal.Router.HandleFunc("/hal/slack-webhook", a.slackHandler)
		// Someday we won't need this :D
		hal.Router.HandleFunc("/hubot/slack-webhook", a.slackHandler)
		hal.Logger.Debug("slack - added HTTP request handlers")
	}

	return nil
}

func (a *adapter) slackHandler(w http.ResponseWriter, r *http.Request) {
	hal.Logger.Debug("slack - HTTP handler received message")

	r.ParseForm()
	parsedRequest := a.parseRequest(r.Form)
	message := a.newMessageFromHTTP(parsedRequest)

	// hal.Logger.Debug(message)
	a.Receive(message)
	w.Write([]byte(""))
}

// Stop shuts down the adapter
func (a *adapter) Stop() error {
	if a.ircEnabled {
		// set up a connection to the IRC gateway
		hal.Logger.Debug("slack - stopping IRC connection")
		a.stopIRCConnection()
		hal.Logger.Debug("slack - stopped IRC connection")
	}
	return nil
}

func (a *adapter) newMessageFromHTTP(req *slackRequest) *hal.Message {
	return &hal.Message{
		User: hal.User{
			ID:   req.UserID,
			Name: req.UserName,
		},
		Room: req.ChannelID,
		Text: req.Text,
	}
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

func (a *adapter) parseRequest(form url.Values) *slackRequest {
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
	Channel     string                   `json:"channel,omitempty"`
	Username    string                   `json:"username,omitempty"`
	Text        string                   `json:"text,omitempty"`
	IconEmoji   string                   `json:"icon_emoji,omitempty"`
	IconURL     string                   `json:"icon_url,omitempty"`
	UnfurlLinks bool                     `json:"unfurl_links,omitempty"`
	Fallback    string                   `json:"fallback,omitempty"`
	Color       string                   `json:"color,omitempty"`
	Fields      []map[string]interface{} `json:"fields,omitempty"`
}

type slackField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
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
			hal.Logger.Info(channel)
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

func (a *adapter) ircServer() string {
	return a.team + `.irc.slack.com:6667`
}

func (a *adapter) sendHTTP(res *hal.Response, strings ...string) error {
	hal.Logger.Debug("slack - sending HTTP response")
	for _, str := range strings {
		s := &slackPayload{
			Username: a.botname,
			Channel:  res.Message.Room,
			Text:     str,
		}

		opts := res.Envelope.Options
		if i, ok := opts["iconEmoji"]; ok {
			s.IconEmoji = i.(string)
		}

		if i, ok := opts["iconURL"]; ok {
			s.IconURL = i.(string)
		}

		if i, ok := opts["unfurlLinks"]; ok {
			s.UnfurlLinks = i.(bool)
		}

		if i, ok := opts["fallback"]; ok {
			s.Fallback = i.(string)
		}

		if i, ok := opts["color"]; ok {
			s.Color = i.(string)
		}

		if i, ok := opts["fields"]; ok {
			s.Fields = i.([]map[string]interface{})
		}

		u := fmt.Sprintf("https://%s.slack.com/services/hooks/incoming-webhook?token=%s", a.team, a.token)
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
