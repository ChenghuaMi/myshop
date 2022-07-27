package rpc

import (
	"net/rpc"
	"time"
)

type Conn interface {
	Call(methodName string,req interface{},resp interface{}) error
	Close() error
	CreateTime() time.Time
}

type CreateConnHandle func() (Conn,error)

type Connect struct {
	client *rpc.Client
	create_time time.Time
}
func (con *Connect) Call(methodName string,req interface{},resp interface{}) error {
	return con.client.Call(methodName,req,resp)
}
func(con *Connect) Close() error {
	return con.client.Close()
}
func(con *Connect) CreateTime() time.Time {
	return con.create_time
}