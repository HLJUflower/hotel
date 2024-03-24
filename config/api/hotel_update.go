package api

//更新信息
import (
	"hotel/global"
	"time"
)

type Hotel_Update struct {
	UpdateId   int       `json:"update_id" gorm:"update_id"`
	ServeName  string    `json:"serve_name" gorm:"serve_name"`
	Status     string    `gorm:"status" json:"status"`
	CreateBy   string    `gorm:"create_by" json:"create_by"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	UpdateBy   string    `gorm:"update_by" json:"update_by"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
	Remark     string    `gorm:"remark" json:"remark"`
}

// CreateUpdate 更新信息创建
func CreateUpdate(u HotelRole, s string) {
	n := Hotel_Update{}
	n.ServeName = s + " Hotel_Update"
	n.CreateBy = u.RoleName
	n.CreateTime = time.Now()
	global.App.DB.Table("hotel_update").Create(&n)
}
