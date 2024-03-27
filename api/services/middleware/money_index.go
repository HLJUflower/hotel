package middleware

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"hotel/config/api"
	"hotel/global"
	"hotel/utils/logging"
	"image/color"
	"strconv"
	"time"
)

// MoneyTotal 财务管理
func MoneyTotal(u api.HotelRole) *fyne.Container {
	var mySpace = layout.NewSpacer()
	can, err := strconv.Atoi(u.DataScope)
	if err != nil {
		logging.Errorln("Money Power Error : ", err)
	}
	if can <= 2 {
		sum, num := 0, 0
		menu := make([]api.Menu, 100)
		menus := make([][5]string, 100)
		now := time.Now()
		menus[0][0] = "ServeName"
		menus[0][1] = "Money"
		menus[0][2] = "CreateTime"
		menus[0][3] = "Room"
		menus[0][4] = "Remark"
		for i := 1; time.Now().AddDate(0, -1, 0).Before(now); i++ {
			global.App.DB.Table("menu").Where("create_time < ?", now).Last(&menu[i])
			menus[i][0] = menu[i].ServeName
			menus[i][1] = menu[i].Money
			menus[i][2] = menu[i].CreateTime.Format("2006-01-02 15:04:05")
			menus[i][3] = menu[i].Room
			menus[i][4] = menu[i].Remark
			mon, _ := strconv.Atoi(menus[i][1])
			sum += mon
			now = menu[i].CreateTime
			num++
		}

		list := widget.NewTable(
			func() (int, int) {
				return num + 2, 5
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("Past Plans")
			},
			func(i widget.TableCellID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(menus[i.Row][i.Col])
			},
		)
		list.SetColumnWidth(0, 360)
		list.SetColumnWidth(1, 160)
		list.SetColumnWidth(2, 260)
		list.SetColumnWidth(3, 100)
		list.SetColumnWidth(4, 100)
		moneylayout := container.New(layout.NewGridLayout(1), list)
		return moneylayout
	} else {
		always := canvas.NewText("Sorry, You Don't Have Enough Permissions", color.White)
		always.TextSize = 60
		alway := container.New(layout.NewGridLayout(1), mySpace, always, mySpace)
		return alway
	}
}
