package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myshop/client/global"
	"time"
)

func ZipLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		end := time.Now()
		cost := end.Sub(start)
		status := ctx.Writer.Status()
		fields := []zap.Field{
			zap.String("path",ctx.Request.URL.Path),
			zap.Time("start time",start),
			zap.Time("end time",end),
			zap.Duration("cost",cost),
			zap.Int("status",status),
			zap.String("error",ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		}
		if status == 200 {
			global.Log.Info(ctx.Request.RequestURI,fields...)
		}else{
			global.Log.Error(ctx.Request.RequestURI,fields...)
		}

	}
}
