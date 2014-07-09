package memory

import (
	"fmt"
	"github.com/danryan/hal"
)

func init() {
	hal.RegisterStore("memory", New)
}

type store struct {
	hal.BasicStore
	data map[string][]byte
}

// New returns an new initialized store
func New(robot *hal.Robot) (hal.Store, error) {
	s := &store{
		data: map[string][]byte{},
	}
	s.SetRobot(robot)
	return s, nil
}

func (s *store) Open() error {
	return nil
}

func (s *store) Close() error {
	return nil
}

func (s *store) Get(key string) ([]byte, error) {
	if val, ok := s.data[key]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("key %s was not found", key)
}

func (s *store) Set(key string, data []byte) error {
	s.data[key] = data
	return nil
}

func (s *store) Delete(key string) error {
	if _, ok := s.data[key]; !ok {
		return fmt.Errorf("key %s was not found", key)
	}
	delete(s.data, key)
	return nil
}
