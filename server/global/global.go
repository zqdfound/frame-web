package global

import (
	"frame-web/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GLOBAL_VP *viper.Viper
	CONFIG    config.Server
	LOG       *zap.Logger
	DB        *gorm.DB
	REDIS     redis.UniversalClient
)
