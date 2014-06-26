package hal

import (
	"github.com/ccding/go-logging/logging"
	"os"
	"strconv"
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
func GetenvDefault(key string, def string) string {
	// if len(args) < 2 {
	// 	return "" //, errors.New("GetenvDefault requires two arguments (envVar string, defaultValue string)")
	// }
	// key := args[0]
	// def := args[1]
	str := os.Getenv(key)
	if str == "" {
		return def
	}
	return str
}

// GetenvConvert attempts to convert an environment variable into type _t_,
// and returns nil if the variable is an empty string (not set)
// func GetenvConvert(key string, t interface{}) interface{} {
// 	str := strings.ToLower(os.Getenv(key))
// 	if str == "" {
// 		return nil
// 	}
// 	switch t.(type) {
// 	case int:
// 		i, err := strconv.Atoi(str)
// 		if err != nil {
// 			return nil
// 		}
// 		return i
// 	case float64:
// 		f, err := strconv.ParseFloat(str, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return f
// 	case bool:
// 		b, err := strconv.ParseBool(str)
// 		if err != nil {
// 			return nil
// 		}
// 		return b
// 	}
// 	return nil
// }

// GetenvConvertDefault attempts to convert an environment variable into type _t_,
// falling back to a default _def_
// func GetenvConvertDefault(key string, t, def interface{}) interface{} {
// 	env := GetenvConvert(key, t)
// 	if env == nil {
// 		return def
// 	}
// 	return env
// }

// func GetenvDefaultString(key string, def string) string {
// 	env := os.Getenv(key)
// 	if env == "" {
// 		return def
// 	}
// 	return env
// }

func GetenvDefaultInt(key string, def int) int {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	i, _ := strconv.Atoi(env)
	return i
}

func GetenvDefaultBool(key string, def bool) bool {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	i, _ := strconv.ParseBool(env)
	return i
}

// GetenvConvert with errors
// func GetenvConvert(key string, i interface{}) (interface{}, error) {
// 	str := strings.ToLower(os.Getenv(key))
// 	if str == "" {
// 		return str, errors.New(key + ` environment variable is not set`)
// 	}
// 	switch i.(type) {
// 	case int:
// 		return strconv.Atoi(str)
// 	case bool:
// 		return strconv.ParseBool(str)
// 	}
// 	return nil, errors.New("cannot convert string to unsupported type")
// }

func newLogger() *logging.Logger {
	format := "%25s [%s] %8s: %s\n time,name,levelname,message"
	timeFormat := time.RFC3339
	levelStr := strings.ToUpper(GetenvDefault("HAL_LOG_LEVEL", "INFO"))
	level := logging.GetLevelValue(levelStr)
	logger, _ := logging.WriterLogger("hal", level, format, timeFormat, os.Stdout, true)
	return logger
}
