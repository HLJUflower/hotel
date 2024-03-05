package main

import (
	"hotel/app"
	"hotel/bootstrap"
	"hotel/config/sql"
	"hotel/global"
	"hotel/utils/logging"
	"hotel/utils/times"
	"time"
)

// 提前运行
func init() {
	//启动日志更新
	logging.Setup()
}

func main() {
	global.App.ConfigViper = bootstrap.InitConfig()
	global.App.DB = sql.InitDB()
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
	app.Login()
}
