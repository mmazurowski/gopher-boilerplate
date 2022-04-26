package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Formatter struct {
}

func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	serviceFromEnv := os.Getenv("SERVICE_NAME")

	serviceName := serviceFromEnv

	if len(serviceFromEnv) == 0 {
		serviceName = "UNKNOWN_APPLICATION"
	}

	message := map[string]any{
		"level":       entry.Level,
		"message":     entry.Message,
		"timestamp":   entry.Time.Format(time.RFC3339),
		"application": serviceName,
	}

	serialized, err := json.Marshal(message)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}

	return append(serialized, '\n'), nil
}

func Factory() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&Formatter{})

	return logger
}
