package proto

type LoginReq struct {
	Username string
	Password string
}

type MemberResp struct {
	Id string
	Username string
}

type MemberInfo interface {
	Login(req LoginReq,res *MemberResp) error
}
