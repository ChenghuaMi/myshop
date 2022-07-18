package middleware

import (
	"github.com/gin-gonic/gin"
	"myshop/client/response"
	"fmt"
	"myshop/client/util"
)

func ParseMid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		result,err := util.ParseToken(token)
		fmt.Println("err:",err)
		fmt.Println(result)
		if err != nil {
			fmt.Println("?????????????????????????????????")
			response.ResultFail(500,err.Error(),nil,ctx)
			ctx.Abort()
		}
		ctx.Set("user_id",result.(*util.UserClaims).Id)
		ctx.Next()
	}
}