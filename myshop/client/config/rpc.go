package config

type Rpc struct {
	Servers map[string]Service `mapstructure:"servers"`
}
type Service struct {
	CertFile string				`mapstructure:"cert_file"`
	TlsServerName string		`mapstructure:"tls_server_name"`
	Network string				`mapstructure:"network"`
	Address string				`mapstructure:"address"`
}
