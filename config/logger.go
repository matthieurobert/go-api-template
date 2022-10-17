package config

import "github.com/sirupsen/logrus"

// InitLogger : create the logrus instance and set configuration
func initLogger(env Env) *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	switch env.LogLevel {
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
		return logger
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
		return logger
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
		return logger
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
		return logger
	case "info":
		logger.SetLevel(logrus.InfoLevel)
		return logger
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
		return logger
	default:
		logger.SetLevel(logrus.InfoLevel)
		return logger
	}
}
