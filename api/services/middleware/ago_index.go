package middleware

//
//import (
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/layout"
//	"fyne.io/fyne/v2/widget"
//	"hotel/config/api"
//	"hotel/global"
//	"time"
//)
//
//func agoShow() *fyne.Container {
//	tks := make([]api.LoginInfor, 50)
//	tkss1 := make([][4]string, 50)
//	num1 := 1
//
//	tkss1[0][0] = "user_name"
//	tkss1[0][1] = "status"
//	tkss1[0][2] = "msg"
//	tkss1[0][3] = "login_time"
//
//	now := time.Now()
//
//	for i := 0; i < 20; i++ {
//		global.App.DB.Table("login_infor").Where("login_time < ?", now).Last(&tks)
//		tkss1[num1][0] = tks[i].UserName
//		tkss1[num1][1] = tks[i].Status
//		tkss1[num1][2] = tks[i].Msg
//		tkss1[num1][3] = tks[i].LoginTime.Format("2006-01-02 15:04:05")
//		now = tks[i].LoginTime
//		num1++
//	}
//
//	list1 := widget.NewTable(
//		func() (int, int) {
//			return num1 + 2, 5
//		},
//		func() fyne.CanvasObject {
//			return widget.NewLabel("Login user : ")
//		},
//		func(i widget.TableCellID, o fyne.CanvasObject) {
//			o.(*widget.Label).SetText(tkss1[i.Row][i.Col])
//		},
//	)
//	list1.SetColumnWidth(0, 260)
//	list1.SetColumnWidth(1, 260)
//	list1.SetColumnWidth(2, 260)
//	list1.SetColumnWidth(3, 260)
//	joblayout1 := container.New(layout.NewGridLayout(1), list1)
//
//	return joblayout1
//}
//
//// Ago 历史记录查询
//func Ago() *fyne.Container {
//	tab := container.NewAppTabs(
//		container.NewTabItem("Show", container.New(layout.NewGridLayout(1), agoShow())),
//	)
//	tab.SetTabLocation(container.TabLocationTop)
//	ans := container.New(layout.NewGridLayout(1), tab)
//	return ans
//}
