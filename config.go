package hal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/danryan/env"
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

func newLogger() *logrus.Logger {
	level, err := logrus.ParseLevel(Config.LogLevel)
	if err != nil {
		panic(err)
	}
	logger := logrus.New()
	logger.Level = level
	
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
