package adapters

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/app/commands"
	"github.com/mmazurowski/gopher-boilerplate/src/modules/hello-say-hello/framework"
	"github.com/mmazurowski/gopher-boilerplate/src/shared/api"
)

func HelloAction(app *framework.Application, event events.APIGatewayProxyRequest, ctx context.Context) (events.APIGatewayProxyResponse, error) {
	payload, err := FromAPIGatewayEvent(event)

	if err != nil {
		return api.RespondWith(map[string]string{}, 500)
	}

	cmd := commands.NewSayHelloCommand(payload.Name)

	result, err := app.Bus.Handle(ctx, cmd)

	if err != nil {
		return api.RespondWith(map[string]string{}, 500)
	}

	output := result.(map[string]string)

	return api.RespondWith(output, 200)
}
