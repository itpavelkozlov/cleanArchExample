package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger interface {
	Info(msg interface{})
	Warn(msg interface{})
	Error(msg interface{})
	Debug(msg interface{})
}

type LoggerImp struct {
	log *logrus.Logger
}

func (l LoggerImp) Info(msg interface{}) {
	l.log.Info(msg)
}

func (l LoggerImp) Warn(msg interface{}) {
	l.log.Warn(msg)
}

func (l LoggerImp) Error(msg interface{}) {
	l.log.Error(msg)
}

func (l LoggerImp) Debug(msg interface{}) {
	l.log.Debug(msg)
}

func NewLogger() Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	return &LoggerImp{
		log: logrus.New(),
	}
}
