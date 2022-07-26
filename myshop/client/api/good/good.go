package good

import (
	"github.com/gin-gonic/gin"
	"myshop/client/form"
	"myshop/client/response"
	"myshop/client/service"
)

func GetAttr(ctx *gin.Context) {
	var goodForm form.Good
	if err := ctx.ShouldBind(&goodForm);err != nil {
		response.ResultFail(500,err.Error(),nil,ctx)
		return
	}
	attrs,err := service.ServiceGood.GetGoodsAttribuetList(goodForm.Id)
	if err != nil {
		response.ResultFail(500,err.Error(),nil,ctx)
		return
	}
	response.ResultOk(200,"ok",attrs,ctx)
}
