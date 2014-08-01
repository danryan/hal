package hipchat

import (
	"fmt"
	"github.com/daneharrigan/hipchat"
	"github.com/danryan/env"
	"github.com/danryan/hal"
	"strings"
)

func init() {
	hal.RegisterAdapter("hipchat", New)
}

// adapter struct
type adapter struct {
	hal.BasicAdapter
	user     string
	nick     string
	name     string
	password string
	resource string
	rooms    []string
	client   *hipchat.Client
	// config   *config
}

type config struct {
	User     string `env:"required key=HAL_HIPCHAT_USER"`
	Password string `env:"required key=HAL_HIPCHAT_PASSWORD"`
	Rooms    string `env:"required key=HAL_HIPCHAT_ROOMS"`
	Resource string `env:"key=HAL_HIPCHAT_RESOURCE default=bot"`
}

// New returns an initialized adapter
func New(robot *hal.Robot) (hal.Adapter, error) {
	c := &config{}
	env.MustProcess(c)

	a := &adapter{
		user:     c.User,
		password: c.Password,
		resource: c.Resource,
		rooms:    func() []string { return strings.Split(c.Rooms, ",") }(),
	}
	a.SetRobot(robot)
	return a, nil
}

// Run starts the adapter
func (a *adapter) Run() error {
	go a.startConnection()
	return nil
}

// Stop shuts down the adapter
func (a *adapter) Stop() error {
	// hipchat package doesn't provide an explicit stop command
	return nil
}

// Send sends a regular response
func (a *adapter) Send(res *hal.Response, strings ...string) error {
	for _, str := range strings {
		a.client.Say(res.Message.Room, a.name, str)
	}
	return nil
}

// Reply sends a direct response
func (a *adapter) Reply(res *hal.Response, strings ...string) error {
	newStrings := make([]string, len(strings))
	for _, str := range strings {
		s := fmt.Sprintf("@%s: %s", mentionName(res.Envelope.User), str)
		newStrings = append(newStrings, s)
	}

	return a.Send(res, newStrings...)
}

// Emote is not implemented.
func (a *adapter) Emote(res *hal.Response, strings ...string) error {
	return nil
}

// Topic is not implemented.
func (a *adapter) Topic(res *hal.Response, strings ...string) error {
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

func (a *adapter) newMessage(msg *hipchat.Message) *hal.Message {
	from := strings.Split(msg.From, "/")
	user, _ := a.Robot.Users.GetByName(from[1])

	return &hal.Message{
		User: user,
		Room: from[0],
		Text: msg.Body,
	}
}

func mentionName(u *hal.User) string {
	mn, ok := u.Options["mentionName"]
	if !ok {
		return ""
	}
	return mn.(string)
}

func (a *adapter) startConnection() error {
	client, err := hipchat.NewClient(a.user, a.password, a.resource)
	if err != nil {
		hal.Logger.Error(err.Error())
		return err
	}

	client.Status("chat")

	for _, user := range client.Users() {
		// retrieve the name and mention name of our bot from the server
		if user.Id == client.Id {
			a.name = user.Name
			a.nick = user.MentionName
			// skip adding the bot to the users map
			continue
		}
		// Initialize a newUser object in case we need it.
		newUser := hal.User{
			ID:   user.Id,
			Name: user.Name,
			Options: map[string]interface{}{
				"mentionName": user.MentionName,
			},
		}
		// Prepopulate our users map because we can easily do so.
		// If a user doesn't exist, set it.
		u, err := a.Robot.Users.Get(user.Id)
		if err != nil {
			a.Robot.Users.Set(user.Id, newUser)
		}
		// If the user doesn't match completely (say, if someone changes their name),
		// then adjust what we have stored.
		if u.Name != user.Name || mentionName(&u) != user.MentionName {
			a.Robot.Users.Set(user.Id, newUser)
		}
	}

	// Make a map of room JIDs to human names
	roomJids := make(map[string]string, len(client.Rooms()))
	for _, room := range client.Rooms() {
		roomJids[room.Name] = room.Id
	}
	client.Status("chat")
	// Only join the rooms we want
	for _, room := range a.rooms {
		hal.Logger.Debugf("%s - joined %s", a, room)
		client.Join(roomJids[room], a.name)
	}

	a.client = client
	a.Robot.Alias = a.nick

	// send an empty string every 60 seconds so hipchat doesn't disconnect us
	go client.KeepAlive()

	for message := range client.Messages() {
		from := strings.Split(message.From, "/")
		// ignore messages directly from the channel
		// TODO: don't do this :)
		if len(from) < 2 {
			continue
		}
		// ingore messages from our bot
		if from[1] == a.name {
			continue
		}

		msg := a.newMessage(message)
		a.Receive(msg)
	}
	return nil
}
