package client

import (
	"go.uber.org/zap"
	"myshop/client/global"
)

type CommonClient struct {
	*MemberClient
}
var CommClient = new(CommonClient)


func call(serverName string,methodName string,req interface{},resp interface{}) error {
	conn,err := global.RpcClient.NewConnect(serverName)
	if err != nil {
		global.Log.Error("创建连接失败 :" + serverName,zap.Error(err))
		return err
	}
	err = conn.Call(methodName,req,resp)
	if err != nil {
		global.Log.Error("rpc 调度 失败：" + serverName + ":" + methodName,zap.Error(err))
		return err
	}
	return nil
}