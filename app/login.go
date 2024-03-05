package app

import (
	"github.com/gin-gonic/gin"
	"hotel/utils/logging"
)

// User 登录验证信息
type User struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

// Login 登录函数
func Login() {
	r := gin.Default()
	//加载登录页面
	r.LoadHTMLFiles("index.html")
	logging.Infoln("log init success!")
	// 反馈基本信息
	r.GET("/index", func(c *gin.Context) {
		c.JSON(
			200,
			"Login Get",
		)
	})
	//接收表单
	r.POST("/index", func(c *gin.Context) {
		c.JSON(
			200,
			"Login Post",
		)
	})
	//接收注册信息
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(
			200,
			"Login Put",
		)
	})
	//接收删除信息
	r.DELETE("/index", func(c *gin.Context) {
		//u := User{
		//	ID:   123,
		//	Name: "Hello",
		//	Age:  20,
		//}
		c.JSON(
			200,
			"Login Delete",
		)
	})
	//3.监听端口，默认8888
	err2 := r.Run(":8888")
	if err2 != nil {
		return
	}
}
