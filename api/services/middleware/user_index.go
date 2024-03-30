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

func userCreate(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	RoleName := widget.NewLabel("RoleName : ")
	RoleNameempty := widget.NewEntry()
	RoleNameempty.SetPlaceHolder("RoleName")
	RoleNamelayout := container.New(layout.NewGridLayout(5), mySpace, RoleName, mySpace, RoleNameempty, mySpace)
	RoleNamelay := container.New(layout.NewGridLayout(1), mySpace, RoleNamelayout, mySpace)

	RolePasswd := widget.NewLabel("RolePasswd : ")
	RolePasswdempty := widget.NewEntry()
	RolePasswdempty.SetPlaceHolder("RolePasswd")
	RolePasswdlayout := container.New(layout.NewGridLayout(5), mySpace, RolePasswd, mySpace, RolePasswdempty, mySpace)
	RolePasswdlay := container.New(layout.NewGridLayout(1), mySpace, RolePasswdlayout, mySpace)

	RoleKey := widget.NewLabel("RolePasswd : ")
	RoleKeyempty := widget.NewEntry()
	RoleKeyempty.SetPlaceHolder("RolePasswd")
	RoleKeylayout := container.New(layout.NewGridLayout(5), mySpace, RoleKey, mySpace, RoleKeyempty, mySpace)
	RoleKeylay := container.New(layout.NewGridLayout(1), mySpace, RoleKeylayout, mySpace)

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

	remark := widget.NewLabel("Remark : ")
	remarkempty := widget.NewEntry()
	remarkempty.SetPlaceHolder("Remark")
	remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarkempty, mySpace)
	remarklay := container.New(layout.NewGridLayout(1), mySpace, remarklayout, mySpace)

	userCreate := widget.NewButton("Create", func() {
		user := api.HotelRole{}

		user.RoleName = RoleNameempty.Text
		user.Phone = phoneempty.Text
		user.Status = statusempty.Text
		user.Remark = remarkempty.Text
		user.DelFlag = "0"
		user.CreateTime = time.Now()
		user.CreateBy = u.RoleName
		user.RoleKey = RoleKeyempty.Text
		if user.RoleKey == "ceo" || user.RoleKey == "leader" {
			user.DataScope = "1"
		} else if user.RoleKey == "money" {
			user.DataScope = "2"
		} else if user.RoleKey == "user_996" {
			user.DataScope = "3"
		} else {
			user.DataScope = "4"
		}
		user.RolePasswd = RolePasswdempty.Text

		global.App.DB.Table("hotel_role").Create(&user)
	})

	userCreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, userCreate, mySpace, mySpace)

	userlayout := container.New(layout.NewGridLayout(1), RoleNamelay, RolePasswdlay, phonelay, statuslay, RoleKeylay, remarklay, mySpace, userCreatelayout, mySpace)

	return userlayout
}

func userDelete(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	RoleName := widget.NewLabel("RoleName : ")
	RoleNameempty := widget.NewEntry()
	RoleNameempty.SetPlaceHolder("RoleName")
	RoleNamelayout := container.New(layout.NewGridLayout(5), mySpace, RoleName, mySpace, RoleNameempty, mySpace)
	RoleNamelay := container.New(layout.NewGridLayout(1), mySpace, RoleNamelayout, mySpace)

	RolePasswd := widget.NewLabel("RolePasswd : ")
	RolePasswdempty := widget.NewEntry()
	RolePasswdempty.SetPlaceHolder("RolePasswd")
	RolePasswdlayout := container.New(layout.NewGridLayout(5), mySpace, RolePasswd, mySpace, RolePasswdempty, mySpace)
	RolePasswdlay := container.New(layout.NewGridLayout(1), mySpace, RolePasswdlayout, mySpace)

	RoleKey := widget.NewLabel("RoleKey : ")
	RoleKeyempty := widget.NewEntry()
	RoleKeyempty.SetPlaceHolder("RoleKey")
	RoleKeylayout := container.New(layout.NewGridLayout(5), mySpace, RoleKey, mySpace, RoleKeyempty, mySpace)
	RoleKeylay := container.New(layout.NewGridLayout(1), mySpace, RoleKeylayout, mySpace)

	phone := widget.NewLabel("phone : ")
	phoneempty := widget.NewEntry()
	phoneempty.SetPlaceHolder("phone")
	phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phoneempty, mySpace)
	phonelay := container.New(layout.NewGridLayout(1), mySpace, phonelayout, mySpace)

	userDelete := widget.NewButton("Delete", func() {
		user := api.HotelRole{}

		user.RoleName = RoleNameempty.Text
		user.Phone = phoneempty.Text
		user.RoleKey = RoleKeyempty.Text
		if user.RoleKey == "ceo" || user.RoleKey == "leader" {
			user.DataScope = "1"
		} else if user.RoleKey == "money" {
			user.DataScope = "2"
		} else if user.RoleKey == "user_996" {
			user.DataScope = "3"
		} else {
			user.DataScope = "4"
		}
		user.RolePasswd = RolePasswdempty.Text

		user.Status = "1"
		user.DelFlag = "2"
		user.UpdateTime = time.Now()
		user.UpdateBy = u.RoleName
		global.App.DB.Table("user").Update(&user)
	})

	userDeletelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, userDelete, mySpace, mySpace)

	userDeletelay := container.New(layout.NewGridLayout(1), RoleNamelay, RolePasswdlay, phonelay, RoleKeylay, mySpace, userDeletelayout, mySpace)

	return userDeletelay
}

