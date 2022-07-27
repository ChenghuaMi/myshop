package rpc

import (
	"context"
	"errors"
	"time"
)


type Pool interface {
	GetConn(ctx context.Context) (Conn,error)
}

type managerPool struct {
	Pools map[string]Pool
}
func NewManagerPool() *managerPool {
	return &managerPool{Pools: make(map[string]Pool)}
}
func (m *managerPool) AddPool(name string,p Pool) {
	m.Pools[name] = p
}
func (m *managerPool) GetPool(name string) (Pool,bool) {
	p,ok := m.Pools[name]
	return p,ok
}

type PoolOption struct {
	PoolSize int
	PoolTTl time.Duration
	CreateConnHandle
}

type pool struct {
	poolSize int
	poolTTl time.Duration
	Conns chan Conn
	CreateConnHandle
}
func initPool(op *PoolOption) (*pool,error) {
	if op.PoolSize == 0 {
		return nil,errors.New("连接池大小未设置")
	}
	p := &pool{
		poolSize:         op.PoolSize,
		poolTTl: 		  op.PoolTTl,
		Conns:            make(chan Conn,op.PoolSize),
		CreateConnHandle: op.CreateConnHandle,
	}
	return p,p.init()
}
func(p *pool) init() error {
	if p.CreateConnHandle == nil {
		return errors.New("创建连接池函数未设置")
	}
	for i := 0;i<p.poolSize;i++ {
		conn,err := p.CreateConnHandle()
		if err != nil {
			return err
		}
		p.Conns <- conn
	}
	return nil
}
func(p *pool) GetConn(ctx context.Context) (Conn,error) {
	for  {
		select {
		case <-ctx.Done():
			return nil,errors.New("连接超时。。。。")
		case conn := <-p.Conns:
			if d := time.Since(conn.CreateTime());d > p.poolTTl {
				conn.Close()
				p.CreateConn()
			}
			return conn,nil
		}
	}
}
func(p *pool) CreateConn(){
	go func() {
		for  {
			if p.CreateConnHandle == nil {
				return
			}
			conn,err := p.CreateConnHandle()
			if err != nil {
				continue
			}
			p.Conns <-conn
			return
		}
	}()
}
