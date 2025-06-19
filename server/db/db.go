package db

//
//import (
//	zlog "frame-web/zap"
//	"github.com/spf13/viper"
//
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//var DB *gorm.DB
//
//func InitDB() {
//	dsn := viper.GetString("database.dsn")
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		zlog.Error("Failed to connect to database", "error", err)
//		panic("failed to connect database")
//	}
//	zlog.Info("Database connection established")
//}
