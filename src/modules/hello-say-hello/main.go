package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/adapters"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/framework"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	app, err := framework.CreateApplication(ctx)

	if err != nil {
		panic(err)
	}

	return adapters.HelloAction(app, event, ctx)
}

func main() {
	lambda.Start(Handler)
}
