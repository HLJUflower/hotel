package middleware

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"hotel/config/api"
	"hotel/global"
	"strconv"
	"time"
)

func roomCreate(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	roomID := widget.NewLabel("roomID : ")
	roomIDempty := widget.NewEntry()
	roomIDempty.SetPlaceHolder("roomID")
	roomIDlayout := container.New(layout.NewGridLayout(5), mySpace, roomID, mySpace, roomIDempty, mySpace)
	roomIDlay := container.New(layout.NewGridLayout(1), mySpace, roomIDlayout, mySpace)

	leader := widget.NewLabel("leader : ")
	leaderempty := widget.NewEntry()
	leaderempty.SetPlaceHolder("leader")
	leaderlayout := container.New(layout.NewGridLayout(5), mySpace, leader, mySpace, leaderempty, mySpace)
	leaderlay := container.New(layout.NewGridLayout(1), mySpace, leaderlayout, mySpace)

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

	roomCreate := widget.NewButton("Create", func() {
		room := api.Rooms{}
		roomids, _ := strconv.Atoi(roomIDempty.Text)
		room.RoomId = int64(roomids)
		room.Leader = leaderempty.Text
		room.Phone = phoneempty.Text
		room.Status = statusempty.Text
		room.Remark = remarkempty.Text
		room.DelFlag = "0"
		room.CreateTime = time.Now()
		room.CreateBy = u.RoleName
		global.App.DB.Table("room").Create(&room)
	})

	roomCreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, roomCreate, mySpace, mySpace)

	roomlayout := container.New(layout.NewGridLayout(1), roomIDlay, leaderlay, phonelay, statuslay, remarklay, mySpace, roomCreatelayout, mySpace)

	return roomlayout
}

func roomDelete(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	delRoom := widget.NewLabel("Del_RoomID : ")
	delRoomempty := widget.NewEntry()
	delRoomempty.SetPlaceHolder("Del_RoomID")
	delRoomlayout := container.New(layout.NewGridLayout(5), mySpace, delRoom, mySpace, delRoomempty, mySpace)
	delRoomlay := container.New(layout.NewGridLayout(1), mySpace, delRoomlayout, mySpace)

	userName := widget.NewLabel("userName : ")
	userNameempty := widget.NewEntry()
	userNameempty.SetPlaceHolder("userName")
	userNamelayout := container.New(layout.NewGridLayout(5), mySpace, userName, mySpace, userNameempty, mySpace)
	userNamelay := container.New(layout.NewGridLayout(1), mySpace, userNamelayout, mySpace)

	userPassword := widget.NewLabel("userPassword : ")
	userPasswordempty := widget.NewEntry()
	userPasswordempty.SetPlaceHolder("userPassword")
	userPasswordlayout := container.New(layout.NewGridLayout(5), mySpace, userPassword, mySpace, userPasswordempty, mySpace)
	userPasswordlay := container.New(layout.NewGridLayout(1), mySpace, userPasswordlayout, mySpace)

	roomDelete := widget.NewButton("Delete", func() {
		room := api.Rooms{}
		global.App.DB.Table("room").Where("room_id = ?", delRoomempty.Text).Find(&room)
		room.Status = "1"
		room.DelFlag = "2"
		room.UpdateTime = time.Now()
		room.UpdateBy = u.RoleName
		global.App.DB.Table("room").Update(&room)
	})

	roomDeletelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, roomDelete, mySpace, mySpace)

	roomDeletelay := container.New(layout.NewGridLayout(1), delRoomlay, userNamelay, userPasswordlay, mySpace, roomDeletelayout, mySpace)

	return roomDeletelay
}

func roomUpdate(u api.HotelRole) *fyne.Container {
	mySpace := layout.NewSpacer()

	roomID := widget.NewLabel("roomID : ")
	roomIDempty := widget.NewEntry()
	roomIDempty.SetPlaceHolder("roomID")
	roomIDlayout := container.New(layout.NewGridLayout(5), mySpace, roomID, mySpace, roomIDempty, mySpace)
	roomIDlay := container.New(layout.NewGridLayout(1), mySpace, roomIDlayout, mySpace)

	leader := widget.NewLabel("leader : ")
	leaderempty := widget.NewEntry()
	leaderempty.SetPlaceHolder("leader")
	leaderlayout := container.New(layout.NewGridLayout(5), mySpace, leader, mySpace, leaderempty, mySpace)
	leaderlay := container.New(layout.NewGridLayout(1), mySpace, leaderlayout, mySpace)

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

	roomUpdate := widget.NewButton("Update", func() {
		room := api.Rooms{}
		global.App.DB.Table("room").Where("room_id = ?", roomID.Text).Find(&room)
		if leaderempty.Text != "" {
			room.Leader = leaderempty.Text
		}

		if phoneempty.Text != "" {
			room.Phone = phoneempty.Text
		}

		if statusempty.Text != "" {
			room.Status = statusempty.Text
		}

		if remarkempty.Text != "" {
			room.Remark = remarkempty.Text
		}

		room.UpdateTime = time.Now()
		room.UpdateBy = u.RoleName
		global.App.DB.Table("room").Update(&room)
	})

	roomCreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, roomUpdate, mySpace, mySpace)

	roomlayout := container.New(layout.NewGridLayout(1), roomIDlay, leaderlay, phonelay, statuslay, remarklay, mySpace, roomCreatelayout, mySpace)

	return roomlayout
}

