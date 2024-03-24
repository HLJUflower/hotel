package api

//角色信息
import (
	"hotel/global"
	"time"
)

type HotelRole struct {
	RoleId     int       `gorm:"role_id"`
	RoleName   string    `gorm:"role_name"`
	RolePasswd string    `gorm:"role_passwd"`
	Phone      string    `gorm:"phone"`
	RoleKey    string    `gorm:"role_key"`
	RoleSort   int       `gorm:"role_sort"`
	DataScope  string    `gorm:"data_scope"`
	Status     string    `gorm:"status"`
	DelFlag    string    `gorm:"del_flag"`
	CreateBy   string    `gorm:"create_by"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateBy   string    `gorm:"update_by"`
	UpdateTime time.Time `gorm:"update_time"`
	Remark     string    `gorm:"remark"`
}

func (r HotelRole) TableName() string {
	return "login_infor"
}

//增删改查 判断权限

// FindHotelPost 岗位信息查找
func FindHotelPost(s string, u HotelRole) []HotelPost {
	var user []HotelPost
	global.App.DB.Table("hotel_role").Where("post_code = ?", s).Find(&user)
	CreateUpdate(u, "hotel_post")
	return user
}

// FindHotelRole 角色信息查询
func FindHotelRole(s string, u HotelRole) []HotelRole {
	var res []HotelRole
	global.App.DB.Table("hotel_role").Where("role_key = ?", s).Find(res)
	CreateUpdate(u, "hotel_role")
	return res
}

// Find_login_infor 访客信息查询
func Find_login_infor(s string, u HotelRole) []LoginInfor {
	CreateUpdate(u, "login_infor")
	var res []LoginInfor
	global.App.DB.Table("login_infor").Where("user_name = ?", s).Find(res)
	return res
}

// Find_menu 账单信息查询
func Find_menu(s string, u HotelRole) []Menu {
	CreateUpdate(u, "menu")
	var res []Menu
	global.App.DB.Table("menu").Where("room = ?", s).Find(res)
	return res
}

// Find_rooms 房间信息查询
func Find_rooms(s string, u HotelRole) []Rooms {
	CreateUpdate(u, "rooms")
	var res []Rooms
	global.App.DB.Table("rooms").Where("status = ?", s).Find(res)
	return res
}

// Find_sleep_user 客户信息查询
func Find_sleep_user(s string, u HotelRole) []SleepUser {
	CreateUpdate(u, "sleep_user")
	var res []SleepUser
	global.App.DB.Table("sleep_user").Where("role_name = ?", s).Find(res)
	return res
}

// Find_task 计划信息查询
func Find_task(s string, u HotelRole) []Task {
	CreateUpdate(u, "sleep_user")
	var res []Task
	global.App.DB.Table("update").Where("serve_name = ?", s).Find(res)
	return res
}

// Find_update 更新信息查询
func Find_update(s string, u HotelRole) []Hotel_Update {
	CreateUpdate(u, "sleep_user")
	var res []Hotel_Update
	global.App.DB.Table("task").Where("job_name = ?", s).Find(res)
	return res
}

// Delete 角色信息删除
func (r HotelRole) Delete(u HotelRole, del HotelRole) {
	CreateUpdate(u, "hotel_role")
	global.App.DB.Table("hotel_role").Delete(del)
}

// Update 角色信息更新
func (r HotelRole) Update(u HotelRole, updt HotelRole) {
	CreateUpdate(u, "hotel_role")
	global.App.DB.Table("hotel_role").Update(updt)
}
