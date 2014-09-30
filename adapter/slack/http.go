package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danryan/hal"
)

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

func (a *adapter) slackHandler(w http.ResponseWriter, r *http.Request) {
	hal.Logger.Debug("slack - HTTP handler received message")

	r.ParseForm()
	parsedRequest := a.parseRequest(r.Form)
	message := a.newMessageFromHTTP(parsedRequest)

	// hal.Logger.Debug(message)
	a.Receive(message)
	w.Write([]byte(""))
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
