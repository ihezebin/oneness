package hook

import (
	"github.com/sirupsen/logrus"
)

const FieldKeyServiceName = "service"

type FieldsHook struct {
	Fields logrus.Fields
}

var _ logrus.Hook = &FieldsHook{}

// NewFieldsHook Use to create the FieldsHook
// Used to add common log attributes
// 用于添加公共的日志属性
func NewFieldsHook(fields logrus.Fields) *FieldsHook {
	return &FieldsHook{
		Fields: fields,
	}
}

// Levels implement levels
func (hook *FieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire implement fire
func (hook *FieldsHook) Fire(entry *logrus.Entry) error {
	for k, v := range hook.Fields {
		entry.Data[k] = v
	}
	return nil
}
