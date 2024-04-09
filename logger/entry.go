package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Entry struct {
	*logrus.Entry
}

func (e *Entry) WithError(err error) *Entry {
	return &Entry{Entry: e.Entry.WithError(err)}
}

func (e *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: e.Entry.WithField(key, value)}
}

func (e *Entry) WithFields(fields map[string]interface{}) *Entry {
	return &Entry{Entry: e.Entry.WithFields(fields)}
}

func (e *Entry) Log(ctx context.Context, level logrus.Level, args ...interface{}) {
	e.Entry.WithContext(ctx).Log(level, args...)
}
func (e *Entry) Trace(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Trace(args...)
}
func (e *Entry) Debug(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Debug(args...)
}
func (e *Entry) Print(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Print(args...)
}
func (e *Entry) Info(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Info(args...)
}
func (e *Entry) Warn(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Warn(args...)
}
func (e *Entry) Warning(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Warning(args...)
}
func (e *Entry) Error(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Error(args...)
}
func (e *Entry) Panic(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Panic(args...)
}
func (e *Entry) Fatal(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Fatal(args...)
}

func (e *Entry) Logf(ctx context.Context, level logrus.Level, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Logf(level, format, args...)
}
func (e *Entry) Tracef(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Tracef(format, args...)
}
func (e *Entry) Debugf(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Debugf(format, args...)
}
func (e *Entry) Printf(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Printf(format, args...)
}
func (e *Entry) Infof(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Infof(format, args...)
}
func (e *Entry) Warnf(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Warnf(format, args...)
}
func (e *Entry) Warningf(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Warningf(format, args...)
}
func (e *Entry) Errorf(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Errorf(format, args...)
}
func (e *Entry) Panicf(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Panicf(format, args...)
}
func (e *Entry) Fatalf(ctx context.Context, format string, args ...interface{}) {
	e.Entry.WithContext(ctx).Fatalf(format, args...)
}

func (e *Entry) Logln(ctx context.Context, level logrus.Level, args ...interface{}) {
	e.Entry.WithContext(ctx).Logln(level, args...)
}
func (e *Entry) Traceln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Traceln(args...)
}
func (e *Entry) Debugln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Debugln(args...)
}
func (e *Entry) Println(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Println(args...)
}
func (e *Entry) Infoln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Infoln(args...)
}
func (e *Entry) Warnln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Warnln(args...)
}
func (e *Entry) Warningln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Warningln(args...)
}
func (e *Entry) Errorln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Errorln(args...)
}
func (e *Entry) Panicln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Panicln(args...)
}
func (e *Entry) Fatalln(ctx context.Context, args ...interface{}) {
	e.Entry.WithContext(ctx).Fatalln(args...)
}
