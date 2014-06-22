package gobot

import (
	"log"
	"os"

	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/kelseyhightower/envconfig"
)

type Config struct {
	Name        string
	AdapterName string
	Logger      *log.Logger
	// AdapterName string
}

// TODO: panic if required env vars are not present
// https://github.com/paulhammond/slackcat/blob/master/slackcat.go#L26-L52
func NewConfig() *Config {
	config := &Config{}
	config.Name = "gobot"
	config.AdapterName = "shell"
	config.Logger = log.New(os.Stdout, "[gobot] ", 0)

	return config
}
