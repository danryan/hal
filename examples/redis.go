package main

import (
	"fmt"
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/irc"
	_ "github.com/danryan/hal/adapter/shell"
	_ "github.com/danryan/hal/adapter/slack"
	_ "github.com/danryan/hal/adapter/test"
	_ "github.com/danryan/hal/store/redis"
	"github.com/davecgh/go-spew/spew"
	"os"
)

var pingHandler = hal.Hear(`ping`, func(res *hal.Response) error {
	return res.Send("PONG")
})

var getHandler = hal.Hear(`get (.+)`, func(res *hal.Response) error {
	key := res.Match[1]
	val, err := res.Robot.Store.Get(key)
	if err != nil {
		res.Send(err.Error())
		return err
	}
	return res.Send(string(val))
})

var setHandler = hal.Hear(`set (.+) (.+)`, func(res *hal.Response) error {
	key := res.Match[1]
	val := res.Match[2]
	err := res.Robot.Store.Set(key, []byte(val))
	if err != nil {
		res.Send(err.Error())
		return err
	}
	return res.Send("OK")
})

var deleteHandler = hal.Hear(`delete (.+)`, func(res *hal.Response) error {
	key := res.Match[1]
	if err := res.Robot.Store.Delete(key); err != nil {
		res.Send(err.Error())
		return err
	}
	return res.Send("OK")
})

var usersHandler = hal.Hear(`show users`, func(res *hal.Response) error {
	// users, _ := res.Robot.Store.Get("hal:users")
	lines := []string{}
	for _, user := range res.Robot.Users.All() {
		lines = append(lines, spew.Sdump(user))
	}
	return res.Send(lines...)
})

var userHandler = hal.Hear(`user (.+)`, func(res *hal.Response) error {
	id := res.Match[1]
	user, _ := res.Robot.Users.Get(id)
	// users, _ := res.Robot.Store.Get("hal:users")
	line := spew.Sdump(user)
	return res.Send(line)
})

func main() {
	os.Exit(Run())
}

// Run returns an int so we can return a proper exit code
func Run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	robot.Handle(
		getHandler,
		setHandler,
		deleteHandler,
		usersHandler,
		userHandler,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}
