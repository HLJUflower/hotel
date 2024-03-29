package middleware

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"hotel/config/api"
	"hotel/global"
	"time"
)

// 客户登记

func sleepUserCreate(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	roleName := widget.NewLabel("role_name : ")
	roleNameempty := widget.NewEntry()
	roleNameempty.SetPlaceHolder("role_name")
	roleNamelayout := container.New(layout.NewGridLayout(5), mySpace, roleName, mySpace, roleNameempty, mySpace)
	roleNamelay := container.New(layout.NewGridLayout(1), mySpace, roleNamelayout, mySpace)

	phone := widget.NewLabel("phone : ")
	phoneempty := widget.NewEntry()
	phoneempty.SetPlaceHolder("phone")
	phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phoneempty, mySpace)
	phonelay := container.New(layout.NewGridLayout(1), mySpace, phonelayout, mySpace)

	status := widget.NewLabel("Status : ")
	statusempty := widget.NewEntry()
	statusempty.SetPlaceHolder("Status")
	statuslayout := container.New(layout.NewGridLayout(5), mySpace, status, mySpace, statusempty, mySpace)
	statuslay := container.New(layout.NewGridLayout(1), mySpace, statuslayout, mySpace)

	remark := widget.NewLabel("sleep_userID : ")
	remarkempty := widget.NewEntry()
	remarkempty.SetPlaceHolder("sleep_userID")
	remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarkempty, mySpace)
	remarklay := container.New(layout.NewGridLayout(1), mySpace, remarklayout, mySpace)

	sleepUsercreate := widget.NewButton("Create", func() {
		sleepUser := api.SleepUser{}
		sleepUser.RoleName = roleNameempty.Text
		sleepUser.Phone = phoneempty.Text
		sleepUser.Status = statusempty.Text
		sleepUser.Remark = remarkempty.Text
		sleepUser.DelFlag = "0"
		sleepUser.CreateTime = time.Now()
		sleepUser.CreateBy = u.RoleName
		global.App.DB.Table("sleep_user").Create(&sleepUser)
	})

	sleepUsercreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, sleepUsercreate, mySpace, mySpace)

	sleepUserlayout := container.New(layout.NewGridLayout(1), roleNamelay, phonelay, statuslay, remarklay, mySpace, sleepUsercreatelayout, mySpace)

	return sleepUserlayout
}

func sleepUserupdate(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	roleName := widget.NewLabel("role_name : ")
	roleNameempty := widget.NewEntry()
	roleNameempty.SetPlaceHolder("role_name")
	roleNamelayout := container.New(layout.NewGridLayout(5), mySpace, roleName, mySpace, roleNameempty, mySpace)
	roleNamelay := container.New(layout.NewGridLayout(1), mySpace, roleNamelayout, mySpace)

	phone := widget.NewLabel("phone : ")
	phoneempty := widget.NewEntry()
	phoneempty.SetPlaceHolder("phone")
	phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phoneempty, mySpace)
	phonelay := container.New(layout.NewGridLayout(1), mySpace, phonelayout, mySpace)

	status := widget.NewLabel("Status : ")
	statusempty := widget.NewEntry()
	statusempty.SetPlaceHolder("Status")
	statuslayout := container.New(layout.NewGridLayout(5), mySpace, status, mySpace, statusempty, mySpace)
	statuslay := container.New(layout.NewGridLayout(1), mySpace, statuslayout, mySpace)

	remark := widget.NewLabel("sleep_RoomID : ")
	remarkempty := widget.NewEntry()
	remarkempty.SetPlaceHolder("sleep_RoomID")
	remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarkempty, mySpace)
	remarklay := container.New(layout.NewGridLayout(1), mySpace, remarklayout, mySpace)

	sleepUserUpdate := widget.NewButton("Update", func() {
		sleepUser := api.SleepUser{}
		global.App.DB.Table("sleep_user").Where("role_name = ?", roleNameempty.Text).Find(&sleepUser)
		if phoneempty.Text != "" {

			sleepUser.Phone = phoneempty.Text
		}

		if statusempty.Text != "" {
			sleepUser.Status = statusempty.Text
		}

		if remarkempty.Text != "" {
			sleepUser.Remark = remarkempty.Text
		}

		sleepUser.UpdateTime = time.Now()
		sleepUser.UpdateBy = u.RoleName
		global.App.DB.Table("sleep_user").Update(&sleepUser)
	})

	sleepUsercreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, sleepUserUpdate, mySpace, mySpace)

	sleepUserlayout := container.New(layout.NewGridLayout(1), roleNamelay, phonelay, statuslay, remarklay, mySpace, sleepUsercreatelayout, mySpace)

	return sleepUserlayout
}

func sleepUsershow(u api.HotelRole) *fyne.Container {
	var mySpace = layout.NewSpacer()

	phone := widget.NewLabel("phone : ")
	phoneempty := widget.NewEntry()
	phoneempty.SetPlaceHolder("phone")
	phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phoneempty, mySpace)
	phonelay := container.New(layout.NewGridLayout(1), mySpace, phonelayout, mySpace)

	sleepUseShow := widget.NewButton("Show", func() {
		sleepUser := api.SleepUser{}
		global.App.DB.Table("sleep_user").Where("phone = ?", phoneempty.Text).Find(&sleepUser)
		w := global.A.NewWindow("Information")

		roleName := widget.NewLabel("role_name : ")
		roleNameempty := widget.NewLabel(sleepUser.RoleName)
		roleNamelayout := container.New(layout.NewGridLayout(5), mySpace, roleName, mySpace, roleNameempty, mySpace)
		roleNamelay := container.New(layout.NewGridLayout(1), mySpace, roleNamelayout, mySpace)

		phone := widget.NewLabel("phone : ")
		phoneempty := widget.NewLabel(sleepUser.Phone)
		phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phoneempty, mySpace)
		phonelay := container.New(layout.NewGridLayout(1), mySpace, phonelayout, mySpace)

		status := widget.NewLabel("Status : ")
		statusempty := widget.NewLabel(sleepUser.Status)
		statuslayout := container.New(layout.NewGridLayout(5), mySpace, status, mySpace, statusempty, mySpace)
		statuslay := container.New(layout.NewGridLayout(1), mySpace, statuslayout, mySpace)

		remark := widget.NewLabel("sleep_roomID : ")
		remarkempty := widget.NewLabel(sleepUser.Remark)
		remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarkempty, mySpace)
		remarklay := container.New(layout.NewGridLayout(1), mySpace, remarklayout, mySpace)

		content := container.New(layout.NewGridLayout(1), roleNamelay, phonelay, statuslay, remarklay)
		w.SetContent(content)
		w.Show()
	})

	sleepUsercreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, sleepUseShow, mySpace, mySpace)

	sleepUserlayout := container.New(layout.NewGridLayout(1), phonelay, mySpace, sleepUsercreatelayout, mySpace)

	return sleepUserlayout

}

func SleepUserLog(u api.HotelRole) *fyne.Container {
	tab := container.NewAppTabs(
		container.NewTabItem("Create", container.New(layout.NewGridLayout(1), sleepUserCreate(u))),
		container.NewTabItem("Update", container.New(layout.NewGridLayout(1), sleepUserupdate(u))),
		container.NewTabItem("Show", container.New(layout.NewGridLayout(1), sleepUsershow(u))),
	)
	tab.SetTabLocation(container.TabLocationTop)
	ans := container.New(layout.NewGridLayout(1), tab)
	return ans
}
