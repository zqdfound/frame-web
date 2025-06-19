package global

import (
	"frame-web/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GLOBAL_VP     *viper.Viper
	GLOBAL_CONFIG config.Server
	GLOBAL_LOG    *zap.Logger
)
