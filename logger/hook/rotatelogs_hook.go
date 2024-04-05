package hook

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ logrus.Hook = &RotateLogsHook{}

// NewRotateLogsHook
// errLevel the level which output to err file
func NewRotateLogsHook(path string, rotateTime time.Duration, expireTime time.Duration) (*RotateLogsHook, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, errors.Wrapf(err, "make dir:%s error", dir)
	}

	normalExt := filepath.Ext(path)
	errExt := fmt.Sprintf(".err%s", normalExt)
	errPath := strings.ReplaceAll(path, normalExt, errExt)

	normalWriter, err := newRotateLog(path, rotateTime, expireTime)
	errWriter, err := newRotateLog(errPath, rotateTime, expireTime)
	if err != nil {
		return nil, errors.Wrapf(err, "open error file:%s error", errPath)
	}

	return &RotateLogsHook{
		normalWriter: normalWriter,
		errWriter:    errWriter,
		errLevel:     logrus.ErrorLevel,
	}, nil
}

type RotateLogsHook struct {
	normalWriter, errWriter *rotatelogs.RotateLogs
	errLevel                logrus.Level
}

func newRotateLog(path string, rotateTime time.Duration, expireTime time.Duration) (*rotatelogs.RotateLogs, error) {
	return rotatelogs.New(
		// The pattern used to generate actual log file names.
		// You should use patterns using the strftime (3) format.
		// 分割后的文件名称
		path+".%Y%m%d%H%M%S",
		// Interval between file rotation. By default logs are rotated every 86400 seconds.
		// Note: Remember to use time.Duration values.
		// 设置日志切割时间间隔
		rotatelogs.WithRotationTime(rotateTime),
		// Path where a symlink for the actual log file is placed.
		// This allows you to always check at the same location for log files even if the logs were rotated
		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(filename),
		// Time to wait until old logs are purged. By default no logs are purged,
		// which certainly isn't what you want. Note: Remember to use time.Duration values.
		// 设置最大保存时间
		rotatelogs.WithMaxAge(expireTime),
		// The number of files should be kept. By default, this option is disabled.
		// Note: MaxAge should be disabled by specifing WithMaxAge(-1) explicitly.
		//rotatelogs.WithRotationCount(1),
		// Ensure a new file is created every time New() is called.
		// If the base file name already exists, an implicit rotation is performed.
		rotatelogs.ForceNewFile(),
	)
}

func (l *RotateLogsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *RotateLogsHook) Fire(entry *logrus.Entry) error {
	data, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		return errors.Wrapf(err, "format log error")
	}

	if entry.Level <= l.errLevel {
		_, err = l.errWriter.Write(data)
		if err != nil {
			return errors.Wrapf(err, "write error log error")
		}
	} else {
		_, err = l.normalWriter.Write(data)
		if err != nil {
			return errors.Wrapf(err, "write normal log error")
		}
	}
	return nil
}

func (l *RotateLogsHook) Close() error {
	if l == nil {
		return nil
	}

	if l.normalWriter != nil {
		if err := l.normalWriter.Close(); err != nil {
			return errors.Wrapf(err, "close normal writer error")
		}
	}

	if l.errWriter != nil {
		if err := l.errWriter.Close(); err != nil {
			return errors.Wrapf(err, "close err writer error")
		}
	}

	return nil
}
