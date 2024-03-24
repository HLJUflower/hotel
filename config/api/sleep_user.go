package api

//客户信息
import (
	"hotel/global"
	"time"
)

type SleepUser struct {
	RoleId     string    `gorm:"role_id" json:"role_id"`
	RoleName   string    `gorm:"role_name" json:"role_name"`
	Phone      string    `gorm:"phone" json:"phone"`
	Status     string    `gorm:"status" json:"status"`
	DelFlag    string    `gorm:"del_flag" json:"del_flag"`
	CreateBy   string    `gorm:"create_by" json:"create_by"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	UpdateBy   string    `gorm:"update_by" json:"update_by"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
	Remark     string    `gorm:"remark" jso:"remark"`
}

// 客户信息更新
func (r SleepUser) Update(u HotelRole) {
	CreateUpdate(u, "sleep_user")
	global.App.DB.Table("sleep_user").Save(r)
}

// 客户信息创建
func (r SleepUser) Create(u HotelRole) {
	CreateUpdate(u, "sleep_user")
	global.App.DB.Table("sleep_user").Save(r)
}
