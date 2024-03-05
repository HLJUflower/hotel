package sql

import (
	"github.com/jinzhu/gorm"
	"hotel/utils/logging"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:zl20040203@(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logging.Errorf("DB connect failed : ", err)
		return nil
	}
	return db
}
