package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/conf"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var cfg struct {
		Web struct {
			APIHost         string        `conf:"defatult:0.0.0.0:3000"`
			DebugHost       string        `conf:"defatult:0.0.0.0:4000"`
			ReadTimeout     time.Duration `conf:"defatult:5s"`
			WriteTimeout    time.Duration `conf:"defatult:5s"`
			ShutdownTimeout time.Duration `conf:"defatult:5s"`
		}
	}

	if err := conf.Parse(os.Args[1:], "SALES", &cfg); err != nil {

	}

	return nil
}
