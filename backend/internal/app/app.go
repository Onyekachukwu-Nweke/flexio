package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
