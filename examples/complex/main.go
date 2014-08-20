package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	"github.com/danryan/hal/handler"
	_ "github.com/danryan/hal/store/memory"
	"os"
)

// HAL is just another Go package, which means you are free to organize things
// however you deem best.

// You can define your handlers in the same file...
var pingHandler = hal.Hear(`ping`, func(res *hal.Response) error {
	return res.Send("PONG")
})

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	// Or define them inside another function...
	fooHandler := hal.Respond(`foo`, func(res *hal.Response) error {
		return res.Send("BAR")
	})

	tableFlipHandler := &hal.Handler{
		Method:  hal.HEAR,
		Pattern: `tableflip`,
		Run: func(res *hal.Response) error {
			return res.Send(`(╯°□°）╯︵ ┻━┻`)
		},
	}

	robot.Handle(
		pingHandler,
		fooHandler,
		tableFlipHandler,

		// Or stick them in an entirely different package, and reference them
		// exactly in the way you would expect.
		handler.Ping,

		// Or use a hal.Handler structure complete with usage...
		&hal.Handler{
			Method:  hal.RESPOND,
			Pattern: `SYN`,
			Usage:   `hal syn - replies with "ACK"`,
			Run: func(res *hal.Response) error {
				return res.Reply("ACK")
			},
		},

		// Or even inline!
		hal.Hear(`yo`, func(res *hal.Response) error {
			return res.Send("lo")
		}),
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
