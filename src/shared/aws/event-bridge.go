package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func CreateEventBridgeClient(ctx context.Context) (*cloudwatchevents.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return nil, err
	}

	return cloudwatchevents.NewFromConfig(cfg), nil
}
