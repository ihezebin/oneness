package logger

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = defaultLogger()

func Logger() *logrus.Logger {
	return logger
}

func defaultLogger() *logrus.Logger {
	l := &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.DateTime,
		},
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	return l
}

func ResetLoggerWithOptions(opts ...Option) {
	logger = defaultLogger()
	for _, opt := range opts {
		opt(logger)
	}
}

func WithError(err error) *Entry {
	return &Entry{Entry: logger.WithError(err)}
}

func WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: logger.WithField(key, value)}
}

func WithFields(fields map[string]interface{}) *Entry {
	return &Entry{Entry: logger.WithFields(fields)}
}

func Log(ctx context.Context, level logrus.Level, args ...interface{}) {
	logger.WithContext(ctx).Log(level, args...)
}
func Trace(ctx context.Context, args ...interface{})   { logger.WithContext(ctx).Trace(args...) }
func Debug(ctx context.Context, args ...interface{})   { logger.WithContext(ctx).Debug(args...) }
func Print(ctx context.Context, args ...interface{})   { logger.WithContext(ctx).Print(args...) }
func Info(ctx context.Context, args ...interface{})    { logger.WithContext(ctx).Info(args...) }
func Warn(ctx context.Context, args ...interface{})    { logger.WithContext(ctx).Warn(args...) }
func Warning(ctx context.Context, args ...interface{}) { logger.WithContext(ctx).Warning(args...) }
func Error(ctx context.Context, args ...interface{})   { logger.WithContext(ctx).Error(args...) }
func Panic(ctx context.Context, args ...interface{})   { logger.WithContext(ctx).Panic(args...) }
func Fatal(ctx context.Context, args ...interface{})   { logger.WithContext(ctx).Fatal(args...) }

func Logf(ctx context.Context, level logrus.Level, format string, args ...interface{}) {
	logger.WithContext(ctx).Logf(level, format, args...)
}
func Tracef(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Tracef(format, args...)
}
func Debugf(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Debugf(format, args...)
}
func Printf(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Printf(format, args...)
}
func Infof(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Infof(format, args...)
}
func Warnf(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Warnf(format, args...)
}
func Warningf(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Warningf(format, args...)
}
func Errorf(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Errorf(format, args...)
}
func Panicf(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Panicf(format, args...)
}
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Fatalf(format, args...)
}

func Logln(ctx context.Context, level logrus.Level, args ...interface{}) {
	logger.WithContext(ctx).Logln(level, args...)
}
func Traceln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Traceln(args...)
}
func Debugln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Debugln(args...)
}
func Println(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Println(args...)
}
func Infoln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Infoln(args...)
}
func Warnln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Warnln(args...)
}
func Warningln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Warningln(args...)
}
func Errorln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Errorln(args...)
}
func Panicln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Panicln(args...)
}
func Fatalln(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Fatalln(args...)
}
