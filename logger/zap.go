package logger

import (
	"context"
	"fengqi/dodo-apex-robot/config"
	"go.uber.org/zap"
)

var _zap *zap.Logger
var levelMap = map[string]zap.AtomicLevel{
	"debug": zap.NewAtomicLevelAt(zap.DebugLevel),
	"info":  zap.NewAtomicLevelAt(zap.InfoLevel),
	"warn":  zap.NewAtomicLevelAt(zap.WarnLevel),
	"error": zap.NewAtomicLevelAt(zap.ErrorLevel),
	"panic": zap.NewAtomicLevelAt(zap.PanicLevel),
	"fatal": zap.NewAtomicLevelAt(zap.FatalLevel),
}

func InitZap(ctx context.Context) error {
	var err error
	var zc = zap.NewProductionConfig()

	if level, ok := levelMap[config.Log.Level]; ok {
		zc.Level = level
	}

	if config.Log.Level != "debug" && config.Log.Path != "" {
		zc.OutputPaths = []string{config.Log.Path}
		zc.ErrorOutputPaths = []string{config.Log.Path}
	}

	zc.Encoding = "json"
	_zap, err = zc.Build()
	_zap.Info("zap logger init success")

	return err
}

func Zap() *zap.Logger {
	if _zap == nil {
		if err := InitZap(context.TODO()); err != nil {
			panic(err)
		}
	}
	return _zap
}
