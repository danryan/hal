package hal

import (
	"fmt"
	"github.com/ccding/go-logging/logging"
	// "github.com/davecgh/go-spew/spew"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Robot receives messages from an adapter and sends them to listeners
type Robot struct {
	*Config
	Name    string
	Alias   string
	Adapter Adapter
	Logger  *logging.Logger
	Port    string
	Router  *http.ServeMux

	handlers   []Handler
	signalChan chan os.Signal
}

// Handlers returns the robot's handlers
func (robot *Robot) Handlers() []Handler {
	return robot.handlers
}

// NewRobot returns a new Robot instance
func NewRobot() (*Robot, error) {
	config := NewConfig()
	robot := &Robot{
		Name:       config.Name,
		Logger:     config.Logger,
		Port:       config.Port,
		Config:     config,
		Router:     newRouter(),
		signalChan: make(chan os.Signal, 1),
	}

	adapter, err := NewAdapter(robot)
	if err != nil {
		robot.Logger.Error(err)
		return nil, err
	}
	robot.SetAdapter(adapter)

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
	// robot.Logger.
	robot.Logger.Info("Shutting down.")

	return nil
}

// Run starts up the robot
func (robot *Robot) Run() error {
	stop := false

	go robot.Adapter.Run()
	// Start the HTTP server after the adapter, as adapter.Run() adds additional
	// handlers to the router.
	go http.ListenAndServe(`:`+robot.Port, robot.Router)

	signal.Notify(robot.signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for !stop {
		select {
		case sig := <-robot.signalChan:
			switch sig {
			// case syscall.SIGHUP:
			// robot.Logger.Info("Reloading")
			case syscall.SIGINT, syscall.SIGTERM:
				stop = true
			}
		}
	}

	robot.Stop()

	return nil
}

func (robot *Robot) respondRegex(pattern string) string {
	str := `^(?:`
	if robot.Alias != "" {
		str += `(?:` + robot.Alias + `|` + robot.Name + `)`
	} else {
		str += robot.Name
	}
	str += `[:,]?)\s+(?:` + pattern + `)`
	return str
}

// newRouter initializes a new http.ServeMux and sets up several default routes
func newRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/hal/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "PONG")
	})

	router.HandleFunc("/hal/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server time is: %s\n", time.Now())
	})

	return router
}

func (robot *Robot) SetName(name string) {
	robot.Name = name
}

func (robot *Robot) SetLogger(logger *logging.Logger) {
	robot.Logger = logger
}

func (robot *Robot) SetAdapter(adapter Adapter) {
	robot.Adapter = adapter
}

func (robot *Robot) SetPort(port string) {
	robot.Port = port
}

func (robot *Robot) SetRouter(router *http.ServeMux) {
	robot.Router = router
}
