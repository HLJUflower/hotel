package middleware

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"hotel/config/api"
	"hotel/global"
	"time"
)

// JobPlan 季度计划
func JobPlan(u api.HotelRole) *fyne.Container {
	var mySpace = layout.NewSpacer()
	tk := api.Task{}
	tks := make([]api.Task, 10)
	tkss := make([][5]string, 10)
	num := 0
	now := time.Now().Format("2006.01.02 15:04:05")
	tkss[0][0] = "JobName"
	tkss[0][1] = "JobGroup"
	tkss[0][2] = "JobDetail"
	tkss[0][3] = "Remark"
	tkss[0][4] = "Status"
	for i := 1; i < 10; i++ {
		global.App.DB.Table("task").Where("create_time < ?", now).Last(&tks[i])
		tkss[i][0] = tks[i].JobName
		tkss[i][1] = tks[i].JobGroup
		tkss[i][2] = tks[i].JobDetail
		tkss[i][3] = tks[i].Remark
		tkss[i][4] = tks[i].Status
		fmt.Println(1)
		global.App.DB.Table("task").Where("create_time < ?", now).First(&tk)
		if tk == tks[i] {
			num = i
			break
		} else {
			now = tks[i].CreateTime.Format("2006.01.02 15:04:05")
		}
	}
	global.App.DB.Table("task").Last(&tk)
	tkgb := widget.NewLabel(tk.JobDetail)
	tkpb := widget.NewLabel("This Issue : ")
	tklay := container.New(layout.NewGridLayout(5), mySpace, tkpb, mySpace, tkgb, mySpace)
	tklayout := container.New(layout.NewGridLayout(1), mySpace, tklay, mySpace)
	list := widget.NewTable(
		func() (int, int) {
			return num + 2, 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Past Plans")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tkss[i.Row][i.Col])
		},
	)
	list.SetColumnWidth(0, 160)
	list.SetColumnWidth(1, 160)
	list.SetColumnWidth(2, 260)
	list.SetColumnWidth(3, 260)
	list.SetColumnWidth(4, 260)
	joblayout := container.New(layout.NewGridLayout(1), tklayout, list)
	return joblayout
}
