package log

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// NewLogger ...
func NewLogger() *log.Logger {
	logger := log.New()
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true

	return logger
}
