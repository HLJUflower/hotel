package middleware

import (
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
		if roomIDempty.Text != "" {
			roomids, _ := strconv.Atoi(roomIDempty.Text)
			room.RoomId = int64(roomids)
		}
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

func roomShow(u api.HotelRole) {

}

// Room RoomInsert 房间增删改查
func Room(u api.HotelRole) *fyne.Container {
	tab := container.NewAppTabs(
		container.NewTabItem("Show", container.New(layout.NewGridLayout(1), roomCreate(u))),
		container.NewTabItem("Create", container.New(layout.NewGridLayout(1), roomDelete(u))),
		container.NewTabItem("Update", container.New(layout.NewGridLayout(1), roomUpdate(u))),
	)
	tab.SetTabLocation(container.TabLocationTop)
	ans := container.New(layout.NewGridLayout(1), tab)
	return ans
}
