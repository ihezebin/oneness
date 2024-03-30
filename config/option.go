package config

type Option func(c *Config)

func WithDataType(dataType DataType) Option {
	return func(c *Config) {
		c.dataType = dataType
	}
}

func WithDestTagName(tagName DestTagName) Option {
	return func(c *Config) {
		c.destTagName = tagName
	}
}

func WithFileName(fileName string) Option {
	return func(c *Config) {
		c.fileName = fileName
	}

}
