package config

import (
	"strings"
	"testing"
)

func TestConfigFile(t *testing.T) {
	config := NewWithFilePath("config.json")
	t.Log(config.FilePaths())
	dest := make(map[string]interface{})
	if err := config.Load(&dest); err != nil {
		t.Fatal(err)
	}

	t.Log(dest)
}

func TestConfigReader(t *testing.T) {
	data := `{
  "name": "oneness",
  "version": "1.0.0"
}`
	config := NewWithReader(strings.NewReader(data))
	dest := make(map[string]interface{})
	if err := config.Load(&dest); err != nil {
		t.Fatal(err)
	}

	t.Log(dest)
}
