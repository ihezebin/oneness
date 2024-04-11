package logger

import (
	"strings"
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
		l.AddHook(hook.NewCallerHook(false))
	}
}
func WithCallerPrettyHook() Option {
	return func(l *logrus.Logger) {
		l.AddHook(hook.NewCallerHook(false))
	}
}

func WithTimestampHook() Option {
	return func(l *logrus.Logger) {
		l.AddHook(hook.NewTimestampHook())
	}
}

type Level string

const (
	LevelPanic   Level = "panic"
	LevelFatal   Level = "fatal"
	LevelError   Level = "error"
	LevelWarn    Level = "warn"
	LevelWarning Level = "warning"
	LevelInfo    Level = "info"
	LevelDebug   Level = "debug"
	LevelTrace   Level = "trace"
)

func WithLevel(level Level) Option {
	level2LogrusLevel := func(level Level) logrus.Level {
		switch Level(strings.ToLower(string(level))) {
		case LevelPanic:
			return logrus.PanicLevel
		case LevelFatal:
			return logrus.FatalLevel
		case LevelError:
			return logrus.ErrorLevel
		case LevelWarn, LevelWarning:
			return logrus.WarnLevel
		case LevelInfo:
			return logrus.InfoLevel
		case LevelDebug:
			return logrus.DebugLevel
		case LevelTrace:
			return logrus.TraceLevel
		default:
			return logrus.InfoLevel
		}
	}

	return func(l *logrus.Logger) {
		l.SetLevel(level2LogrusLevel(level))
	}
}

func WithServiceName(serviceName string) Option {
	return func(l *logrus.Logger) {
		l.AddHook(hook.NewFieldsHook(logrus.Fields{
			hook.FieldKeyServiceName: serviceName,
		}))
	}
}
