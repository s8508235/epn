package log

import (
	log "github.com/sirupsen/logrus"
)

// NewLogger ...
func NewLogger() *log.Logger {
	logger := log.New()

	return logger
}
