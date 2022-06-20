package utils

import "github.com/sirupsen/logrus"

// Logger interface acts as a contract for a specific logger.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

func NewLogrusLogger() Logger {
	log := logrus.New()
	log.Level = logrus.InfoLevel
	return log
}
