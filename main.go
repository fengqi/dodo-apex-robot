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
	_ = logger.InitZap(context.TODO())
}

func main() {
	http.HandleFunc("/images/", cache.ImageHandler)
	http.HandleFunc("/webhook", webhook.Handler)

	logger.Zap().Info("Starting server at port " + strconv.Itoa(config.Port))
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Port), nil); err != nil {
		logger.Zap().Fatal("ListenAndServe", zap.Error(err))
	}
}
