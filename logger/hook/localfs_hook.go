package hook

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ logrus.Hook = &LocalFsHook{}

// NewLocalFsHook
// errLevel the level which output to err file
func NewLocalFsHook(path string) (*LocalFsHook, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, errors.Wrapf(err, "make dir:%s error", dir)
	}

	normalExt := filepath.Ext(path)
	errExt := fmt.Sprintf(".err%s", normalExt)
	errPath := strings.ReplaceAll(path, normalExt, errExt)

	normalWriter, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.Wrapf(err, "open normal file:%s error", path)
	}
	errWriter, err := os.OpenFile(errPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.Wrapf(err, "open error file:%s error", errPath)
	}

	return &LocalFsHook{
		normalWriter: normalWriter,
		errWriter:    errWriter,
		errLevel:     logrus.ErrorLevel,
	}, nil
}

type LocalFsHook struct {
	normalWriter, errWriter *os.File
	errLevel                logrus.Level
}

func (l *LocalFsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LocalFsHook) Fire(entry *logrus.Entry) error {
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

func (l *LocalFsHook) Close() error {
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
