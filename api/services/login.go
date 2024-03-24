package services

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"hotel/config/api"
	"hotel/global"
	"image/color"
	"time"
)

// User 登录验证信息
type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func Select() {
	w := global.A.NewWindow("Hotel")
	wid, hig := 900, 600
	w.Resize(fyne.NewSize(float32(wid), float32(hig)))
	//enroll
	text := canvas.NewText("enroll", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 40
	textlayout := container.New(layout.NewGridLayout(3), mySpace, text, mySpace)

	username := widget.NewLabel("UserName")
	usernamelog := widget.NewEntry()
	usernamelog.SetPlaceHolder("UserName")
	usernamelayout := container.New(layout.NewGridLayout(5), mySpace, username, mySpace, usernamelog, mySpace)

	password := widget.NewLabel("Password")
	passwordlog := widget.NewPasswordEntry()
	passwordlog.SetPlaceHolder("Password")
	passwordlayout := container.New(layout.NewGridLayout(5), mySpace, password, mySpace, passwordlog, mySpace)

	passwordagain := widget.NewLabel("Password Again")
	passwordagainlog := widget.NewPasswordEntry()
	passwordagainlog.SetPlaceHolder("Password Again")
	passwordagainlayout := container.New(layout.NewGridLayout(5), mySpace, passwordagain, mySpace, passwordagainlog, mySpace)

	phone := widget.NewLabel("Phone/WeChat")
	phonelog := widget.NewEntry()
	phonelog.SetPlaceHolder("Contact details")
	phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phonelog, mySpace)

	rolekey := widget.NewLabel("Role")
	rolekeylog := widget.NewEntry()
	rolekeylog.SetPlaceHolder("Job title")
	rolekeylayout := container.New(layout.NewGridLayout(5), mySpace, rolekey, mySpace, rolekeylog, mySpace)

	remark := widget.NewLabel("Remark")
	remarklog := widget.NewEntry()
	remarklog.SetPlaceHolder("Remark")
	remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarklog, mySpace)
	zcbtn := widget.NewButton("enroll", func() {

	})
	zcbttn := container.New(layout.NewGridLayout(3), mySpace, zcbtn, mySpace)
	zcbtttn := container.NewVBox(zcbttn, mySpace)

	//login
	textlogin := canvas.NewText("Welcome to the Hotel Management System", color.White)
	textlogin.Alignment = fyne.TextAlignCenter
	textlogin.TextSize = 40

	userlogin := widget.NewEntry()
	passwordlogin := widget.NewPasswordEntry()
	userlogin.SetPlaceHolder("Account")
	passwordlogin.SetPlaceHolder("Password")
	textupSpacelogin := layout.NewSpacer()
	textdownSpacelogin := layout.NewSpacer()
	userlayoutlogin := container.New(layout.NewGridLayout(5), textupSpacelogin, widget.NewLabel("UserName"), textupSpacelogin, userlogin, textupSpacelogin)
	passwordlayoutlogin := container.New(layout.NewGridLayout(5), textupSpacelogin, widget.NewLabel("Password"), textupSpacelogin, passwordlogin, textupSpacelogin)

	loginBtnlogin := widget.NewButton("Login", func() {
		if userlogin.Text == "" || passwordlogin.Text == "" {
			dialog.ShowInformation("Login Error", "The username or password cannot be empty!", w)
		} else {
			realuserlogin := api.HotelRole{}
			global.App.DB.Table("hotel_role").Where("role_name = ?", userlogin.Text).Find(&realuserlogin)
			if realuserlogin.RolePasswd == password.Text {
				loginuserlogin := api.LoginInfor{
					UserName:  realuserlogin.RoleName,
					Status:    "0",
					Msg:       "",
					LoginTime: time.Now(),
				}
				global.App.DB.Table("login_infor").Create(&loginuserlogin)
				global.App.Logger.Infoln(
					"Login ",
					loginuserlogin.InfoId,
					" ",
					loginuserlogin.UserName,
					" ",
					loginuserlogin.LoginTime,
				)
				w.Close()
				Index(realuserlogin)
			} else {
				dialog.ShowInformation("Login Error", "Wrong username or password!", w)
				global.App.Logger.Infoln(userlogin.Text, " ", "login error")
			}
		}

	})

	titlelogin := container.New(layout.NewGridLayout(1), textupSpacelogin, textlogin, textdownSpacelogin)
	titlelogin.MinSize().Min(fyne.NewSize(float32(wid/8), float32(hig)))
	userpwdlogin := container.NewVBox(textupSpacelogin, userlayoutlogin, textupSpacelogin, passwordlayoutlogin, textupSpacelogin)
	loginBttonlogin := container.New(layout.NewGridLayout(3), textupSpacelogin, loginBtnlogin, textdownSpacelogin)
	loginBttnlogin := container.NewVBox(textupSpacelogin, loginBttonlogin, textdownSpacelogin)

	tab := container.NewAppTabs(
		container.NewTabItem("Login", container.New(layout.NewGridLayout(1), titlelogin, userpwdlogin, loginBttnlogin)),
		container.NewTabItem("Enroll", container.New(layout.NewGridLayout(1), mySpace, textlayout, mySpace, usernamelayout, mySpace, passwordlayout, mySpace, passwordagainlayout, mySpace, phonelayout, mySpace, rolekeylayout, mySpace, remarklayout, mySpace, zcbtttn)),
	)
	tab.SetTabLocation(container.TabLocationTop)
	w.SetContent(tab)
	w.Show()
	global.A.Run()
}

// Login 登录函数+
func Login() {

	//zcbtn := widget.NewButton("enroll", func() {
	//	//w.Close()
	//	middleware.CreateUser()
	//})
	//zcbttn := container.New(layout.NewGridLayout(3), textupSpace, zcbtn, textupSpace)
	//zcbtttn := container.NewVBox(zcbttn, textupSpace)
	//text.TextStyle = fyne.TextStyle{Italic: true}

}

//	fmt.Println("-------------------------------------")
//	fmt.Println("**********欢迎进入酒店管理系统**********")
//	fmt.Println("-------------------------------------")
//	fmt.Println("如果您还未注册，请输入（注册）；否则请您输入（登录）！")
//	s := ""
//SC:
//	fmt.Scanln(s)
//	if s == "注册" {
//		middleware.CreateUser()
//	} else if s == "登录" {

//			tag := false
//			for tag == false {
//				fmt.Println("请输入用户名")
//				fmt.Scanln(u.ID)
//				fmt.Println("请输入密码")
//				fmt.Scanln(u.Password)
//
//
//			}
//		} else {
//			fmt.Println("请准确输入信息！")
//			goto SC
//	//	}
//
//}
