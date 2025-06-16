package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config/config-dev") // 配置文件名(不带扩展名)
	viper.SetConfigType("yaml")              // 配置文件类型
	viper.AddConfigPath(".")                 // 配置文件路径

	// 设置环境变量前缀并自动加载
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Failed to read config file: %v", err)
	}
}
