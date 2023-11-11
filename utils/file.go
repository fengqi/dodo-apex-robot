package utils

import (
	"fengqi/dodo-apex-robot/logger"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CheckPath(path string) {
	path = filepath.Dir(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			logger.Client.Error("create path err", zap.String("path", path), zap.Error(err))
		}
	}
}
