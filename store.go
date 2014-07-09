package hal

import (
	"fmt"
)

// Store interface for storage backends to implement
type Store interface {
	Open() error
	Close() error
	Get(string) ([]byte, error)
	Set(key string, data []byte) error
	Delete(string) error
}

type store struct {
	name    string
	newFunc func(*Robot) (Store, error)
}

// BasicStore struct to be embedded in other stores
type BasicStore struct {
	Robot *Robot
}

// SetRobot sets the adapter's Robot
func (s *BasicStore) SetRobot(r *Robot) {
	s.Robot = r
}

// Stores is a map of registered stores
var Stores = map[string]store{}

// RegisterStore registers a new store
func RegisterStore(name string, newFunc func(*Robot) (Store, error)) {
	Stores[name] = store{
		name:    name,
		newFunc: newFunc,
	}
}

// NewStore returns an initialized store
func NewStore(robot *Robot) (Store, error) {
	name := Config.StoreName
	if _, ok := Stores[name]; !ok {
		return nil, fmt.Errorf("%s is not a registered store", Config.StoreName)
	}

	store, err := Stores[name].newFunc(robot)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *BasicStore) String() string {
	return Config.StoreName
}
