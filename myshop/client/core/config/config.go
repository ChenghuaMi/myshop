package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"myshop/client/config"
)

func Viper(config *config.Config,configFile string) *viper.Viper {
	vip := viper.New()
	vip.SetConfigFile(configFile)
	vip.SetConfigType("yml")
	if err := vip.ReadInConfig();err != nil {
		panic(fmt.Sprint("fail error %s \n",err.Error()))
	}
	vip.WatchConfig()
	cfg := func(vip *viper.Viper) {
		vip.Unmarshal(config)
	}
	vip.OnConfigChange(func(in fsnotify.Event) {
		cfg(vip)
	})
	cfg(vip)
	return vip
}
