package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmazurowski/gopher-boilerplate/src/shared/domain/features"
	eventbus "github.com/mmazurowski/gopher-boilerplate/src/shared/infrastructure/event-bus"
	"github.com/mmazurowski/gopher-boilerplate/src/shared/infrastructure/http"
	"github.com/mmazurowski/gopher-boilerplate/src/shared/infrastructure/logger"
)

func execution(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	featureTag := features.Greeter()
	log := logger.CreateFeatureLogger(featureTag)
	config := createConfig()
	eventBus, err := eventbus.New(ctx)

	if err != nil {
		return http.InternalServerError(err)
	}

	deps := dependencies{logger: log, eventBus: eventBus, config: config}

	return handler(ctx, event, deps)
}

func main() {
	lambda.Start(execution)
}
