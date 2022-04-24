package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/adapters"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/framework"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	app := framework.CreateApplication()
	lambdaContext, _ := lambdacontext.FromContext(ctx)

	return adapters.HelloAction(app, event, lambdaContext)
}

func main() {
	lambda.Start(Handler)
}
