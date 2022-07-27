package models

type Member struct {
	Id string
	Username string
	Password string
	Status int
	Nickname string
	Phone string
}
func (*Member) TableName() string {
	return "member"
}
