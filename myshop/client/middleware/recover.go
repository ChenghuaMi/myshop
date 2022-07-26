package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myshop/client/global"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover();err != nil {
				global.Log.Error(ctx.Request.URL.Path,
					zap.Any("error",err),
					zap.String("method",ctx.Request.Method),
					)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
