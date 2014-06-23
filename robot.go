package hal

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Robot receives messages from an adapter and sends them to listeners
type Robot struct {
	*Config
	Name       string
	Alias      string
	Adapter    Adapter
	Logger     *log.Logger
	handlers   []Handler
	signalChan chan os.Signal
	router     *http.ServeMux

	// Listeners map[string][]Handler
}

// Handlers returns the robot's handlers
func (robot *Robot) Handlers() []Handler {
	return robot.handlers
}

// NewRobot returns a new Robot instance
func NewRobot() (*Robot, error) {
	robot := &Robot{}
	config := NewConfig()

	adapter, err := NewAdapter(config.AdapterName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	adapter.SetRobot(robot)

	robot.Name = config.Name
	robot.Logger = config.Logger
	robot.Adapter = adapter
	robot.signalChan = make(chan os.Signal, 1)
	robot.router = http.NewServeMux()

	return robot, nil
}

// Handle registers a new handler with the robot
func (robot *Robot) Handle(handlers ...Handler) {
	robot.handlers = append(robot.handlers, handlers...)
}

// Receive dispatches messages to our handlers
func (robot *Robot) Receive(msg *Message) error {
	for _, handler := range robot.handlers {
		response := NewResponse(robot, msg)

		err := handler.Handle(response)
		if err != nil {
			return err
		}
	}
	return nil
}

// Stop initiates the shutdown process
func (robot *Robot) Stop() error {
	robot.Adapter.Stop()
	log.Println("Shutting down...")

	return nil
}

// Run starts up the robot
func (robot *Robot) Run() error {
	defer robot.Stop()

	stop := false

	go robot.Adapter.Run()

	signal.Notify(robot.signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for !stop {
		select {
		case sig := <-robot.signalChan:
			switch sig {
			case syscall.SIGHUP:
				log.Println("Reloading...")
			case syscall.SIGINT, syscall.SIGTERM:
				stop = true
			}
		}
	}

	return nil
}

func (robot *Robot) RespondRegex(pattern string) string {
	str := `^(?:`
	if robot.Alias != "" {
		str += `(?:` + robot.Alias + `|` + robot.Name + `)`
	} else {
		str += robot.Name
	}
	str += `[:,]?)\s+(?:` + pattern + `)`
	return str
	// return `^(?:(?:h|hubot)[:,]?)\s+(?:` + pattern + `)`
}
