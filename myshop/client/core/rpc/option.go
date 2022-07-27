package rpc

import "time"

var (
	DefaultPoolSize = 10
	DefaultPoolTTl = 10 * time.Second
)
type diaOption struct {
	poolSize int
	poolTTl time.Duration
	servers map[string]*Server
}

type Server struct {
	CertFile string
	TlsServerName string
	Openssl bool
	Network string
	Address string
}

func NewDiaOption() *diaOption {
	return &diaOption{
		poolSize: DefaultPoolSize,
		poolTTl: DefaultPoolTTl,
		servers:  make(map[string]*Server),
	}
}

type DiaOption interface {
	apply(option *diaOption)
}

type fnDialOption struct {
	f func(option *diaOption)
}
func (fn *fnDialOption) apply(option *diaOption) {
	fn.f(option)
}
func newFnDialOption(f func(option *diaOption)) *fnDialOption {
	return &fnDialOption{
		f: f,
	}
}
func SetDialOption(serverName string,s *Server) DiaOption {
	return newFnDialOption(func(option *diaOption) {
		if s.CertFile != "" && s.TlsServerName != "" {
			s.Openssl = true
		}
		option.servers[serverName] = s
	})
}