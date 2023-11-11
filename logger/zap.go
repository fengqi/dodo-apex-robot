package logger

import (
	"context"
	"fengqi/dodo-apex-robot/config"
	"go.uber.org/zap"
)

var Client *zap.Logger

var levelMap = map[string]zap.AtomicLevel{
	"debug": zap.NewAtomicLevelAt(zap.DebugLevel),
	"info":  zap.NewAtomicLevelAt(zap.InfoLevel),
	"warn":  zap.NewAtomicLevelAt(zap.WarnLevel),
	"error": zap.NewAtomicLevelAt(zap.ErrorLevel),
	"panic": zap.NewAtomicLevelAt(zap.PanicLevel),
	"fatal": zap.NewAtomicLevelAt(zap.FatalLevel),
}

func InitZap(ctx context.Context) {
	var zc = zap.NewProductionConfig()

	if level, ok := levelMap[config.Log.Level]; ok {
		zc.Level = level
	}

	if config.Log.Level != "debug" && config.Log.Path != "" {
		zc.OutputPaths = []string{config.Log.Path}
		zc.ErrorOutputPaths = []string{config.Log.Path}
	}

	zc.Encoding = "json"

	var err error
	Client, err = zc.Build()
	if err != nil {
		panic(err)
	}
	Client.Info("zap logger init success")
}
