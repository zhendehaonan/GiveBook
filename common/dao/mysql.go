package dao

import (
	"GiveBook/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1234@(localhost)/db2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.SingularTable(true)
	db.AutoMigrate(&model.User{})
	DB = db
	return DB
}

// 获取DB对象
func GetDB() *gorm.DB {
	return DB
}
