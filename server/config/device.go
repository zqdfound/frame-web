package config

type Device struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"` // 主机
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
