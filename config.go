package hal

import (
	"fmt"
	"github.com/ccding/go-logging/logging"
	"github.com/danryan/env"
	"net/http"
	"os"
	"strings"
	"time"
)

// Config struct
type config struct {
	Name        string `env:"key=HAL_NAME default=hal"`
	Alias       string `env:"key=HAL_ALIAS"`
	AdapterName string `env:"key=HAL_ADAPTER default=shell"`
	StoreName   string `env:"key=HAL_STORE default=memory"`
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

// newRouter initializes a new http.ServeMux and sets up several default routes
func newRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/hal/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "PONG")
	})

	router.HandleFunc("/hal/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server time is: %s\n", time.Now().UTC())
	})

	return router
}
