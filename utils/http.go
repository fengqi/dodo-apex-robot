package utils

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/logger"
	"fengqi/dodo-apex-robot/message"
	"go.uber.org/zap"
	"net/http"
)

func Response(w http.ResponseWriter) {
	bytes, _ := json.Marshal(&message.Response{})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func ResponseData(w http.ResponseWriter, data interface{}) {
	logger.Zap().Debug("response", zap.Any("data", data))

	bytes, _ := json.Marshal(&message.Response{Data: data})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
