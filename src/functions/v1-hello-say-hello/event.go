package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator"
)

type GreetEvent struct {
	Name string `json:"name" validate:"required"`
}

func parseGatewayEvent(event events.APIGatewayProxyRequest) (*GreetEvent, error) {
	var payload GreetEvent
	err := json.Unmarshal([]byte(event.Body), &payload)

	if err != nil {
		return nil, err
	}

	validate := validator.New()

	err = validate.Struct(payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}
