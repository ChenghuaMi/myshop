package bootstrap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"myshop/client/api"
	"myshop/client/cache"
	"myshop/client/core/config"
	"myshop/client/global"
	"myshop/client/models"
	"time"
)

const filepath = "E:\\go\\src\\myshop\\client\\config\\conf\\config.yml"
func init() {
	config.Viper(&global.Config,filepath)
	LoadCache()
	LoadZap()
	LoadDb()
	api.InitMiddleware()
	api.InitRouter()
}
func LoadZap() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime= func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	wt := zapcore.AddSync(WriteInfo())
	core := zapcore.NewCore(encoder,wt,zapcore.InfoLevel)
	global.Log = zap.New(core,zap.AddCaller(),zap.Hooks())
	//global.Log.Info("test",zap.String("uri","baidu.com"))
}
func LoadCache() {
	switch global.Config.Cache.CacheType {
	case "freetype":
		cache.CacheObj = cache.NewCache(cache.NewFreeCache())
	case "redis":
		cache.CacheObj = cache.NewCache(cache.NewRedis())
	default:
		cache.CacheObj = cache.NewCache(cache.NewFreeCache())
	}
	//cache.CacheObj.Set("name",[]byte("mch"),&cache.Options{})
	//result,err := cache.CacheObj.Get("name")
	//fmt.Println(string(result.([]byte)),err)
}
func LoadDb() {
	models.InitDb(&global.Config)
}
func WriteInfo() io.Writer {
	return &lumberjack.Logger{
		Filename:   global.Config.Log.Filename,
		MaxSize:    global.Config.Log.MaxSize,
		MaxAge:     global.Config.Log.MaxAge,
		MaxBackups: global.Config.Log.MaxBackups,
		LocalTime:  false,
		Compress:   false,
	}
}
