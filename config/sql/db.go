package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"hotel/utils/logging"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:zl20040203@(localhost:3306)/hotel?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logging.Errorf("DB connect failed : ", err)
		return nil
	}
	//db.Exec("config/sql/db.sql")
	return db
}
