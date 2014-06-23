package hal

import (
	"fmt"
	"net/http"
)

// BasicAdapter declares common functions shared by all adapters
type BasicAdapter struct {
	robot *Robot
}

// Robot returns the adapter's robot instance
func (a *BasicAdapter) Robot() *Robot {
	return a.robot
}

// SetRobot sets the adapter's robot instance
func (a *BasicAdapter) SetRobot(r *Robot) {
	a.robot = r
}

func (a *BasicAdapter) preRun() {
	a.robot.Logger.Infof("Starting %s adapter.", a)
	// TODO: probably not useful for production
	a.robot.Router.HandleFunc("/hal/adapter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", a)
	})
}

func (a *BasicAdapter) postRun() {
	a.robot.Logger.Infof("Started %s adapter.", a)
}

func (a *BasicAdapter) run() {
	a.preRun()
	a.postRun()
}

func (a *BasicAdapter) preStop() {
	a.robot.Logger.Infof("Stopping %s adapter.", a)
}

func (a *BasicAdapter) postStop() {
	a.robot.Logger.Infof("Stopped %s adapter.", a)
}

func (a *BasicAdapter) stop() {
	a.preStop()
	a.postStop()
}

func (a *BasicAdapter) String() string {
	return a.robot.Adapter.Name()
}

func (a *BasicAdapter) Name() string {
	return "basic"
}
