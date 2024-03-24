package services

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	middleware2 "hotel/api/services/middleware"
	"hotel/config/api"
	"hotel/global"
	"strconv"
)

var mySpace = layout.NewSpacer()

// 住户登记
func zhuhudengji(w fyne.Window, content *fyne.Container) *fyne.Container {
	zhbtn := widget.NewButton("UserLogin", func() {
		w.Close()
		middleware2.SleepUserLog()
	})
	zhbttn := container.New(layout.NewGridLayout(3), mySpace, zhbtn, mySpace)
	zhbtttn := container.NewVBox(zhbttn, mySpace)
	content.Add(zhbtttn)
	return content

}

// 房间查询
func fangjianchaxun(w fyne.Window, content *fyne.Container) *fyne.Container {
	fjbtn := widget.NewButton("Room", func() {
		w.Close()
		middleware2.RoomInsert()
	})
	fjbttn := container.New(layout.NewGridLayout(3), mySpace, fjbtn, mySpace)
	fjbtttn := container.NewVBox(fjbttn, mySpace)
	content.Add(fjbtttn)
	return content

}

// 季度计划
func jidujihua(w fyne.Window, content *fyne.Container) *fyne.Container {
	jdbtn := widget.NewButton("Plan", func() {
		w.Close()
		middleware2.JobPlan()
	})
	jdbttn := container.New(layout.NewGridLayout(3), mySpace, jdbtn, mySpace)
	jdbtttn := container.NewVBox(jdbttn, mySpace)
	content.Add(jdbtttn)
	return content

}

// 员工查询
func yuangongchaxun(w fyne.Window, content *fyne.Container) *fyne.Container {
	ygbtn := widget.NewButton("Employee", func() {
		w.Close()
		middleware2.UserManager()
	})
	ygbttn := container.New(layout.NewGridLayout(3), mySpace, ygbtn, mySpace)
	ygbtttn := container.NewVBox(ygbttn, mySpace)
	content.Add(ygbtttn)
	return content

}

// 财务汇总
func caiwuhuizong(w fyne.Window, content *fyne.Container) *fyne.Container {
	cwbtn := widget.NewButton("Money", func() {
		w.Close()
		middleware2.MoneyTotal()
	})
	cwbttn := container.New(layout.NewGridLayout(3), mySpace, cwbtn, mySpace)
	cwbtttn := container.NewVBox(cwbttn, mySpace)
	content.Add(cwbtttn)
	return content

}

func Index(u api.HotelRole) {
	w := global.A.NewWindow("index")
	wid, hig := 900, 600
	w.Resize(fyne.NewSize(float32(wid), float32(hig)))
	content := container.New(layout.NewGridLayout(1))
	content.Add(mySpace)
	vip, _ := strconv.Atoi(u.DataScope)
	if vip == 4 {
		zhuhudengji(w, content)
		fangjianchaxun(w, content)
	} else if vip == 3 {
		yuangongchaxun(w, content)
		jidujihua(w, content)
	} else if vip == 2 {
		yuangongchaxun(w, content)
		jidujihua(w, content)
		caiwuhuizong(w, content)
	} else if vip == 1 {
		zhuhudengji(w, content)
		fangjianchaxun(w, content)
		yuangongchaxun(w, content)
		jidujihua(w, content)
		caiwuhuizong(w, content)
	}
	lobtn := widget.NewButton("LogOut", func() {
		w.Close()
		Login()
	})
	lobttn := container.New(layout.NewGridLayout(3), mySpace, lobtn, mySpace)
	lobtttn := container.NewVBox(lobttn, mySpace)
	content.Add(lobtttn)
	w.SetContent(content)
	w.Show()
}
