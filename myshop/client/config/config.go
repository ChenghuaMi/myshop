package config

type Config struct {
	App `mapstructure:"app"`
	Mysql `mapstructure:"mysql"`
}
