package proto

type LoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type MemberResp struct {
	Id string
	Username string
}

type MemberInfo interface {
	Login(req LoginReq,res *MemberResp) error
}
