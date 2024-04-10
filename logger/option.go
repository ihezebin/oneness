package logger

import (
	"time"

	"github.com/ihezebin/oneness/logger/hook"
	"github.com/sirupsen/logrus"
)

type Option func(*logrus.Logger)

func WithLocalFsHook(path string) Option {
	localFsHook, err := hook.NewLocalFsHook(path)
	if err != nil {
		panic(err)
	}
	return func(l *logrus.Logger) {
		l.AddHook(localFsHook)
	}
}
func WithRotateLogsHook(path string, rotateTime time.Duration, expireTime time.Duration) Option {
	rotateLogsHook, err := hook.NewRotateLogsHook(path, rotateTime, expireTime)
	if err != nil {
		panic(err)
	}
	return func(l *logrus.Logger) {
		l.AddHook(rotateLogsHook)
	}
}

func WithCallerHook() Option {
	return func(l *logrus.Logger) {
		l.AddHook(hook.NewCallerHook())
	}
}

func WithTimestampHook() Option {
	return func(l *logrus.Logger) {
		l.AddHook(hook.NewTimestampHook())
	}
}

func WithLevel(level logrus.Level) Option {
	return func(l *logrus.Logger) {
		l.SetLevel(level)
	}
}

func WithServiceName(serviceName string) Option {
	return func(l *logrus.Logger) {
		l.AddHook(hook.NewFieldsHook(logrus.Fields{
			hook.FieldKeyServiceName: serviceName,
		}))
	}
}
