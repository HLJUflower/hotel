package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"hotel/global"
	"hotel/utils/logging"
)

//Viper（处理所有类型的配置需求和格式）初始化

func InitConfig() *viper.Viper {
	//配置文件路径
	config := "config.html"
	//初始化viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	//判断配置文件是否有错误
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	//监听配置文件
	err := v.WatchRemoteConfig()
	if err != nil {
		return nil
	}
	//配置文件重载配置
	v.OnConfigChange(func(in fsnotify.Event) {
		logging.Info("config file changed:", in.Name)
		if err := v.Unmarshal(&global.App.Config); err != nil {
			logging.Errorf("config file begin changed failed:", err)
		}
	})
	return v
}
