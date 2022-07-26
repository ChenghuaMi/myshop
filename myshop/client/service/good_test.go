package service

import (
	"myshop/client/core/config"
	"myshop/client/global"
	"myshop/client/models"
	"testing"
)

const filepath = "E:\\go\\src\\myshop\\client\\config\\conf\\config.yml"
func init() {
	config.Viper(&global.Config,filepath)
	models.InitDb(&global.Config)
}
func TestGood(t *testing.T) {
	ServiceGood.GetGoodsAttribuetList(1)
}
