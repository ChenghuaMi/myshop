package main

import (
	_ "myshop/server/bootstrap"
	"myshop/server/global"
	"myshop/server/server"
)
func main() {

	global.TlsServer.Register(server.ServiceMember)
	global.TlsServer.Run(global.Config.Rpc.Addr)
}
