package hal

import (
	"fmt"
	"github.com/ccding/go-logging/logging"
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
	robot.Logger.Debugf("%s - robot received message", robot.Adapter.Name())
	for _, handler := range robot.handlers {
		response := NewResponse(robot, msg)
		err := handler.Handle(response)
		if err != nil {
			robot.Logger.Error(err)
			return err
		}
	}
	return nil
}

// Stop initiates the shutdown process
func (robot *Robot) Stop() error {
	robot.Adapter.Stop()
	robot.Logger.Info("Stopping robot.")

	return nil
}

// Run initiates the startup process
func (robot *Robot) Run() error {

	robot.Logger.Info("Starting robot.")
	go robot.Adapter.Run()
	// Start the HTTP server after the adapter, as adapter.Run() adds additional
	// handlers to the router.
	robot.Logger.Debug("Starting HTTP server. ")
	go http.ListenAndServe(`:`+robot.Port, robot.Router)

	signal.Notify(robot.signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	stop := false
	for !stop {
		select {
		case sig := <-robot.signalChan:
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				stop = true
			}
		}
	}
	// Stop listening for new signals
	signal.Stop(robot.signalChan)

	// Initiate the shutdown process for our robot
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
		fmt.Fprintf(w, "Server time is: %s\n", time.Now().UTC())
	})

	return router
}

func (robot *Robot) SetName(name string) {
	robot.Name = name
}

func (robot *Robot) SetAdapter(adapter Adapter) {
	robot.Adapter = adapter
}
