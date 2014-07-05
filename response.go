package hal

import (
	"github.com/ccding/go-logging/logging"
	// "regexp"
)

// Response struct
type Response struct {
	Robot    *Robot
	Message  *Message
	Match    [][]string
	Listener Listener
	Logger   *logging.Logger
	// Match    []string
	// Envelope *Envelope
}

// Envelope contains metadata about the chat message.
type Envelope struct {
	Room    string
	User    *User
	Message *Message
}

// User is the user of the response's message
func (res *Response) UserID() string {
	return res.Message.User.ID
}

// Room is the room of the response's message
func (res *Response) Room() string {
	return res.Message.Room
}

// Text is the text of the response's message
func (res *Response) Text() string {
	return res.Message.Text
}

// NewResponse returns a new Response object
func NewResponse(robot *Robot, msg *Message) *Response {
	return &Response{
		Message: msg,
		Robot:   robot,
		Logger:  robot.Logger,
	}
}

// Send posts a message back to the chat source
func (response *Response) Send(strings ...string) error {
	if err := response.Robot.Adapter.Send(response, strings...); err != nil {
		response.Logger.Error(err)
		return err
	}
	return nil


// Reply posts a message mentioning the current user
func (response *Response) Reply(strings ...string) error {
	if err := response.Robot.Adapter.Reply(response, strings...); err != nil {
		response.Logger.Error(err)
		return err
	}
	return nil
}

// Emote posts an emote back to the chat source
func (response *Response) Emote(strings ...string) error {
	if err := response.Robot.Adapter.Emote(response, strings...); err != nil {
		response.Logger.Error(err)
		return err
	}
	return nil
}

// Topic posts a topic changing message
func (response *Response) Topic(strings ...string) error {
	if err := response.Robot.Adapter.Topic(response, strings...); err != nil {
		response.Logger.Error(err)
		return err
	}
	return nil
}

// Play posts a sound message
func (response *Response) Play(strings ...string) error {
	if err := response.Robot.Adapter.Play(response, strings...); err != nil {
		response.Logger.Error(err)
		return err
	}
	return nil
}
