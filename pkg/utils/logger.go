package utils

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger(service string) *Logger {

	var logger = &Logger{logrus.WithFields(logrus.Fields{"service": service})}

	return logger
}

type Type string

const (
	Platform Type = "Platform"
	ADB      Type = "DB"
)

func (logger Logger) WithError(errType Type, err error) Logger {
	return Logger{logger.WithFields(logrus.Fields{"error": err, "Type": errType})}
}

func (logger Logger) WithToken(requestToken string) Logger {
	return Logger{logger.WithField("requestToken", requestToken)}
}
