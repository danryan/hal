package hal

import (
	"github.com/ccding/go-logging/logging"
	"os"
	"strings"
	"time"
)

// Config struct
type Config struct {
	Name        string
	AdapterName string
	Logger      *logging.Logger
	Port        string
}

// TODO: panic if required env vars are not present
// https://github.com/paulhammond/slackcat/blob/master/slackcat.go#L26-L52
func NewConfig() *Config {
	return &Config{
		Name:        GetenvDefault("HAL_NAME", "hal"),
		AdapterName: GetenvDefault("HAL_ADAPTER", "shell"),
		Logger:      newLogger(),
		Port:        GetenvDefault("PORT", "9000"),
	}
}

// GetenvDefault tests if an environment variable is set, and returns a default
// value if the variable is an empty string.
func GetenvDefault(args ...string) string {
	if len(args) < 2 {
		return "" //, errors.New("GetenvDefault requires two arguments (envVar string, defaultValue string)")
	}
	key := args[0]
	def := args[1]
	str := os.Getenv(key)
	if str == "" {
		return def
	}
	return str
}

func newLogger() *logging.Logger {
	format := "%25s [%s] %8s: %s\n time,name,levelname,message"
	timeFormat := time.RFC3339
	levelStr := strings.ToUpper(GetenvDefault("HAL_LOG_LEVEL", "INFO"))
	level := logging.GetLevelValue(levelStr)
	logger, _ := logging.WriterLogger("hal", level, format, timeFormat, os.Stdout, true)
	return logger
}
