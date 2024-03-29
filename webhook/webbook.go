package webhook

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/common"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/dodo"
	"fengqi/dodo-apex-robot/logger"
	"fengqi/dodo-apex-robot/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Client.Error("webhook read body err", zap.Error(err))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Client.Error("webhook close body err", zap.Error(err))
		}
	}(r.Body)

	logger.Client.Debug("webhook", zap.ByteString("body", bytes))

	var req common.Request
	if err := json.Unmarshal(bytes, &req); err != nil {
		logger.Client.Error("webhook unmarshal request err", zap.Error(err))
		return
	}

	// 解密回调数据
	bytes, err = utils.AesDecrypt(config.Dodo.SecretKey, req.Payload)
	if err != nil {
		logger.Client.Error("webhook decrypt err", zap.Error(err))
		return
	}

	var subject dodo.Subject
	if err := json.Unmarshal(bytes, &subject); err != nil {
		logger.Client.Error("webhook unmarshal subject err", zap.Error(err))
		return
	}

	logger.Client.Debug("subject", zap.Any("subject", subject))

	// 地址校验，原样返回收到的信息，证明解密成功
	if subject.Type == 2 {
		utils.ResponseData(w, subject.Data)
		return
	}

	go business(w, r, subject.Data)
	utils.Response(w)
}

// 业务消息处理
func business(w http.ResponseWriter, r *http.Request, raw json.RawMessage) {
	var business dodo.Business
	if err := json.Unmarshal(raw, &business); err != nil {
		logger.Client.Error("business unmarshal err", zap.Error(err))
		return
	}

	switch business.EventType {
	case "2001": // 消息事件
		messageHandle(w, r, business)
		return

	default: // 其他的暂时不适配
		return
	}
}
