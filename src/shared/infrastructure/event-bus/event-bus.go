package eventbus

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/mmazurowski/gopher-boilerplate/src/shared/infrastructure/aws"
)

type EventBus struct {
	busName         string
	applicationName string
	client          *cloudwatchevents.Client
}

func New(ctx context.Context) (*EventBus, error) {
	client, err := aws.CreateEventBridgeClient(ctx)

	if err != nil {
		return nil, err
	}

	busName := os.Getenv("ENTERPRISE_EVENT_BRIDGE")
	applicationName := os.Getenv("SERVICE_NAME")

	if len(busName) == 0 {
		return nil, errors.New("ENTERPRISE_EVENT_BRIDGE env variable must be present")
	}

	if len(applicationName) == 0 {
		return nil, errors.New("SERVICE_NAME env variable must be present")
	}

	return &EventBus{busName: busName, applicationName: applicationName, client: client}, nil
}

func (bus *EventBus) Publish(ctx context.Context, messages []EventMessage) error {
	entries := make([]types.PutEventsRequestEntry, len(messages))

	for i := range messages {
		message := messages[i]

		entries[i] = types.PutEventsRequestEntry{
			DetailType:   &message.detailType,
			EventBusName: &bus.busName,
			Source:       &bus.applicationName,
			Time:         &message.time,
			Detail:       &message.message,
		}
	}

	input := cloudwatchevents.PutEventsInput{Entries: entries}

	_, err := bus.client.PutEvents(ctx, &input)

	if err != nil {
		return err
	}

	return nil
}
