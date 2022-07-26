package models

import (
	"myshop/client/core/config"
	"myshop/client/global"
	"testing"
)
const filepath = "E:\\go\\src\\myshop\\client\\config\\conf\\config.yml"
func init() {
	config.Viper(&global.Config,filepath)
	InitDb(&global.Config)
}
func TestAddAttribute(t *testing.T) {
	attrs := []Attribute{
		{Name: "颜色", CategoryId: 1, Sort: 1,AttributeDetail: []AttributeDetail{
			{Id: 1,Name: "红色"},
			{Id: 2,Name: "白色"},
			{Id: 3,Name: "绿色"},
		}},
		{Name: "配置", CategoryId: 1, Sort: 2,AttributeDetail: []AttributeDetail{
			{Id: 4,Name: "I5"},
			{Id: 5,Name: "I3"},
			{Id: 6,Name: "I7"},
		}},
		{Name: "显卡", CategoryId: 1, Sort: 3,AttributeDetail: []AttributeDetail{
			{Id: 7,Name: "独显"},
			{Id: 8,Name: "集成"},
		}},
		{Name: "存储容量", CategoryId: 2, Sort: 2,AttributeDetail: []AttributeDetail{
		{Id: 9,Name: "500G"},
		{Id: 10,Name: "1T"},
		}},
	}
	DB.Create(&attrs)
}
