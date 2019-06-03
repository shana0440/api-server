package config

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func parseLogLevel(level string) (log.Level, error) {
	switch level {
	case "TRACE":
		return log.TraceLevel, nil
	case "DEBUG":
		return log.DebugLevel, nil
	case "INFO":
		return log.InfoLevel, nil
	case "WARN":
		return log.WarnLevel, nil
	case "ERROR":
		return log.ErrorLevel, nil
	default:
		return log.TraceLevel, errors.New("unsupport log level")
	}
}
