package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int
	Msg string
	Data interface{}
}

func ResultOk(code int,msg string,data interface{},ctx *gin.Context) {
	ctx.JSON(200,&Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
func ResultFail(code int,msg string,data interface{},ctx *gin.Context) {
	ctx.JSON(500,&Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
