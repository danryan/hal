package redis

import (
	"github.com/danryan/hal"
)

func init() {
	hal.RegisterStore("redis", New)
}

type store struct {
	hal.BasicStore
}

func (s *store) Name() string { return "redis" }

// New returns an initialized store
func New(robot *hal.Robot) (hal.Store, error) {
	s := &store{}
	s.SetRobot(robot)
	return s, nil
}

func (s *store) Open() error {
	return nil
}
func (s *store) Close() error {
	return nil
}
func (s *store) Load() error {
	return nil
}
func (s *store) Save() error {
	return nil
}
func (s *store) Get(key string) ([]byte, error) {
	return []byte{}, nil
}
func (s *store) Set(key string, data []byte) error {
	return nil
}
func (s *store) Delete(key string) error {
	return nil
}
