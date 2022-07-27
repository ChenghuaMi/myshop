package rpc

import (
	"crypto/tls"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"fmt"
)

type TlsServer struct {
	opt *tlsoption
}
func NewTlsServer(chg ...ChangeOption) *TlsServer {
	opt := defaultTlsOption
	for _,item := range chg {
		item.apply(&opt)
	}
	return &TlsServer{
		opt: &opt,
	}
}

func(l *TlsServer) Register(server interface{}) error {
	return rpc.Register(server)
}
func(l *TlsServer) Run(addr string) {
	lis,err := l.Listen(addr)
	if err != nil {
		return
	}
	for  {
		conn,err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
func(l *TlsServer) Listen(addr string) (net.Listener,error) {
	if !l.opt.openssl {
		return net.Listen("tcp",addr)
	}
	fmt.Println("tls.........................")
	cert,err := tls.LoadX509KeyPair("rpc/keys/server.pem","rpc/keys/server.key")
	if err != nil {
		panic(err)
	}
	conf := tls.Config{Certificates: []tls.Certificate{cert}}
	return tls.Listen("tcp",addr,&conf)
}