package bootstrap

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myshop/server/core"
	"myshop/server/global"
	"myshop/server/models"
	"myshop/server/rpc"
)
const DSN="%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	LoadConfig()
	LoadMysql()
	loadRpc()
}
func LoadConfig() {
	core.LoadViper()
}
func LoadMysql() {
	dsn := fmt.Sprintf(DSN,global.Config.Mysql.Username,global.Config.Mysql.Password,global.Config.Mysql.Host,global.Config.Mysql.Port,global.Config.Mysql.Dbname)
	fmt.Println(dsn)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("err:" + err.Error())
	}
	models.CreateUid(dsn)
	global.Db = db
}


func loadRpc() {
	var changeOption []rpc.ChangeOption
	changeOption = append(changeOption,rpc.AddTlsOptionFunc("rpc/keys/server.pem","rpc/keys/server.key"))
	global.TlsServer = rpc.NewTlsServer(changeOption...)
}