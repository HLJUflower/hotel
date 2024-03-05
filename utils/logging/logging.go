package logging

//日志
import (
	"github.com/sirupsen/logrus"
	"hotel/global"
)

// Setup 初始化日志配置
func Setup() {
	global.App.Logger = logrus.New()
}

// Debug 调试
func Debug(v ...interface{}) {
	global.App.Logger.Debug(v)
}

func Debugln(v ...interface{}) {
	global.App.Logger.Debugln(v)
}

func Debugf(format string, args ...interface{}) {
	global.App.Logger.Debugf(format, args...)
}

// Info 正常级别日志
func Info(v ...interface{}) {
	global.App.Logger.Info(v)
}

func Infoln(v ...interface{}) {
	global.App.Logger.Infoln(v)
}

func Infof(format string, args ...interface{}) {
	global.App.Logger.Infof(format, args...)
}

// Warn 警告级别日志
func Warn(v ...interface{}) {
	global.App.Logger.Warn(v)
}

func Warnln(v ...interface{}) {
	global.App.Logger.Warnln(v)
}

func Warnf(format string, args ...interface{}) {
	global.App.Logger.Warnf(format, args...)
}

// Error 错误级别日志
func Error(v ...interface{}) {
	global.App.Logger.Error(v)
}

func Errorln(v ...interface{}) {
	global.App.Logger.Errorln(v)
}

func Errorf(format string, args ...interface{}) {
	global.App.Logger.Errorf(format, args...)
}

// Fatal 致命错误级别日志
func Fatal(v ...interface{}) {
	global.App.Logger.Fatal(v)
}

func Fatalln(v ...interface{}) {
	global.App.Logger.Fatalln(v)
}

func Fatalf(format string, args ...interface{}) {
	global.App.Logger.Fatalf(format, args...)
}
