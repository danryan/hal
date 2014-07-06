package hal

import (
	"github.com/ccding/go-logging/logging"
	"github.com/danryan/env"
	"os"
	"strings"
	"time"
)

// Config struct
type config struct {
	Name        string `env:"key=HAL_NAME default=hal"`
	AdapterName string `env:"key=HAL_ADAPTER default=shell"`
	Port        int    `env:"key=PORT default=9000"`
	LogLevel    string `env:"key=HAL_LOG_LEVEL default=info"`
}

func newConfig() *config {
	c := &config{}
	env.MustProcess(c)
	return c
}

func newLogger() *logging.Logger {
	format := "%25s [%s] %8s: %s\n time,name,levelname,message"
	timeFormat := time.RFC3339
	levelStr := strings.ToUpper(Config.LogLevel)
	level := logging.GetLevelValue(levelStr)
	logger, _ := logging.WriterLogger("hal", level, format, timeFormat, os.Stdout, true)
	return logger
}
