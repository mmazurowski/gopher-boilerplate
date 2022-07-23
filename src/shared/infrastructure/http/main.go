package http

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

const internalServerError int = 500
const notFound int = 404
const created int = 201
const success int = 200
const validationError int = 422

func RespondWith(content map[string]string, httpStatus int) events.APIGatewayProxyResponse {
	errorResponse := events.APIGatewayProxyResponse{
		StatusCode: internalServerError,
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
	return RespondWith(content, success), nil
}

func NotFound(content map[string]string) (events.APIGatewayProxyResponse, error) {
	return RespondWith(content, notFound), nil
}

func Created(content map[string]string) (events.APIGatewayProxyResponse, error) {
	return RespondWith(content, created), nil
}

func ValidationError(err error) (events.APIGatewayProxyResponse, error) {
	return RespondWith(map[string]string{
		"type":    "ValidationError",
		"status":  "422",
		"details": err.Error(),
	}, validationError), nil
}

func InternalServerError(err error) (events.APIGatewayProxyResponse, error) {
	return RespondWith(map[string]string{
		"type":    "InternalServerError",
		"status":  "500",
		"details": err.Error(),
	}, internalServerError), nil
}
