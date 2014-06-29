package hal

import (
	_ "crypto/tls"
	_ "encoding/json"
	_ "errors"
	_ "github.com/mattn/go-xmpp"
	"net/http"
	"net/url"
)

// HipchatAdapter struct
type HipchatAdapter struct {
	BasicAdapter
	jid      string
	password string
	rooms    string
}

// Send sends a regular response
func (a *HipchatAdapter) Send(res *Response, strings ...string) error {
	return nil
}

// Reply sends a direct response
func (a *HipchatAdapter) Reply(res *Response, strings ...string) error {
	newStrings := make([]string, len(strings))
	for _, str := range strings {
		newStrings = append(newStrings, res.UserID()+`: `+str)
	}

	a.Send(res, newStrings...)

	return nil
}

// Emote is not implemented.
func (a *HipchatAdapter) Emote(res *Response, strings ...string) error {
	return nil
}

// Topic sets the topic
func (a *HipchatAdapter) Topic(res *Response, strings ...string) error {
	for _, str := range strings {
		_ = str
	}
	return nil
}

// Play is not implemented.
func (a *HipchatAdapter) Play(res *Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *HipchatAdapter) Receive(msg *Message) error {
	a.Logger.Debug("hipchat - adapter received message")
	a.Robot.Receive(msg)
	a.Logger.Debug("hipchat - adapter sent message to robot")

	return nil
}

// Run starts the adapter
func (a *HipchatAdapter) Run() error {
	a.preRun()

	a.Logger.Debug("hipchat - adding HTTP request handlers")
	a.Router.HandleFunc("/hal/hipchat-webhook", a.hipchatHandler)
	a.postRun()

	return nil
}

func (a *HipchatAdapter) hipchatHandler(w http.ResponseWriter, r *http.Request) {
	a.Logger.Debug("hipchat - HTTP handler received message")

	r.ParseForm()
	parsedRequest := a.parseRequest(r.Form)
	message := a.newMessageFromHTTP(parsedRequest)

	// a.Logger.Debug(message)
	a.Receive(message)
	w.Write([]byte(""))
}

// Stop shuts down the adapter
func (a *HipchatAdapter) Stop() error {
	a.stop()
	return nil
}

func (a *HipchatAdapter) Name() string {
	return "hipchat"
}

// TODO: implement
func (a *HipchatAdapter) newMessageFromHTTP(req *hipchatRequest) *Message {
	return &Message{
		User: &User{
			ID: req.UserName,
		},
		Room: req.ChannelID,
		Text: req.Text,
	}
}

func (a *HipchatAdapter) parseRequest(form url.Values) *hipchatRequest {
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

func (a *HipchatAdapter) sendHTTP(res *Response, strings ...string) error {
	a.Logger.Debug("hipchat - sending HTTP response")
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

func (a *HipchatAdapter) newClient() *HipchatClient {
	return &HipchatClient{}
}