func userUpdate(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	RoleName := widget.NewLabel("RoleName : ")
	RoleNameempty := widget.NewEntry()
	RoleNameempty.SetPlaceHolder("RoleName")
	RoleNamelayout := container.New(layout.NewGridLayout(5), mySpace, RoleName, mySpace, RoleNameempty, mySpace)
	RoleNamelay := container.New(layout.NewGridLayout(1), mySpace, RoleNamelayout, mySpace)

	RolePasswd := widget.NewLabel("RolePasswd : ")
	RolePasswdempty := widget.NewEntry()
	RolePasswdempty.SetPlaceHolder("RolePasswd")
	RolePasswdlayout := container.New(layout.NewGridLayout(5), mySpace, RolePasswd, mySpace, RolePasswdempty, mySpace)
	RolePasswdlay := container.New(layout.NewGridLayout(1), mySpace, RolePasswdlayout, mySpace)

	RoleKey := widget.NewLabel("RolePasswd : ")
	RoleKeyempty := widget.NewEntry()
	RoleKeyempty.SetPlaceHolder("RolePasswd")
	RoleKeylayout := container.New(layout.NewGridLayout(5), mySpace, RoleKey, mySpace, RoleKeyempty, mySpace)
	RoleKeylay := container.New(layout.NewGridLayout(1), mySpace, RoleKeylayout, mySpace)

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

	remark := widget.NewLabel("Remark : ")
	remarkempty := widget.NewEntry()
	remarkempty.SetPlaceHolder("Remark")
	remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarkempty, mySpace)
	remarklay := container.New(layout.NewGridLayout(1), mySpace, remarklayout, mySpace)

	userUpdate := widget.NewButton("Update", func() {
		user := api.HotelRole{}
		user.RolePasswd = RolePasswdempty.Text
		global.App.DB.Table("hotel_role").Create(&user)
		global.App.DB.Table("user").Where("role_name = ?", RoleNameempty.Text).Find(&user)
		if RoleNameempty.Text != "" {
			user.RoleName = RoleNameempty.Text
		}

		if phoneempty.Text != "" {
			user.Phone = phoneempty.Text
		}

		if statusempty.Text != "" {
			user.Status = statusempty.Text
		}

		if remarkempty.Text != "" {
			user.Remark = remarkempty.Text
		}
		if RoleKeyempty.Text != "" {
			user.RoleKey = RoleKeyempty.Text
		}
		if user.RoleKey == "ceo" || user.RoleKey == "leader" {
			user.DataScope = "1"
		} else if user.RoleKey == "money" {
			user.DataScope = "2"
		} else if user.RoleKey == "user_996" {
			user.DataScope = "3"
		} else {
			user.DataScope = "4"
		}
		user.UpdateTime = time.Now()
		user.UpdateBy = u.RoleName
		global.App.DB.Table("user").Update(&user)
	})

	userCreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, userUpdate, mySpace, mySpace)

	userlayout := container.New(layout.NewGridLayout(1), RoleNamelay, RolePasswdlay, phonelay, statuslay, RoleKeylay, remarklay, mySpace, userCreatelayout, mySpace)

	return userlayout
}

func userShow(u api.HotelRole) *fyne.Container {
	var mySpace = layout.NewSpacer()
	tks := make([]api.HotelRole, 100)
	tkss1 := make([][6]string, 100)
	num1 := 1

	tkss1[0][0] = "RoleName"
	tkss1[0][1] = "RolePasswd"
	tkss1[0][2] = "RoleKey"
	tkss1[0][3] = "Phone"
	tkss1[0][4] = "Remark"
	tkss1[0][5] = "Status"

	global.App.DB.Table("hotel_role").Where("del_flag = ?", "0").Find(&tks)

	for i := 0; i < len(tks); i++ {
		tkss1[num1][0] = tks[i].RoleName
		tkss1[num1][1] = tks[i].RolePasswd
		tkss1[num1][2] = tks[i].RoleKey
		tkss1[num1][3] = tks[i].Phone
		tkss1[num1][4] = tks[i].Remark
		tkss1[num1][5] = tks[i].Status
		num1++
	}

	tkpb := widget.NewLabel("Hotel user : ")
	tklay := container.New(layout.NewGridLayout(2), tkpb, mySpace)
	tklayout := container.New(layout.NewGridLayout(1), mySpace, tklay, mySpace)
	list1 := widget.NewTable(
		func() (int, int) {
			return num1 + 2, 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Hotel user : ")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tkss1[i.Row][i.Col])
		},
	)
	list1.SetColumnWidth(0, 160)
	list1.SetColumnWidth(1, 160)
	list1.SetColumnWidth(2, 260)
	list1.SetColumnWidth(3, 160)
	list1.SetColumnWidth(4, 160)
	list1.SetColumnWidth(4, 160)
	joblayout1 := container.New(layout.NewGridLayout(1), tklayout, list1)

	return joblayout1
}

// User user 员工增删改查
func User(u api.HotelRole) *fyne.Container {
	tab := container.NewAppTabs(
		container.NewTabItem("Create", container.New(layout.NewGridLayout(1), userCreate(u))),
		container.NewTabItem("Delete", container.New(layout.NewGridLayout(1), userDelete(u))),
		container.NewTabItem("Update", container.New(layout.NewGridLayout(1), userUpdate(u))),
		container.NewTabItem("Show", container.New(layout.NewGridLayout(1), userShow(u))),
	)
	tab.SetTabLocation(container.TabLocationTop)
	ans := container.New(layout.NewGridLayout(1), tab)
	return ans
}
