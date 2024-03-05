package global

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"hotel/config"
)

// Application 项目启动时的变量
type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Logger      *logrus.Logger
}

// 启动项实例化
var App = new(Application)
