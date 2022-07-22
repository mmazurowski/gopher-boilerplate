package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	eventbus "github.com/mmazurowski/gopher-boilerplate/src/shared/infrastructure/event-bus"
	"github.com/mmazurowski/gopher-boilerplate/src/shared/infrastructure/http"
	log "github.com/sirupsen/logrus"
)

type dependencies struct {
	config   *config
	logger   *log.Logger
	eventBus *eventbus.EventBus
}

func handler(ctx context.Context, gatewayEvent events.APIGatewayProxyRequest, deps dependencies) (events.APIGatewayProxyResponse, error) {
	deps.logger.Info("Hello there! " + deps.config.appName)

	greetEvent, err := parseGatewayEvent(gatewayEvent)

	if err != nil {
		return http.ValidationError(err)
	}

	message := map[string]any{
		"Message": fmt.Sprintf("Welcome: %s", greetEvent.Name),
	}

	domainEvent, _ := eventbus.CreateEvent("greetings.welcome", message)

	eventSeries := []eventbus.EventMessage{*domainEvent}

	err = deps.eventBus.Publish(ctx, eventSeries)

	if err != nil {
		return http.InternalServerError(err)
	}

	response := map[string]string{"hello": greetEvent.Name}

	return http.Success(response)
}
