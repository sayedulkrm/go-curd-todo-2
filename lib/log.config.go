package lib

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ConfigureLogger() {

	// Configure logrus to use the provided format
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceColors:     true,
		},
	}

	// Set the logrus logger as the default logger for Logrus
	logrus.SetFormatter(logger.Formatter)
	logrus.SetOutput(logger.Out)
	logrus.SetLevel(logger.Level)
}
