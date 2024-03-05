package app

import (
	"github.com/gin-gonic/gin"
	"hotel/global"
	"hotel/utils/logging"
)

// User 登录验证信息
type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// Login 登录函数
func Login() {
	r := gin.Default()
	//加载登录页面
	r.LoadHTMLFiles("../resource/html/index.html")

	logging.Infoln("log init success!")
	// 反馈基本信息
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200,
			"../resource/html/index.html",
			gin.H{
				"title": "index",
			},
		)
	})
	//接收表单
	r.POST("/index", func(c *gin.Context) {
		//接收输入
		username := c.PostForm("username")
		password := c.PostForm("password")
		u := User{}
		//sql查询
		global.App.DB.Where(User{ID: username}).Find(u)
		if u.Password == password {
			//这里应该跳转到主页
			c.JSON(200,
				"Login Success",
			)
		} else {
			//提示登陆失败
			c.JSON(200,
				"Login Faild",
			)
		}
	})
	//3.监听端口，默认8888
	err2 := r.Run(":8888")
	if err2 != nil {
		return
	}
}
