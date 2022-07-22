package http

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func RespondWith(content map[string]string, httpStatus int) events.APIGatewayProxyResponse {
	errorResponse := events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	var buf bytes.Buffer

	body, err := json.Marshal(content)

	if err != nil {
		return errorResponse
	}

	json.HTMLEscape(&buf, body)

	response := events.APIGatewayProxyResponse{
		StatusCode: httpStatus,
		Body:       buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return response
}

func Success(content map[string]string) (events.APIGatewayProxyResponse, error) {
	return RespondWith(content, 200), nil
}

func NotFound(content map[string]string) (events.APIGatewayProxyResponse, error) {
	return RespondWith(content, 404), nil
}

func Created(content map[string]string) (events.APIGatewayProxyResponse, error) {
	return RespondWith(content, 201), nil
}

func ValidationError(err error) (events.APIGatewayProxyResponse, error) {
	return RespondWith(map[string]string{
		"type":    "ValidationError",
		"status":  "422",
		"details": err.Error(),
	}, 422), nil
}

func InternalServerError(err error) (events.APIGatewayProxyResponse, error) {
	return RespondWith(map[string]string{
		"type":    "InternalServerError",
		"status":  "500",
		"details": err.Error(),
	}, 500), nil
}
