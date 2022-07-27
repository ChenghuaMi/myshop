package server

import (
	"myshop/server/logic"
	"myshop/server/proto"
	"fmt"
)

type Member struct {

}
var ServiceMember = new(Member)

func (m *Member) Login(req proto.LoginReq,resp *proto.MemberResp) error {
	fmt.Println("rpc.........................")
	return logic.LogicMember.Login(req,resp)
}
