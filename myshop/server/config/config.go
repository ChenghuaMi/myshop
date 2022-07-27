package config

type Config struct {
	Rpc	`mapstructure:"rpc"`
	Mysql `mapstructure:"mysql"`
}
