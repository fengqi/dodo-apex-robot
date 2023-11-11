package main

import (
	"context"
	"fengqi/dodo-apex-robot/cache"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/logger"
	"fengqi/dodo-apex-robot/webhook"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func init() {
	config.Load("./config.json")
	logger.InitZap(context.TODO())
	cache.InitCache(config.CachePath)
}

func main() {
	http.HandleFunc("/images/", cache.ImageHandler)
	http.HandleFunc("/webhook", webhook.Handler)

	logger.Client.Info("Starting server at port " + strconv.Itoa(config.Port))
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Port), nil); err != nil {
		logger.Client.Fatal("ListenAndServe", zap.Error(err))
	}
}
