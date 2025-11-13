package logger

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

// InitLogger initializes a colored, structured logger
func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// Colored output for better readability
	logger.SetOutput(colorable.NewColorableStdout())

	// Beautiful formatting
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		PadLevelText:    true,
	})

	// Set log level (can be configured later)
	// Debug: Very detailed, for development
	// Info:  Important events, for production
	// Warn:  Warning messages
	// Error: Error conditions
	logger.SetLevel(logrus.InfoLevel)

	return logger
}


// Example usage:
// log := logger.Init()
// log.Info("Server started")
// log.WithFields(logrus.Fields{
//     "city": "London",
//     "temp": 15.5,
// }).Info("Weather fetched")
