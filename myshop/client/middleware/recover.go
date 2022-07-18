package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover();err != nil {
				logger,_ := zap.NewProduction()
				defer logger.Sync()
				logger.Error("error",zap.String("url",ctx.Request.URL.Path),zap.Time("time",time.Now()))
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
