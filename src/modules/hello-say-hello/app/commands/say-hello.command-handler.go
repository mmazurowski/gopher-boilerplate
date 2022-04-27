package commands

import (
	"context"
	"fmt"

	eventbus "github.com/mmazurowski/gopher-boilerplate/src/shared/event-bus"
	log "github.com/sirupsen/logrus"
)

type SayHelloCommandHandler struct {
	logger   *log.Logger
	eventBus *eventbus.EventBus
}

func NewSayHelloCommandHandler(logger *log.Logger, eventBus *eventbus.EventBus) SayHelloCommandHandler {
	return SayHelloCommandHandler{logger: logger, eventBus: eventBus}
}

func (h SayHelloCommandHandler) Name() string {
	return "SAY_HELLO_COMMAND"
}

func (h SayHelloCommandHandler) Handle(ctx context.Context, plainCmd any) any {
	cmd := plainCmd.(SayHelloCommand)

	h.logger.Info("Received command " + cmd.Name())

	message := map[string]any{
		"Message": fmt.Sprintf("Welcome: %s", cmd.name),
	}

	domainEvent, _ := eventbus.CreateEvent("greetings.welcome", message)

	eventSeries := []eventbus.EventMessage{*domainEvent}

	err := h.eventBus.Publish(ctx, eventSeries)

	if err != nil {
		h.logger.Error(err)
		panic(err)
	}

	return map[string]string{
		"Message": fmt.Sprintf("Welcome: %s", cmd.name),
	}
}
