package hook

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type CallerHook struct {
	prettyFilename bool
}

var _ logrus.Hook = &CallerHook{}

func NewCallerHook(prettyFilename bool) logrus.Hook {
	return &CallerHook{
		prettyFilename: prettyFilename,
	}
}

func (s *CallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (s *CallerHook) Fire(entry *logrus.Entry) error {
	caller := getCaller()
	if caller != nil {
		entry.Data[logrus.FieldKeyFunc] = caller.Function
		filename := caller.File
		if s.prettyFilename {
			filename = path.Base(filename)
		}

		entry.Data[logrus.FieldKeyFile] = fmt.Sprintf("%s:%d", filename, caller.Line)
	}
	return nil
}

func getCaller() *runtime.Frame {
	const maxStackDepth = 32
	pcs := make([]uintptr, maxStackDepth)
	// 跳过Callers本身和getCaller函数的帧，通常设置为2。
	n := runtime.Callers(2, pcs)
	if n == 0 {
		return nil
	}

	frames := runtime.CallersFrames(pcs[:n])
	more := true
	var frame runtime.Frame
	for more {
		frame, more = frames.Next()
		if !strings.Contains(frame.Function, "logrus") && !strings.Contains(frame.Function, "logger") {
			// 找到了一个有效的调用者帧
			return &frame // 直接返回当前帧的指针，不必复制。
		}
	}

	return nil
}
