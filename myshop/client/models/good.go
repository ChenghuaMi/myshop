package models

type Goods struct {
	Id int
	ShortName string
	LongName string
	BrandId int
	CategoryId int
	ShopId int
	Status byte
	Sale int
	Attribute []Attribute `gorm:"references:CategoryId;foreignKey:CategoryId"`
}

func(g *Goods) TableName() string {
	return "goods"
}
