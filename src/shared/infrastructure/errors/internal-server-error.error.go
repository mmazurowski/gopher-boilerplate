package errors

import (
	"bytes"
	"encoding/json"
)

type internalServerError struct {
	StatusCode int
	Message    string
	Slug       string
}

func (err internalServerError) Error() string {
	var buf bytes.Buffer
	res, marshalError := json.Marshal(err)

	if marshalError != nil {
		panic("Could not marshall error to string")
	}

	json.HTMLEscape(&buf, res)

	return buf.String()
}

func InternalServerError(message string) error {
	return internalServerError{StatusCode: 500, Slug: "INTERNAL_SERVER_ERROR", Message: message}
}
