package core

import (
	"fmt"
	"frame-web/core/internal"
	"frame-web/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapInit() (logger *zap.Logger) {
	fmt.Printf("create %s directory\n", global.CONFIG.Zap.Director)
	levels := global.CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
