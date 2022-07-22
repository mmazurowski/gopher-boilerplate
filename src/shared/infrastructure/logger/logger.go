package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mmazurowski/gopher-boilerplate/src/shared/domain/features"
	log "github.com/sirupsen/logrus"
)

type Formatter struct {
	feature string
}

func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	serviceFromEnv := os.Getenv("SERVICE_NAME")
	serviceVersion := os.Getenv("SERVICE_VERSION")
	environment := os.Getenv("SERVICE_ENV")
	functionName := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")

	serviceName := serviceFromEnv

	if len(serviceFromEnv) == 0 {
		serviceName = "UNKNOWN_APPLICATION"
	}

	message := map[string]any{
		"level":        entry.Level,
		"message":      entry.Message,
		"timestamp":    entry.Time.Format(time.RFC3339),
		"application":  serviceName,
		"env":          environment,
		"version":      serviceVersion,
		"functionName": functionName,
	}

	if len(f.feature) > 0 {
		message["feature"] = f.feature
	}

	serialized, err := json.Marshal(message)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}

	return append(serialized, '\n'), nil
}

func CreateFeatureLogger(featureTag features.Features) *log.Logger {
	logger := log.New()
	logger.SetFormatter(&Formatter{feature: featureTag.String()})

	return logger
}

func Create() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&Formatter{})

	return logger
}
