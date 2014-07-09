package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/irc"
	_ "github.com/danryan/hal/adapter/shell"
	_ "github.com/danryan/hal/adapter/slack"
	"log"
	"os"
)

var pingHandler = hal.Hear(`ping`, func(res *hal.Response) error {
	return res.Send("PONG")
})

var openDoorsHandler = hal.Respond(`open the pod bay doors`, func(res *hal.Response) error {
	return res.Reply("I'm sorry, Dave. I can't do that.")
})

func main() {
	os.Exit(Run())
}

// Run the robot
func Run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		log.Println(err)
		return 1
	}

	robot.Handle(
		pingHandler,
		openDoorsHandler,
	)

	if err := robot.Run(); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}
