package config

type Redis struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}
