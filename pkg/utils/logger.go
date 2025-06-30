package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a new logger with common configuration
func NewLogger() *logrus.Logger {
	logger := logrus.New()

	// Set log level based on environment (default to Info)
	level := logrus.InfoLevel
	if os.Getenv("DEBUG") == "true" {
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)

	// Configure formatter to include timestamps and colors
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})

	// Write logs to both stdout and file if LOG_FILE is specified
	logFile := os.Getenv("LOG_FILE")
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			// Use multi-writer to write to both stdout and file
			logger.SetOutput(file)
		} else {
			logger.Infof("Failed to log to file %s, using default stderr", logFile)
		}
	}

	return logger
}

// GetLoggerWithField creates a new entry with a field added
func GetLoggerWithField(logger *logrus.Logger, key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// GetLoggerWithFields creates a new entry with multiple fields added
func GetLoggerWithFields(logger *logrus.Logger, fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}
