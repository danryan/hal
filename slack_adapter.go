package hal

import (
	// "github.com/davecgh/go-spew/spew"
	"encoding/json"
	"net/http"
	"net/url"
)

// SlackAdapter struct
type SlackAdapter struct {
	BasicAdapter
	token     string
	team      string
	mode      string
	channels  string //[]string
	botname   string
	iconEmoji string
	linkNames int
}

// Send sends a regular response
func (a *SlackAdapter) Send(res *Response, strings ...string) error {
	a.Logger.Debug("slack - sending response")
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

// TODO: implement
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
	return nil
}

// Run starts the adapter
func (a *SlackAdapter) Run() error {
	a.preRun()
	// set up handlers
	a.Router.HandleFunc("/hal/slack-webhook", a.slackHandler)
	// Someday we won't need this :D
	a.Router.HandleFunc("/hubot/slack-webhook", a.slackHandler)
	a.postRun()

	return nil
}

func (a *SlackAdapter) slackHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	parsedRequest := a.parseRequest(r.Form)
	message := a.newMessage(parsedRequest)

	// a.Logger.Debug(message)
	a.Receive(message)
	w.Write([]byte(""))
}

// Stop shuts down the adapter
func (a *SlackAdapter) Stop() error {
	a.stop()
	return nil
}

func (a *SlackAdapter) Name() string {
	return "slack"
}

// TODO: implement
func (a *SlackAdapter) newMessage(req *slackRequest) *Message {
	return &Message{
		User: &User{
			ID: req.UserName,
		},
		Room: req.ChannelID,
		Text: req.Text,
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

func (a *SlackAdapter) startIRCGateway() error {

	return nil
}
