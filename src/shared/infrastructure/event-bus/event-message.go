package eventbus

import (
	"encoding/json"
	"errors"
	time2 "time"
)

type EventMessage struct {
	detailType string
	message    string
	time       time2.Time
}

func CreateEvent(detailType string, message map[string]any) (*EventMessage, error) {
	serialized, err := json.Marshal(message)

	if err != nil {
		return nil, errors.New("Could not parse event payload for detail: [" + detailType + "]")
	}

	return &EventMessage{detailType: detailType, message: string(serialized), time: time2.Now()}, nil
}
