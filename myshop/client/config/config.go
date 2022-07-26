package config

type Config struct {
	App `mapstructure:"app"`
	Mysql `mapstructure:"mysql"`
	Log
	Cache `mapstructure:"cache"`
}
