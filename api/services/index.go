package services

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"hotel/api/services/middleware"
	"hotel/config/api"
	"hotel/global"
)

var mySpace = layout.NewSpacer()

// 住户登记
// 房间查询
// 季度计划
// 员工查询
// 财务汇总

func Index(u api.HotelRole) {
	w := global.A.NewWindow("index")
	wid, hig := 900, 600
	w.Resize(fyne.NewSize(float32(wid), float32(hig)))

	tab := container.NewAppTabs(
		//季度计划
		container.NewTabItem("Plan", container.New(layout.NewGridLayout(1), middleware.JobPlan(u))),
		//财务报表
		container.NewTabItem("Money", container.New(layout.NewGridLayout(1), middleware.MoneyTotal(u))),
		//客户管理
		container.NewTabItem("Client", container.New(layout.NewGridLayout(1), middleware.SleepUserLog(u))),
		//人员管理
		container.NewTabItem("Manage", container.New(layout.NewGridLayout(1), middleware.User(u))),
		//房间
		container.NewTabItem("Register", container.New(layout.NewGridLayout(1), middleware.Room(u))),
	)
	tab.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tab)
	w.Show()
}
