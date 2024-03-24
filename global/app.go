package global

import (
	"fyne.io/fyne/v2/app"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

// Application 项目启动时的变量
type Application struct {
	ConfigViper *viper.Viper
	Config      Configuration
	Logger      *logrus.Logger
	DB          *gorm.DB
}

var (
	// 启动项实例化
	App = new(Application)
	A   = app.New()
	WG  = sync.WaitGroup{}
)
