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

func jobtext01() *fyne.Container {
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
			return widget.NewLabel("Past Plans : ")
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

func jobtext02(u api.HotelRole) *fyne.Container {
	var mySpace = layout.NewSpacer()

	jobName := widget.NewLabel("JobName : ")
	jobNameempty := widget.NewEntry()
	jobNameempty.SetPlaceHolder("JobName")
	jobNamelayout := container.New(layout.NewGridLayout(5), mySpace, jobName, mySpace, jobNameempty, mySpace)
	jobNamelay := container.New(layout.NewGridLayout(1), mySpace, jobNamelayout, mySpace)

	jobGroup := widget.NewLabel("JobGroup : ")
	jobGroupempty := widget.NewEntry()
	jobGroupempty.SetPlaceHolder("JobGroup")
	jobGrouplayout := container.New(layout.NewGridLayout(5), mySpace, jobGroup, mySpace, jobGroupempty, mySpace)
	jobGrouplay := container.New(layout.NewGridLayout(1), mySpace, jobGrouplayout, mySpace)

	jobDetail := widget.NewLabel("JobDetail : ")
	jobDetailempty := widget.NewEntry()
	jobDetailempty.SetPlaceHolder("JobDetail")
	jobDetaillayout := container.New(layout.NewGridLayout(5), mySpace, jobDetail, mySpace, jobDetailempty, mySpace)
	jobDetaillay := container.New(layout.NewGridLayout(1), mySpace, jobDetaillayout, mySpace)

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

	jobCreate := widget.NewButton("Create", func() {
		job := api.Task{}
		job.JobName = jobNameempty.Text
		job.JobDetail = jobDetailempty.Text
		job.JobGroup = jobGroupempty.Text
		job.Status = statusempty.Text
		job.Remark = remarkempty.Text
		job.CreateTime = time.Now()
		job.CreateBy = u.RoleName
		global.App.DB.Table("task").Create(&job)
	})

	jobCreatelayout := container.New(layout.NewGridLayout(5), mySpace, mySpace, jobCreate, mySpace, mySpace)

	joblayout := container.New(layout.NewGridLayout(1), jobNamelay, jobGrouplay, jobDetaillay, statuslay, remarklay, mySpace, jobCreatelayout, mySpace)

	return joblayout
}

// JobPlan 季度计划
func JobPlan(u api.HotelRole) *fyne.Container {

	tab := container.NewAppTabs(
		container.NewTabItem("Show", container.New(layout.NewGridLayout(1), jobtext01())),
		container.NewTabItem("Create", container.New(layout.NewGridLayout(1), jobtext02(u))),
	)
	tab.SetTabLocation(container.TabLocationTop)
	ans := container.New(layout.NewGridLayout(1), tab)
	return ans
}
