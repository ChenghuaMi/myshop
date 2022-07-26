package service

import (
	"gorm.io/gorm"
	"myshop/client/models"
)

type GoodService struct {

}
var ServiceGood = new(GoodService)
func(s *GoodService) GetGoodsAttribuetList(id int) (attrs []models.Attribute,err error) {
	var gd models.Goods
	err = models.DB.Where("id=?",id).Preload("Attribute", func(db *gorm.DB) *gorm.DB{
		return db.Preload("AttributeDetail", func(db *gorm.DB) *gorm.DB{
			return db.Where("goods_id=?",id).Order("sort desc")
		})
	}).Find(&gd).Error
	//fmt.Printf("%#v\n",gd)
	attrs = gd.Attribute
	return
}
