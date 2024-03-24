package api

//访客信息
import (
	"hotel/global"
	"time"
)

type LoginInfor struct {
	InfoId    int       `gorm:"info_id"`
	UserName  string    `gorm:"user_name"`
	Status    string    `gorm:"status"`
	Msg       string    `gorm:"msg"`
	LoginTime time.Time `gorm:"login_time"`
}

// Create 访客信息创建
func (r LoginInfor) Create(u HotelRole) {
	CreateUpdate(u, "login_infor")
	global.App.DB.Table("hotel_post").Create(r)
}
