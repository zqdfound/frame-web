package global

import (
	"frame-web/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GLOBAL_VP     *viper.Viper
	GLOBAL_CONFIG config.Server
	GLOBAL_LOG    *zap.Logger
	GVA_DB        *gorm.DB
)
