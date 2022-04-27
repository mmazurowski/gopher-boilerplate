package framework

import (
	"context"

	"github.com/mmazurowski/gopher-boilerplate/src/framework/cqrs"
	"github.com/mmazurowski/gopher-boilerplate/src/framework/logger"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/app/commands"
	eventbus "github.com/mmazurowski/gopher-boilerplate/src/shared/event-bus"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	Bus    cqrs.Bus
	Logger *log.Logger
}

func CreateApplication(ctx context.Context) (app *Application, err error) {
	logInstance := logger.Factory()
	eb, err := eventbus.New(ctx)

	if err != nil {
		return nil, err
	}

	handlers := []cqrs.Handler{commands.NewSayHelloCommandHandler(logInstance, eb)}

	return &Application{
		Bus:    cqrs.Factory(handlers),
		Logger: logInstance,
	}, nil
}
