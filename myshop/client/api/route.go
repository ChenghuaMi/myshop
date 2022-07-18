package api

import (
	"github.com/gin-gonic/gin"
	"myshop/client/api/member"
	"myshop/client/core/route"
	"myshop/client/middleware"
)

func InitRouter() {
	route.RegisterRoute(MemberApi)
}
func InitMiddleware() {
	route.RegisterMiddleware(middleware.ZipLog(),middleware.ZipLog(),middleware.Recover())
}

func MemberApi(g *gin.Engine) {
	group := g.Group("member")
	{
		group.POST("/login",member.Login)
		group.POST("/register",member.Register)
		group.GET("/user",middleware.ParseMid(),member.User)
	}
}