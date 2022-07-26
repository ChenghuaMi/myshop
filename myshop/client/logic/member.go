package logic

import (
	"gorm.io/gorm"
	er "myshop/client/errors"
	"myshop/client/form"
	"myshop/client/models"
	"myshop/client/util"
)


type Member struct {

}
var MemberLogic = new(Member)

func(m *Member) Register(reg form.Register) (*models.Member,error){
	one,err := m.FindOne(reg.Username)
	if err != nil {
		return one,er.UserExists
	}
	mod := &models.Member{
		Id:       models.GetWuid(),
		Username: reg.Username,
		Password: util.Md5(reg.Password),
		Status:   0,
		Nickname: "",
		Phone:    "",
	}
	if err := models.DB.Create(mod).Error;err != nil {
		return nil,er.CreateUserFail
	}

	return mod,nil
}
func(m *Member) FindOne(username string) (*models.Member,error) {
	member := &models.Member{}
	err := models.DB.Where("username=?",username).First(member).Error
	if err == gorm.ErrRecordNotFound {
		return member,nil
	}
	return member,er.UserExists
}
func(m *Member) Login(lg form.Login)(*models.Member,error) {
	member := &models.Member{}
	err := models.DB.Where("username=? and password=?",lg.Username,util.Md5(lg.Password)).First(member).Error
	if err == gorm.ErrRecordNotFound {
		return member,er.UserUnExist
	}
	return member,nil
}
