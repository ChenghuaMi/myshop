package models

type Member struct {
	Id string
	Username string
	Password string
	Status int
	Nickname string
	Phone string
}

func(model *Member) TableName() string {
	return "member"
}
