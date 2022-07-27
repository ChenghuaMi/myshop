package rpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/rpc"
	"net/rpc/jsonrpc"
	"fmt"
)
type RpcClient interface {
	NewConnect(serverName string) (Conn,error)
}
type rpcClient struct {
	opt *diaOption
	mp *managerPool
}

func NewRpcClient(dia ...DiaOption) *rpcClient {
	opts := NewDiaOption()

	for _,di := range dia {
		di.apply(opts)
	}
	client := &rpcClient{
		opt: opts,
		mp:  NewManagerPool(),
	}
	poolOpt := &PoolOption{
		PoolSize:        opts.poolSize,
		PoolTTl: 		 opts.poolTTl,
	}
	for serverName,server := range opts.servers {
		poolOpt.CreateConnHandle = client.createConn(serverName,server)
		poo,err := initPool(poolOpt)
		if err != nil {
			fmt.Println("create pool err:" + err.Error())
		}
		client.mp.AddPool(serverName,poo)
	}
	return client
}

func (cli *rpcClient) createConn(serverName string,s *Server) CreateConnHandle {
	fmt.Println(">>>>>")
	return func() (Conn, error) {
		client,err := cli.getClient(s)
		if err != nil {
			fmt.Println("serverName connect error:" + serverName)
			return &Connect{
				client: nil,
			},err
		}
		return &Connect{client: client},nil
	}
}
func(cli *rpcClient) getClient(s *Server) (*rpc.Client,error) {
	if !s.Openssl {
		return jsonrpc.Dial(s.Network,s.Address)
	}
	fmt.Println("tls auth check...................")
	bytes,err := ioutil.ReadFile(s.CertFile)
	if err != nil {
		return nil,errors.New("读取文件失败")
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(bytes)
	conn,err := tls.Dial(s.Network,s.Address,&tls.Config{
		RootCAs: certPool,
		ServerName: s.TlsServerName,
	})
	return jsonrpc.NewClient(conn),err
}

func(cli *rpcClient) NewConnect(serverName string) (Conn,error) {
	if len(cli.opt.servers) > 0 {
		po,ok := cli.mp.GetPool(serverName)
		if !ok {
			return nil,errors.New("连接池 " + serverName + "不存在")
		}
		ctx,_ := context.WithTimeout(context.Background(),10)
		conn,err := po.GetConn(ctx)
		if err != nil {
			return nil,err
		}
		return conn,nil
	}
	return &Connect{},nil
}