package hal

import (
	"fmt"
	// "github.com/ccding/go-logging/logging"
	"net/http"
)

// BasicAdapter declares common functions shared by all adapters
type BasicAdapter struct {
	*Robot
}

func (a *BasicAdapter) SetRobot(r *Robot) {
	a.Robot = r
}

func (a *BasicAdapter) preRun() {
	a.Logger.Infof("Starting %s adapter.", a)
	// TODO: probably not useful for production
	a.Robot.Router.HandleFunc("/hal/adapter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", a)
	})
}

func (a *BasicAdapter) postRun() {
	a.Logger.Infof("Started %s adapter.", a)
}

func (a *BasicAdapter) run() {
	a.preRun()
	a.postRun()
}

func (a *BasicAdapter) preStop() {
	a.Logger.Infof("Stopping %s adapter.", a)
}

func (a *BasicAdapter) postStop() {
	a.Logger.Infof("Stopped %s adapter.", a)
}

func (a *BasicAdapter) stop() {
	a.preStop()
	a.postStop()
}

func (a *BasicAdapter) String() string {
	return a.Robot.Adapter.Name()
}

func (a *BasicAdapter) Name() string {
	return "basic"
}
