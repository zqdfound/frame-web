package zap

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 创建文件输出
	fileSyncer := getWriterSyncer()
	// 创建控制台输出
	consoleSyncer := zapcore.AddSync(os.Stdout)

	// 使用MultiWriteSyncer同时输出到文件和控制台
	multiSyncer := zapcore.NewMultiWriteSyncer(fileSyncer, consoleSyncer)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		multiSyncer,
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

func L() *zap.Logger {
	return &zap.Logger{}
}

func Info(msg string, args ...interface{}) {
	fields := make([]zap.Field, 0, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		switch v := args[i+1].(type) {
		case string:
			fields = append(fields, zap.String(key, v))
		case int:
			fields = append(fields, zap.Int(key, v))
		case error:
			fields = append(fields, zap.Error(v))
		default:
			fields = append(fields, zap.Any(key, v))
		}
	}
	zap.L().Info(msg, fields...)
}

func Error(msg string, args ...interface{}) {
	fields := make([]zap.Field, 0, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		switch v := args[i+1].(type) {
		case string:
			fields = append(fields, zap.String(key, v))
		case int:
			fields = append(fields, zap.Int(key, v))
		case error:
			fields = append(fields, zap.Error(v))
		default:
			fields = append(fields, zap.Any(key, v))
		}
	}
	zap.L().Error(msg, fields...)
}
