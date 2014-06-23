package main

import (
	_ "github.com/davecgh/go-spew/spew"
	"hal/hal"
	"log"
	"os"
)

func main() {
	os.Exit(Run())
}

func Run() int {
	robot, err := hal.NewRobot()

	if err != nil {
		log.Println("Failure!")
		return 1
	}

	robot.Handle(
		respondHandler,
		hearHandler,
	)

	// spew.Dump(robot.Adapter)
	// spew.Dump(robot.Handlers())
	robot.Alias = "hubot"
	log.Println(robot.RespondRegex("do (.+)"))

	if err := robot.Run(); err != nil {
		log.Println("Failure!")
		return 1
	}
	log.Println("Success!")
	return 0
}

// ^(?:(?:h|hubot)[:,]?)\s+(?:do (.+)))
// ^(?:(?:h|hubot)[:,]?)\s+(?:do (.+))
