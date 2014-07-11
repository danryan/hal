package hal

// Response struct
type Response struct {
	Robot   *Robot
	Message *Message
	Match   []string
}

// Envelope contains metadata about the chat message.
type Envelope struct {
	Room    string
	User    *User
	Message *Message
}

// NewResponse returns a new Response object
func NewResponse(robot *Robot, msg *Message) *Response {
	return &Response{
		Message: msg,
		Robot:   robot,
	}
}

// UserID returns the id of the Message's User
func (res *Response) UserID() string {
	return res.Message.User.ID
}

// UserName returns the id of the Message's User
func (res *Response) UserName() string {
	return res.Message.User.Name
}

// UserRoles returns the roles of the Message's User
func (res *Response) UserRoles() []string {
	return res.Message.User.Roles
}

// Room is the room of the response's message
func (res *Response) Room() string {
	return res.Message.Room
}

// Text is the text of the response's message
func (res *Response) Text() string {
	return res.Message.Text
}

// Send posts a message back to the chat source
func (res *Response) Send(strings ...string) error {
	if err := res.Robot.Adapter.Send(res, strings...); err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}

// Reply posts a message mentioning the current user
func (res *Response) Reply(strings ...string) error {
	if err := res.Robot.Adapter.Reply(res, strings...); err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}

// Emote posts an emote back to the chat source
func (res *Response) Emote(strings ...string) error {
	if err := res.Robot.Adapter.Emote(res, strings...); err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}

// Topic posts a topic changing message
func (res *Response) Topic(strings ...string) error {
	if err := res.Robot.Adapter.Topic(res, strings...); err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}

// Play posts a sound message
func (res *Response) Play(strings ...string) error {
	if err := res.Robot.Adapter.Play(res, strings...); err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}
