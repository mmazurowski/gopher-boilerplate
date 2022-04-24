package framework

import (
	"github.com/mmazurowski/gopher-boilerplate/src/framework/cqrs"
	"github.com/mmazurowski/gopher-boilerplate/src/framework/logger"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/app/commands"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	Bus    cqrs.Bus
	Logger *log.Logger
}

func CreateApplication() *Application {
	logInstance := logger.Factory()
	handlers := []cqrs.Handler{commands.NewSayHelloCommandHandler(logInstance)}

	return &Application{
		Bus:    cqrs.Factory(handlers),
		Logger: logInstance,
	}
}
