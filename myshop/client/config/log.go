package config

type Log struct {
	Filename  string `mapstructure:"filename"`
	MaxSize   int `mapstructure:"maxsize"`
	MaxAge    int `mapstructure:"maxage"`
	MaxBackups int `mapstructure:"maxbackups"`
}
