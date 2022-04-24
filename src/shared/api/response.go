package api

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func RespondWith(content map[string]string, httpStatus int) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(content)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	json.HTMLEscape(&buf, body)

	response := events.APIGatewayProxyResponse{
		StatusCode: httpStatus,
		Body:       buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return events.APIGatewayProxyResponse(response), err
}
