package member

import (
	"github.com/gin-gonic/gin"
	"myshop/client/client"
	"myshop/client/form"
	"myshop/client/proto"
	"myshop/client/response"
	"fmt"
	"myshop/client/service"
	"myshop/client/util"
)
func Login(c *gin.Context) {
	var lv form.Login
	if err := c.ShouldBind(&lv);err != nil {
		fmt.Println("abbbb",err.Error())
		response.ResultFail(500,err.Error(),nil,c)
		return
	}
	mod,err := service.MemberService.Login(lv)
	if err != nil {
		response.ResultFail(500,err.Error(),nil,c)
		return
	}
	tokenString,err := util.CreateToken(mod.Id)
	if err != nil {
		response.ResultFail(500,err.Error(),nil,c)
		return
	}

	response.ResultOk(200,"ok",tokenString,c)
}
func Register(c *gin.Context) {
	var lv form.Register
	if err := c.ShouldBind(&lv);err != nil {
		response.ResultFail(500,err.Error(),nil,c)
		return
	}
	mod,err := service.MemberService.Register(lv)
	if err != nil {
		response.ResultFail(500,err.Error(),nil,c)
		return
	}
	response.ResultOk(200,"ok",mod,c)


}
func RpcLogin(c *gin.Context) {
	var lv form.Login
	if err := c.ShouldBind(&lv);err != nil {
		response.ResultFail(500,err.Error(),nil,c)
		return
	}
	var resp proto.MemberResp
	err := client.CommClient.MemberClient.Login(proto.LoginReq{
		Username: lv.Username,
		Password: lv.Password,
	},&resp)
	if err != nil {
		response.ResultFail(500,err.Error(),nil,c)
		return
	}
	response.ResultOk(200,"ok",resp,c)
}
func User(c *gin.Context) {
	fmt.Println("user....................")
	fmt.Println(c.Get("user_id"))
}


