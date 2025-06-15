package db

import (
	zlog "go-log-v/server/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "aaaa"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zlog.Error("Failed to connect to database", "error", err)
		panic("failed to connect database")
	}
	zlog.Info("Database connection established")
}
