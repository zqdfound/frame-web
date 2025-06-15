package zap

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

func Info(msg string, args ...interface{}) {
	fields := make([]zap.Field, 0, len(args))
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			fields = append(fields, zap.String("value", v))
		case int:
			fields = append(fields, zap.Int("value", v))
		case error:
			fields = append(fields, zap.Error(v))
		default:
			fields = append(fields, zap.Any("value", v))
		}
	}
	zap.L().Info(msg, fields...)
}

func Error(msg string, args ...interface{}) {
	fields := make([]zap.Field, 0, len(args))
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			fields = append(fields, zap.String("value", v))
		case int:
			fields = append(fields, zap.Int("value", v))
		case error:
			fields = append(fields, zap.Error(v))
		default:
			fields = append(fields, zap.Any("value", v))
		}
	}
	zap.L().Error(msg, fields...)
}
