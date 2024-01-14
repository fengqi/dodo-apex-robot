package cache

import (
	"fengqi/dodo-apex-robot/logger"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"time"
)

var Client *cache.Cache

func InitCache(path string) {
	Client = cache.New(5*time.Minute, 10*time.Minute)
	_ = Client.LoadFile(path)
	go func() {
		for range time.Tick(10 * time.Minute) {
			logger.Client.Debug("save go cache to file", zap.Int("count", Client.ItemCount()))
			_ = Client.SaveFile(path)
		}
	}()
}
