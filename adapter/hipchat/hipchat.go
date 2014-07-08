package hipchat

import (
	_ "crypto/tls"
	_ "encoding/json"
	_ "errors"
	"github.com/danryan/hal"
	_ "github.com/mattn/go-xmpp"
	"net/http"
	"net/url"
)

// adapter struct
type adapter struct {
	hal.BasicAdapter
	jid      string
	password string
	rooms    string
}

// Send sends a regular response
func (a *adapter) Send(res *hal.Response, strings ...string) error {
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
	hal.Logger.Debug("hipchat - adapter received message")
	a.Robot.Receive(msg)
	hal.Logger.Debug("hipchat - adapter sent message to robot")

	return nil
}

// Run starts the adapter
func (a *adapter) Run() error {
	hal.Logger.Debug("hipchat - adding HTTP request handlers")
	hal.Router.HandleFunc("/hal/hipchat-webhook", a.hipchatHandler)

	return nil
}

func (a *adapter) hipchatHandler(w http.ResponseWriter, r *http.Request) {
	hal.Logger.Debug("hipchat - HTTP handler received message")

	r.ParseForm()
	parsedRequest := a.parseRequest(r.Form)
	message := a.newMessageFromHTTP(parsedRequest)

	a.Receive(message)
	w.Write([]byte(""))
}

// Stop shuts down the adapter
func (a *adapter) Stop() error {
	return nil
}

func (a *adapter) Name() string {
	return "hipchat"
}

func (a *adapter) newMessageFromHTTP(req *hipchatRequest) *hal.Message {
	return &hal.Message{
		User: &hal.User{
			ID: req.UserName,
		},
		Room: req.ChannelID,
		Text: req.Text,
	}
}

func (a *adapter) parseRequest(form url.Values) *hipchatRequest {
	return &hipchatRequest{
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

type hipchatPayload struct {
	Channel   string `json:"channel,omitempty"`
	Username  string `json:"username,omitempty"`
	Text      string `json:"text,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

// the payload of an inbound request (from Hipchat to us).
type hipchatRequest struct {
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

func (a *adapter) sendHTTP(res *hal.Response, strings ...string) error {
	hal.Logger.Debug("hipchat - sending HTTP response")
	// for _, str := range strings {
	// s := &hipchatPayload{
	// Username: a.botname,
	// Channel:  res.Message.Room,
	// Text:     str,
	// }

	// u := `https://` + a.team + `.hipchat.com/services/hooks/hubot?token=` + a.token
	// payload, _ := json.Marshal(s)
	// data := url.Values{}
	// data.Set("payload", string(payload))

	// client := http.Client{}
	// _, err := client.PostForm(u, data)
	// if err != nil {
	// return err
	// }
	// }

	return nil
}

type HipchatClient struct {
}

func (a *adapter) newClient() *HipchatClient {
	return &HipchatClient{}
}
