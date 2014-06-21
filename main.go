package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gobot/gobot"
	"os"
)

func main() {
	os.Exit(Run())
}

func Run() int {
	robot := gobot.NewRobot()

	robot.Handle(
		pingCmd,
		synCmd,
	)

	spew.Dump(robot.Handlers())

	err := robot.Run()
	if err != nil {
		fmt.Println("Failure!")
		return 1
	}
	fmt.Println("Success!")
	return 0
}
