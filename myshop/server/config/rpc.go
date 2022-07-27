package config

type Rpc struct {
	Addr string `mapstructure:"addr"`
	CertFile string `mapstructure:"cert_file"`
	KeyFile string	`mapstructure:"key_file"`
}
