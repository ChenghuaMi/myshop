package config

type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host string `mapstructure:"host"`
	Port int 	`mapstructure:"port"`
	Db string `mapstructure:"db"`
}
