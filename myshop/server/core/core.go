package core

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"fmt"
	"myshop/server/global"
)

const configFile = "E:\\go\\src\\myshop\\server\\config\\conf.yml"
func LoadViper() *viper.Viper{
	vip := viper.New()
	vip.SetConfigFile(configFile)
	vip.SetConfigType("yml")
	if err := vip.ReadInConfig();err != nil {
		panic(fmt.Sprintf("err %s \n",err.Error()))
	}
	vip.WatchConfig()
	cfg := func(v *viper.Viper) {
		v.Unmarshal(&global.Config)
	}
	vip.OnConfigChange(func(in fsnotify.Event) {
		cfg(vip)
	})
	cfg(vip)
	fmt.Println(global.Config)
	return vip
}
