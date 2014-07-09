package redis

import (
	"fmt"
	"github.com/danryan/env"
	"github.com/danryan/hal"
	"github.com/davecgh/go-spew/spew"
	"github.com/garyburd/redigo/redis"
	"net/url"
)

var _ = spew.Sdump()

func init() {
	hal.RegisterStore("redis", New)
}

type store struct {
	hal.BasicStore
	config *config
	client redis.Conn
}

type config struct {
	URL       string `env:"key=HAL_REDIS_URL default=redis://localhost:6379"`
	Namespace string `env:"key=HAL_REDIS_NAMESPACE default=hal"`
}

// New returns an new initialized store
func New(robot *hal.Robot) (hal.Store, error) {
	c := &config{}
	env.MustProcess(c)
	s := &store{
		config: c,
	}
	s.SetRobot(robot)
	return s, nil
}

func (s *store) Open() error {
	uri, err := url.Parse(s.config.URL)
	if err != nil {
		hal.Logger.Error(err)
	}

	conn, err := redis.Dial("tcp", uri.Host)
	if err != nil {
		hal.Logger.Error(err)
		return err
	}
	s.client = conn
	return nil
}

func (s *store) Close() error {
	if err := s.client.Close(); err != nil {
		hal.Logger.Error(err)
		return err
	}
	return nil
}

func (s *store) Get(key string) ([]byte, error) {
	args := s.namespace(key)
	data, err := s.client.Do("GET", args)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return []byte{}, fmt.Errorf("%s not found", key)
	}
	return data.([]byte), nil
}

func (s *store) Set(key string, data []byte) error {
	if _, err := s.client.Do("SET", s.namespace(key), data); err != nil {
		return err
	}
	return nil
}

func (s *store) Delete(key string) error {
	res, err := s.client.Do("DEL", s.namespace(key))
	if err != nil {
		return err
	}
	if res.(int64) < 1 {
		return fmt.Errorf("%s not found", key)
	}
	return nil
}

func (s *store) namespace(key string) string {
	return fmt.Sprintf("%s:%s", s.config.Namespace, key)
}
