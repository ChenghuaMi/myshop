package logic

import (
	"gorm.io/gorm"
	"myshop/server/global"
	"myshop/server/models"
	"myshop/server/proto"
	comm_error "myshop/server/error"
)

type Member struct {

}
var LogicMember = new(Member)
func (m *Member) Login(req proto.LoginReq,resp *proto.MemberResp) error {
	md,err := m.FindOne(req.Username)
	if err != nil {
		return err
	}
	*resp=proto.MemberResp{
		Id:       md.Id,
		Username: md.Username,
	}
	return nil
}

func (m *Member) FindOne(username string) (*models.Member,error) {
	var memberObj models.Member
	err := global.Db.Where("username=?",username).First(&memberObj).Error
	if err == gorm.ErrRecordNotFound {
		return &memberObj,comm_error.ErrorUserNotExist
	}
	if err != nil {
		return nil,comm_error.ErrorSearch
	}
	return &memberObj,nil

}