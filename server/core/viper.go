package core

import (
	"flag"
	"fmt"
	"frame-web/global"
	"github.com/spf13/viper"
)

// 初始化viper
func Viper() *viper.Viper {
	var config string
	//从命令行获取
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" { // 判断命令行参数是否为空
		config = "config/config-dev.yaml" // 默认配置文件路径
		fmt.Printf("使用默认配置文件：%s", config)
	}
	v := viper.New()
	v.SetConfigFile(config) // 配置文件名(不带扩展名)
	v.SetConfigType("yaml") // 配置文件类型
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err = v.Unmarshal(&global.GLOBAL_CONFIG); err != nil { //绑定配置
		panic(err)
	}
	return v
}
