package gobot

type Response struct {
	Robot    *Robot
	Message  *Message
	Match    []string
	Envelope *Envelope
}

// Envelope contains metadata about the chat message.
type Envelope struct {
	Room    string
	User    *User
	Message *Message
}

// Send posts a message back to the chat source
func (r *Response) Send(strings ...string) error {
	r.Robot.Adapter.Send(r.Envelope, strings)
	return nil
}

// Emote posts an emote back to the chat source
func (r *Response) Emote(strings ...string) error {
	r.Robot.Adapter.Emote(r.Envelope, strings)
	return nil
}

// Reply posts a message mentioning the current user
func (r *Response) Reply(strings ...string) error {
	r.Robot.Adapter.Reply(r.Envelope, strings)
	return nil
}

// Topic posts a topic changing message
func (r *Response) Topic(strings ...string) error {
	r.Robot.Adapter.Send(r.Envelope, strings)
	return nil
}

// func (r *Response) Play(strings ...string) error {
// 	r.Robot.Adapter.Play(r.Envelope, strings)
// 	return nil
// }
