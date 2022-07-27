package rpc

type tlsoption struct {
	certFile string
	keyFile string
	openssl bool
}
var defaultTlsOption = tlsoption{}
type ChangeOption interface {
	apply(opt *tlsoption)
}

type ChangeFunc struct {
	f func(opt *tlsoption)
}

func(c *ChangeFunc) apply(opt *tlsoption) {
	c.f(opt)
}

func newChangeFunc(f func(opt *tlsoption)) *ChangeFunc {
	return &ChangeFunc{f: f}
}
func AddTlsOptionFunc(certFile,keyFile string) ChangeOption {
	return newChangeFunc(func(opt *tlsoption) {
		opt.certFile = certFile
		opt.keyFile = keyFile
		if opt.certFile != "" && opt.keyFile != "" {
			opt.openssl = true
		}
	})
}