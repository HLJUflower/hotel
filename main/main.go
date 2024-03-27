package main

import (
	"github.com/flopp/go-findfont"
	"hotel/api/services"
	"hotel/bootstrap"
	"hotel/config/sql"
	"hotel/global"
	"hotel/utils/logging"
	"hotel/utils/times"
	"os"
	"strings"
	"time"
)

// 提前运行
func init() {
	//启动日志更新
	logging.Setup()
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	global.App.ConfigViper = bootstrap.InitConfig()
	global.App.DB = sql.InitDB()
	defer func() {
		global.App.DB.Close()
		logging.Info("SQL CLOSED")
	}()
	os.Setenv("FYNE_FONT", "./msyh.ttf")
	//时间矫正，更新日志
	logging.Info("Set Time E8")
	localTime := time.FixedZone("CST", 8*3600)
	time.Local = localTime
	logging.Info(
		"Start Time : ",
		times.ToStr(),
	)
	//开始运行并进行监听
	logging.Info("Start Listen")
	services.Select()
}
