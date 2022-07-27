package global

import (
	"go.uber.org/zap"
	config2 "myshop/client/config"
	"myshop/client/core/rpc"
)

var (
	Config config2.Config
	Log *zap.Logger
	RpcClient rpc.RpcClient
)
