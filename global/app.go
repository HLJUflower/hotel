package global

import (
	"github.com/spf13/viper"
	"hotel/config"
)

// Application 项目启动时的变量
type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

// 启动项实例化
var App = new(Application)
