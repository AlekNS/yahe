package logger

import (
	"github.com/alekns/yahe/internal/config"
	"github.com/sirupsen/logrus"
)

var rootLogger = logrus.New()

// Get prepares default logger with specified "source" tag.
func Get(source string) *logrus.Entry {
	return rootLogger.WithFields(logrus.Fields{"source": source})
}

// GetRootLogger returns instantiated root logger.
func GetRootLogger() *logrus.Logger {
	return rootLogger
}

func initCommon(settings *config.LoggerSettings) {
	switch settings.ConsoleLevel {
	case "debug":
		rootLogger.SetLevel(logrus.DebugLevel)
	case "warning":
		rootLogger.SetLevel(logrus.WarnLevel)
	case "info":
		rootLogger.SetLevel(logrus.InfoLevel)
	case "error":
		rootLogger.SetLevel(logrus.ErrorLevel)
	default:
		rootLogger.SetLevel(logrus.InfoLevel)
	}
}

// InitLogger .
func InitLogger(settings *config.LoggerSettings) {
	initCommon(settings)
}
