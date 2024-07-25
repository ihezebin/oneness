package hook

import (
	"github.com/sirupsen/logrus"
)

type timestampHook struct{}

const FieldKeyTimestamp = "timestamp"

var _ logrus.Hook = &timestampHook{}

func NewTimestampHook() logrus.Hook {
	return &timestampHook{}
}

func (t timestampHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (t timestampHook) Fire(entry *logrus.Entry) error {
	entry.Data[FieldKeyTimestamp] = entry.Time.Unix()
	return nil
}
