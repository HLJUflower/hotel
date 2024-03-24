package global

import (
	"fyne.io/fyne/v2/app"
	"github.com/jinzhu/gorm"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
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
	App    = new(Application)
	A      = app.New()
	Bundle = i18n.NewBundle(language.English)
)
