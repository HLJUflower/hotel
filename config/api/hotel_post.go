package api

//岗位信息
import (
	"hotel/global"
	"time"
)

type HotelPost struct {
	PostId     int       ` gorm:"post_id" json:"post_id,omitempty"`
	PostCode   string    ` gorm:"post_code" json:"post_code,omitempty"`
	PostName   string    `gorm:"post_name" json:"post_name,omitempty"`
	PostSort   int       `gorm:"post_sort" json:"post_sort,omitempty"`
	Status     string    `gorm:"status" json:"status,omitempty"`
	CreateBy   string    `gorm:"create_by" json:"create_by,omitempty"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	UpdateBy   string    `gorm:"update_by" json:"update_by,omitempty"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
	Remark     string    `gorm:"remark" json:"remark,omitempty"`
}

// Create 岗位信息创建
func (r HotelPost) Create(u HotelRole) {
	n := Hotel_Update{}
	n.UpdateId = u.RoleId
	n.ServeName = "hotel_post" + " Create"
	n.CreateBy = u.RoleName
	n.CreateTime = time.Now()
	global.App.DB.Table("update").Create(n)
	global.App.DB.Table("hotel_post").Create(r)
}

// Delete 岗位信息删除
func (r HotelPost) Delete(u HotelRole) {
	CreateUpdate(u, "hotel_post")
	global.App.DB.Table("hotel_post").Delete(r)

}

// Update 岗位信息更新
func (r HotelPost) Update(u HotelRole) {
	CreateUpdate(u, "hotel_post")
	global.App.DB.Table("hotel_post").Save(r)
}
