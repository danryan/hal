package hal

import (
	"os"
)

// BasicAdapter declares common functions shared by all adapters
type BasicAdapter struct {
	robot      *Robot
	signalChan chan os.Signal
}

// Robot returns the adapter's robot instance
func (a *BasicAdapter) Robot() *Robot {
	return a.robot
}

// SetRobot sets the adapter's robot instance
func (a *BasicAdapter) SetRobot(r *Robot) {
	a.robot = r
}
