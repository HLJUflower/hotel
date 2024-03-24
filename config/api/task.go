package api

//季度计划
import (
	"hotel/global"
	"time"
)

type Task struct {
	JobId      int       `gorm:"job_id" json:"job_id"`
	JobName    string    `gorm:"job_name" json:"job_name"`
	JobGroup   string    `gorm:"job_group" json:"job_group"`
	JobDetail  string    `gorm:"job_detail" json:"job_detail"`
	Status     string    `gorm:"status" json:"status"`
	CreateBy   string    `gorm:"create_by" json:"create_by"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	UpdateBy   string    `gorm:"update_by" json:"update_by"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
	Remark     string    `gorm:"remark" json:"remark"`
}

// Create 季度信息创建
func (r Task) Create(u HotelRole) {
	CreateUpdate(u, "task")
	global.App.DB.Table("hotel_post").Create(r)
}

// Update 季度信息更新
func (r Task) Update(u HotelRole) {
	CreateUpdate(u, "task")
	global.App.DB.Table("hotel_post").Save(r)
}
