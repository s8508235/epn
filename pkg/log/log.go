package log

import (
	filename "github.com/keepeye/logrus-filename"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// NewLogger ...
func NewLogger() *log.Logger {
	logger := log.New()
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true

	// Set specific colors for prefix and timestamp
	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "cyan+bh",
		TimestampStyle: "black+b:168",
	})
	logger.SetFormatter(formatter)
	logger.AddHook(filename.NewHook())
	return logger
}
