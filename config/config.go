package config

import (
	"io"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	kernel      *viper.Viper
	destTagName DestTagName
	dataType    DataType
	fileName    string
	filePaths   []string
	reader      io.Reader
}

func NewWithFilePath(path string, opts ...Option) *Config {
	config := &Config{
		destTagName: defaultDestTagName,
	}
	for _, opt := range opts {
		opt(config)
	}

	// if not sure file type, use json
	if config.dataType == "" {
		config.dataType = DataTypeJson
	}

	filePaths := make([]string, 0)
	// if path is relative path, find from current working directory and app executable path
	if !filepath.IsAbs(path) {
		workPath, err := os.Getwd()
		if err == nil {
			workFilePath := filepath.Join(workPath, path)
			filePaths = append(filePaths, workFilePath)
		}
		appPath := filepath.Dir(os.Args[0])
		appFilePath := filepath.Join(appPath, path)
		filePaths = append(filePaths, appFilePath)
	} else {
		filePaths = append(filePaths, path)
	}

	allDir := true
	kernel := viper.New()
	for _, filePath := range filePaths {
		stat, err := os.Stat(filePath)
		if err == nil && stat != nil {
			if !stat.IsDir() {
				kernel.SetConfigFile(filePath)
				allDir = false
				break
			}
			kernel.AddConfigPath(filePath)
		}
	}
	// if all paths are directories, need set config name and type
	if allDir {
		if config.fileName == "" {
			config.fileName = defaultFileName
		}
		kernel.SetConfigName(config.fileName)
		kernel.SetConfigType(string(config.dataType))
	}

	config.filePaths = filePaths
	config.kernel = kernel

	return config
}

func NewWithReader(reader io.Reader, opts ...Option) *Config {
	config := &Config{
		reader:      reader,
		destTagName: defaultDestTagName,
	}

	for _, opt := range opts {
		opt(config)
	}

	if config.dataType == "" {
		config.dataType = DataTypeJson
	}

	kernel := viper.New()
	// reader need know data type
	kernel.SetConfigType(string(config.dataType))

	config.kernel = kernel

	return config
}

func (c *Config) Kernel() *viper.Viper {
	return c.kernel
}

func (c *Config) Load(dest interface{}) error {
	if c.kernel.ConfigFileUsed() != "" || len(c.filePaths) > 0 {
		if err := c.kernel.ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to load config file path")
		}
	}
	if c.reader != nil {
		if err := c.kernel.MergeConfig(c.reader); err != nil {
			return errors.Wrap(err, "failed to load config reader")
		}
	}

	return c.Kernel().Unmarshal(dest, func(d *mapstructure.DecoderConfig) {
		d.TagName = string(c.destTagName)
	})
}

func (c *Config) FilePaths() []string {
	return c.filePaths
}
