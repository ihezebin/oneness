package logger

import (
	"time"

	"github.com/ihezebin/oneness/logger/hook"
	"github.com/sirupsen/logrus"
)

type Option func(*logrus.Logger)

func WithLocalFsHook(path string) (Option, *hook.LocalFsHook, error) {
	localFsHook, err := hook.NewLocalFsHook(path)
	if err != nil {
		return nil, localFsHook, err
	}
	return func(l *logrus.Logger) {
		l.AddHook(localFsHook)
	}, localFsHook, nil
}
func WithRotateLogsHook(path string, rotateTime time.Duration, expireTime time.Duration) (Option,
	*hook.RotateLogsHook, error) {
	rotateLogsHook, err := hook.NewRotateLogsHook(path, rotateTime, expireTime)
	if err != nil {
		return nil, nil, err
	}
	return func(l *logrus.Logger) {
		l.AddHook(rotateLogsHook)
	}, rotateLogsHook, nil
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
