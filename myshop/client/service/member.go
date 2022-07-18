package service

import (
	"myshop/client/form"
	"myshop/client/logic"
	"myshop/client/models"
)

type Member struct {

}
var MemberService = new(Member)
func(m *Member) Login(lg form.Login)(*models.Member,error) {
	return logic.MemberLogic.Login(lg)
}
func(m *Member) Register(lg form.Register) (*models.Member,error) {
	return logic.MemberLogic.Register(lg)
}
