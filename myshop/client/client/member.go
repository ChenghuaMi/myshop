package client

import "myshop/client/proto"

type MemberClient struct {

}

func(m *MemberClient) Login(req proto.LoginReq,res *proto.MemberResp) error {
	return call("member","Member.Login",req,res)
}