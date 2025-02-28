package config

import (
	"uooobarry/gin-notes/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库连接
func Init() {
	var err error
	// 使用 SQLite 数据库
	DB, err = gorm.Open(sqlite.Open("data/app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}
