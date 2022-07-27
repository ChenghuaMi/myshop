package global

import (
	"gorm.io/gorm"
	"myshop/server/config"
	"myshop/server/rpc"
)

var (
	TlsServer *rpc.TlsServer
	Config config.Config
	Db *gorm.DB
)
