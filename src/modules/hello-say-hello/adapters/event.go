package adapters

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator"
)

type Event struct {
	Name string `json:"name" validate:"required"`
}

func FromAPIGatewayEvent(event events.APIGatewayProxyRequest) (*Event, error) {
	var payload Event
	err := json.Unmarshal([]byte(event.Body), &payload)

	if err != nil {
		return nil, errors.New(" could not decode payload" + err.Error())
	}

	validate := validator.New()

	err = validate.Struct(payload)

	if err != nil {
		return nil, errors.New(" payload did not pass validation" + err.Error())
	}

	return &payload, err
}
