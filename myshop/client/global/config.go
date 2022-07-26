package global

import (
	"go.uber.org/zap"
	config2 "myshop/client/config"
)

var (
	Config config2.Config
	Log *zap.Logger
)
