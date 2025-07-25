package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Device Device `mapstructure:"device" json:"device" yaml:"device"`
	//oss
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
}
