package adapters

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/app/commands"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/framework"
)

func HelloAction(app *framework.Application, event events.APIGatewayProxyRequest, context *lambdacontext.LambdaContext) (events.APIGatewayProxyResponse, error) {
	payload, err := FromAPIGatewayEvent(event)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, err
	}

	cmd := commands.NewSayHelloCommand(payload.Name)

	result, _ := app.Bus.Handle(cmd)

	output := result.(string)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       output,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, err
}
