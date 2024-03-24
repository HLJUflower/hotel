package services

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"hotel/config/api"
	"hotel/global"
	"image/color"
	"strconv"
	"time"
)

// User 登录验证信息
type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func Select() {
	//新建登录页
	w := global.A.NewWindow("Hotel")
	wid, hig := 900, 600
	w.Resize(fyne.NewSize(float32(wid), float32(hig)))
	//enroll 注册页面组件
	//标题
	text := canvas.NewText("enroll", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 40
	textlayout := container.New(layout.NewGridLayout(3), mySpace, text, mySpace)
	//用户名
	username := widget.NewLabel("UserName")
	usernamelog := widget.NewEntry()
	usernamelog.SetPlaceHolder("UserName")
	usernamelayout := container.New(layout.NewGridLayout(5), mySpace, username, mySpace, usernamelog, mySpace)
	//密码
	password := widget.NewLabel("Password")
	passwordlog := widget.NewPasswordEntry()
	passwordlog.SetPlaceHolder("Password")
	passwordlayout := container.New(layout.NewGridLayout(5), mySpace, password, mySpace, passwordlog, mySpace)
	//确认密码
	passwordagain := widget.NewLabel("Password Again")
	passwordagainlog := widget.NewPasswordEntry()
	passwordagainlog.SetPlaceHolder("Password Again")
	passwordagainlayout := container.New(layout.NewGridLayout(5), mySpace, passwordagain, mySpace, passwordagainlog, mySpace)
	//联系方式
	phone := widget.NewLabel("Phone/WeChat")
	phonelog := widget.NewEntry()
	phonelog.SetPlaceHolder("Contact details")
	phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phonelog, mySpace)
	//职务
	rolekey := widget.NewLabel("Role")
	rolekeylog := widget.NewEntry()
	rolekeylog.SetPlaceHolder("Job title")
	rolekeylayout := container.New(layout.NewGridLayout(5), mySpace, rolekey, mySpace, rolekeylog, mySpace)
	//备注
	remark := widget.NewLabel("Remark")
	remarklog := widget.NewEntry()
	remarklog.SetPlaceHolder("Remark")
	remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarklog, mySpace)
	zcbtn := widget.NewButton("enroll", func() {
		user := api.HotelRole{}
		if passwordagainlog.Text == passwordlog.Text {
			user.RolePasswd = passwordlog.Text
			user.RoleName = usernamelog.Text
			user.Phone = phonelog.Text
			user.RoleKey = rolekeylog.Text
			user.Remark = remarklog.Text
			post := api.HotelPost{}
			global.App.DB.Table("hotel_post").Where("post_code = ?", rolekeylog.Text).Find(&post)
			user.DataScope = strconv.Itoa(post.PostSort)
			fmt.Println(post.PostSort)
			user.DelFlag = "0"
			user.Status = "1"
			user.CreateTime = time.Now()
			user.UpdateTime = time.Now()
			global.App.DB.Table("hotel_role").Create(user)
			dialog.ShowInformation("Enroll", "Success!", w)
		} else {
			dialog.ShowInformation("Enroll", "Password Wrong!", w)
		}

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
			if realuserlogin.RolePasswd == passwordlogin.Text {
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
		container.NewTabItem("Enroll", container.New(layout.NewGridLayout(1), mySpace, textlayout, mySpace, usernamelayout, mySpace, passwordlayout, mySpace, passwordagainlayout, mySpace, phonelayout, mySpace, rolekeylayout, mySpace, remarklayout, mySpace, zcbtttn, mySpace)),
	)
	tab.SetTabLocation(container.TabLocationTop)
	w.SetContent(tab)
	w.Show()
	global.A.Run()
}

// Login 登录函数+
//func Login() {

//zcbtn := widget.NewButton("enroll", func() {
//	//w.Close()
//	middleware.CreateUser()
//})
//zcbttn := container.New(layout.NewGridLayout(3), textupSpace, zcbtn, textupSpace)
//zcbtttn := container.NewVBox(zcbttn, textupSpace)
//text.TextStyle = fyne.TextStyle{Italic: true}

//}

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

//
//import (
//	"fmt"
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/layout"
//	"fyne.io/fyne/v2/widget"
//	"hotel/global"
//	"image/color"
//)
//
//var mySpace = layout.NewSpacer()
//
//func myinit (content *fyne.Container){
//	text := canvas.NewText("注册", color.White)
//	text.Alignment = fyne.TextAlignCenter
//	text.TextSize = 40
//	textlayout := container.New(layout.NewGridLayout(3), mySpace, text, mySpace)
//
//	username := widget.NewLabel("UserName")
//	usernamelog := widget.NewEntry()
//	usernamelog.SetPlaceHolder("Please enter your username")
//	usernamelayout := container.New(layout.NewGridLayout(5), mySpace, username, mySpace, usernamelog, mySpace)
//
//	password := widget.NewLabel("Password")
//	passwordlog := widget.NewPasswordEntry()
//	passwordlog.SetPlaceHolder("Please enter your password")
//	passwordlayout := container.New(layout.NewGridLayout(5), mySpace, password, mySpace, passwordlog, mySpace)
//
//	passwordagain := widget.NewLabel("Password Again")
//	passwordagainlog := widget.NewPasswordEntry()
//	passwordagainlog.SetPlaceHolder("Please enter your password again")
//	passwordagainlayout := container.New(layout.NewGridLayout(5), mySpace, passwordagain, mySpace, passwordagainlog, mySpace)
//
//	phone := widget.NewLabel("Phone/WeChat")
//	phonelog := widget.NewEntry()
//	phonelog.SetPlaceHolder("Please enter your contact details")
//	phonelayout := container.New(layout.NewGridLayout(5), mySpace, phone, mySpace, phonelog, mySpace)
//
//	rolekey := widget.NewLabel("Role")
//	rolekeylog := widget.NewEntry()
//	rolekeylog.SetPlaceHolder("Please enter your job title")
//	rolekeylayout := container.New(layout.NewGridLayout(5), mySpace, rolekey, mySpace, rolekeylog, mySpace)
//
//	remark := widget.NewLabel("Remark")
//	remarklog := widget.NewEntry()
//	remarklog.SetPlaceHolder("Please enter your remarks")
//	remarklayout := container.New(layout.NewGridLayout(5), mySpace, remark, mySpace, remarklog, mySpace)
//	zcbtn := widget.NewButton("enroll", func() {
//
//	})
//	zcbttn := container.New(layout.NewGridLayout(3), mySpace, zcbtn, mySpace)
//	zcbtttn := container.NewVBox(zcbttn, mySpace)
//
//	content.Add(layout.NewGridLayout(1), mySpace, textlayout, mySpace, usernamelayout, mySpace, passwordlayout, mySpace, passwordagainlayout, mySpace, phonelayout, mySpace, rolekeylayout, mySpace, remarklayout, mySpace, zcbtttn)
//}
//
//func CreateUser() {
//	fmt.Println(2)
//	w := global.A.NewWindow("create")
//	wid, hig := 900, 600
//	w.Resize(fyne.NewSize(float32(wid), float32(hig)))
//	//newuser := api.HotelRole{}
//	content := container.New()
//
//	w.SetContent(content)
//	w.Show()
//	//fmt.Println("请输入您的昵称：")
//	//fmt.Scanln(newuser.RoleName)
//	//firpwd, secpwd := "", ""
//	//same := false
//	//for same == false {
//	//	fmt.Println("请输入您的密码：")
//	//	fmt.Scanln(firpwd)
//	//	fmt.Println("请您再次输入密码，以保证密码输入正确：")
//	//	fmt.Scanln(secpwd)
//	//	if firpwd == secpwd {
//	//		same = true
//	//		newuser.RolePasswd = firpwd
//	//	} else {
//	//		fmt.Scanln("两次密码不一致！请重新输入！")
//	//	}
//	//}
//	//fmt.Println("请输入您的电话号/微信号：")
//	//fmt.Scanln(newuser.Phone)
//	//newuser.CreateBy = newuser.RoleName
//	//
//	//fmt.Println("请输入您的职位：")
//	//fmt.Scanln(newuser.RoleKey)
//	//adduser := api.HotelPost{}
//	//global.App.DB.Table("hotel_post").Where("role_key = ?", newuser.RoleKey).Find(&adduser)
//	//if adduser.PostCode != newuser.RoleKey {
//	//	newuser.RoleKey = "user_007"
//	//}
//	//global.App.DB.Table("hotel_post").Where("post_code = ?", newuser.DataScope).Find(&adduser)
//	//newuser.DataScope = strconv.Itoa(adduser.PostSort)
//	//newuser.RoleSort = adduser.PostId
//	//newuser.CreateBy = newuser.RoleName
//	//newuser.Status = "1"
//	//newuser.CreateTime = time.Now()
//	//fmt.Println("请输入备注信息：")
//	//fmt.Scanln(newuser.Remark)
//	//global.App.DB.Table("hotel_role").Create(&newuser)
//	//return
//}
