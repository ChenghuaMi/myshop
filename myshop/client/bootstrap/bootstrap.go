package bootstrap

import (
	"myshop/client/api"
	"myshop/client/core/config"
	"myshop/client/global"
	"fmt"
	"myshop/client/models"
)

const filepath = "E:\\go\\src\\myshop\\client\\config\\conf\\config.yml"
func init() {
	config.Viper(&global.Config,filepath)
	LoadDb()
	fmt.Println(global.Config)
	api.InitMiddleware()
	api.InitRouter()
}
func LoadDb() {
	models.InitDb(&global.Config)
}
