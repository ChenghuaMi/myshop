package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func ZipLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger,_ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("test",zap.String("url",ctx.Request.URL.Path),zap.Time("time",time.Now()))
		ctx.Next()
	}
}
