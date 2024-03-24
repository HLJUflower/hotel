package middleware

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
