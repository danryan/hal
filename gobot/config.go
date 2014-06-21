package gobot

import (
	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type Config struct {
	Name string
	Adapter
	Logger *log.Logger
	// AdapterName string
}

// TODO: panic if required env vars are not present
// https://github.com/paulhammond/slackcat/blob/master/slackcat.go#L26-L52
func NewConfig() *Config {
	config := &Config{}
	config.Name = "gobot"
	config.Adapter = &ShellAdapter{}
	config.Logger = log.New(os.Stdout, "[gobot] ", 0)

	return config
}
