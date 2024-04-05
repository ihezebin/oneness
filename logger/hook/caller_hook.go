package hook

import "C"
import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

type CallerHook struct{}

var _ logrus.Hook = &CallerHook{}

func NewCallerHook() logrus.Hook {
	return &CallerHook{}
}

func (s *CallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (s *CallerHook) Fire(entry *logrus.Entry) error {
	caller := getCaller()
	if caller != nil {
		entry.Data[logrus.FieldKeyFunc] = caller.Function
		entry.Data[logrus.FieldKeyFile] = fmt.Sprintf("%s:%d", caller.File, caller.Line)
	}
	return nil
}

const skipCallerDepth = 9

func getCaller() *runtime.Frame {
	pc, file, line, _ := runtime.Caller(skipCallerDepth)
	fn := runtime.FuncForPC(pc).Name()
	return &runtime.Frame{
		Line:     line,
		File:     file,
		Function: fn,
	}
}
