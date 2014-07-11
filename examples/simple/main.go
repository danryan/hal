package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	_ "github.com/danryan/hal/store/memory"
	"os"
)

var pingHandler = hal.Hear(`ping`, func(res *hal.Response) error {
	return res.Send("PONG")
})

var echoHandler = hal.Respond(`echo (.+)`, func(res *hal.Response) error {
	return res.Reply(res.Match[1])
})

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		pingHandler,
		echoHandler,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
