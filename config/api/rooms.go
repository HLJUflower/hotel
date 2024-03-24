package api

//房间信息
import (
	"hotel/global"
	"time"
)

// Rooms 字典类型表
type Rooms struct {
	RoomId     int64     `json:"room_id" gorm:"room_id"`        //字典ID
	Leader     string    `json:"leader" gorm:"leader"`          //字典名称
	Phone      string    `json:"dictType" gorm:"phone"`         //字典类型
	Status     string    `json:"status" gorm:"status"`          //状态 0正常1停用
	DelFlag    string    `json:"del_flag" gorm:"del_flag"`      //备注
	CreateTime time.Time `json:"createTime" gorm:"create_time"` //创建时间
	CreateBy   string    `json:"createBy" gorm:"create_by"`     //创建人
	UpdateTime time.Time `json:"updateTime" gorm:"update_time"` //更新时间
	UpdateBy   string    `json:"updateBy" gorm:"update_by"`     //更新人
	Remark     string    `json:"remark" gorm:"remark"`
}

// Create 房间信息创建
func (r Rooms) Create(u HotelRole) {
	CreateUpdate(u, "rooms")
	global.App.DB.Table("hotel_post").Create(r)
}

// Delete 房间信息删除
func (r Rooms) Delete(u HotelRole) {
	CreateUpdate(u, "rooms")
	global.App.DB.Table("hotel_post").Delete(r)
}

// Update 房间信息更新
func (r Rooms) Update(u HotelRole) {
	CreateUpdate(u, "rooms")
	global.App.DB.Table("hotel_post").Save(r)
}
