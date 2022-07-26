package models

type Attribute struct {
	Id int
	Name string
	CategoryId int
	Sort int
	AttributeDetail []AttributeDetail `gorm:"foreignKey:AttributeId"`
}

func(a *Attribute) TableName() string {
	return "attribute"
}


type AttributeDetail struct {
	Id int
	Name string
	AttributeId int
	Sort int
	GoodsId int
}
func (d *AttributeDetail) TableName() string {
	return "attribute_detail"
}