func roomShow(u api.HotelRole) *fyne.Container {
	var mySpace = layout.NewSpacer()
	tks := make([]api.Rooms, 100)
	tkss1, tkss2 := make([][5]string, 100), make([][5]string, 100)
	num1, num2 := 1, 1

	tkss1[0][0] = "RoomId"
	tkss1[0][1] = "Leader"
	tkss1[0][2] = "Phone"
	tkss1[0][3] = "Remark"
	tkss1[0][4] = "Status"

	tkss2[0][0] = "RoomId"
	tkss2[0][1] = "Leader"
	tkss2[0][2] = "Phone"
	tkss2[0][3] = "Remark"
	tkss2[0][4] = "Status"

	global.App.DB.Table("rooms").Where("del_flag = ?", "0").Find(&tks)

	for i := 0; i < len(tks); i++ {
		if tks[num1].Status == "0" && num1 <= 5 {
			tkss1[num1][0] = strconv.FormatInt(tks[i].RoomId, 10)
			tkss1[num1][1] = tks[i].Leader
			tkss1[num1][2] = tks[i].Phone
			tkss1[num1][3] = tks[i].Remark
			tkss1[num1][4] = tks[i].Status
			fmt.Println(tkss1[num1])
			num1++
		} else if tks[num1].Status == "1" && num2 <= 5 {
			tkss2[num2][0] = strconv.FormatInt(tks[i].RoomId, 10)
			tkss2[num2][1] = tks[i].Leader
			tkss2[num2][2] = tks[i].Phone
			tkss2[num2][3] = tks[i].Remark
			tkss2[num2][4] = tks[i].Status
			num2++
		}
	}

	tkpb := widget.NewLabel("available room : ")
	tklay := container.New(layout.NewGridLayout(2), tkpb, mySpace)
	tklayout := container.New(layout.NewGridLayout(1), mySpace, tklay, mySpace)
	list1 := widget.NewTable(
		func() (int, int) {
			return num1 + 2, 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("available room : ")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tkss1[i.Row][i.Col])
		},
	)
	list1.SetColumnWidth(0, 160)
	list1.SetColumnWidth(1, 160)
	list1.SetColumnWidth(2, 260)
	list1.SetColumnWidth(3, 260)
	list1.SetColumnWidth(4, 260)
	joblayout1 := container.New(layout.NewGridLayout(1), tklayout, list1)

	tkpb2 := widget.NewLabel("fixing room : ")
	tklay2 := container.New(layout.NewGridLayout(2), tkpb2, mySpace)
	tklayout2 := container.New(layout.NewGridLayout(1), mySpace, tklay2, mySpace)
	list2 := widget.NewTable(
		func() (int, int) {
			return num2 + 2, 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("fixing room : ")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tkss2[i.Row][i.Col])
		},
	)
	list1.SetColumnWidth(0, 160)
	list1.SetColumnWidth(1, 160)
	list1.SetColumnWidth(2, 260)
	list1.SetColumnWidth(3, 260)
	list1.SetColumnWidth(4, 260)
	joblayout2 := container.New(layout.NewGridLayout(1), tklayout2, list2)

	joblayout := container.New(layout.NewGridLayout(1), joblayout1, joblayout2)

	return joblayout
}

// Room 房间增删改查
func Room(u api.HotelRole) *fyne.Container {
	tab := container.NewAppTabs(
		container.NewTabItem("Create", container.New(layout.NewGridLayout(1), roomCreate(u))),
		container.NewTabItem("Delete", container.New(layout.NewGridLayout(1), roomDelete(u))),
		container.NewTabItem("Update", container.New(layout.NewGridLayout(1), roomUpdate(u))),
		container.NewTabItem("Show", container.New(layout.NewGridLayout(1), roomShow(u))),
	)
	tab.SetTabLocation(container.TabLocationTop)
	ans := container.New(layout.NewGridLayout(1), tab)
	return ans
}
