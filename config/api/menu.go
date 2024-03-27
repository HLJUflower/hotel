package api

//账单信息
import (
	"hotel/global"
	"time"
)

type Menu struct {
	menuId     int       `json:"menu_id" gorm:"menu_id"`
	ServeName  string    `json:"serve_name" gorm:"serve_name"`
	Room       string    `json:"room" gorm:"room"`
	Money      string    `json:"money" gorm:"money"`
	status     string    `json:"status" gorm:"status"`
	CreateTime time.Time `json:"createTime" gorm:"create_time"`
	CreateBy   string    `json:"createBy" gorm:"create_by"`
	UpdateTime time.Time `json:"updateTime" gorm:"update_time"`
	UpdateBy   string    `json:"updateBy" gorm:"update_by"`
	Remark     string    `json:"remark" gorm:"remark"`
}

// Create 账单信息创建
func (r Menu) Create(u HotelRole) {
	CreateUpdate(u, "menu")
	global.App.DB.Table("menu").Create(r)
}

// Delete 账单信息删除
func (r Menu) Delete(u HotelRole) {
	CreateUpdate(u, "menu")
	global.App.DB.Table("menu").Delete(r)
}

// Update 账单信息更新
func (r Menu) Update(u HotelRole) {
	CreateUpdate(u, "menu")
	global.App.DB.Table("menu").Save(r)
}
