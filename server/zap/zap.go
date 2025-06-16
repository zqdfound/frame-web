package zap

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		getWriterSyncer(),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}
func getWriterSyncer() zapcore.WriteSyncer {
	// 确保日志目录存在
	logDir := "./log"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 创建或打开日志文件
	logPath := filepath.Join(logDir, "log.out")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("创建日志文件失败: %v", err))
	}
	return zapcore.AddSync(file)
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
