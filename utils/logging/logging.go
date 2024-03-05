package logging

//日志
import (
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

// Setup 初始化日志配置
func Setup() {
	Logger = logrus.New()
}

// Debug 调试
func Debug(v ...interface{}) {
	Logger.Debug(v)
}

func Debugln(v ...interface{}) {
	Logger.Debugln(v)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Info 正常级别日志
func Info(v ...interface{}) {
	Logger.Info(v)
}

func Infoln(v ...interface{}) {
	Logger.Infoln(v)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Warn 警告级别日志
func Warn(v ...interface{}) {
	Logger.Warn(v)
}

func Warnln(v ...interface{}) {
	Logger.Warnln(v)
}

func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Error 错误级别日志
func Error(v ...interface{}) {
	Logger.Error(v)
}

func Errorln(v ...interface{}) {
	Logger.Errorln(v)
}

func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Fatal 致命错误级别日志
func Fatal(v ...interface{}) {
	Logger.Fatal(v)
}

func Fatalln(v ...interface{}) {
	Logger.Fatalln(v)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
