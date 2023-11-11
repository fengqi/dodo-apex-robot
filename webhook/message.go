package webhook

// 2001 消息时间

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/cache"
	"fengqi/dodo-apex-robot/dodo"
	"fengqi/dodo-apex-robot/logger"
	"go.uber.org/zap"
	"net/http"
)

// messageHandle 文本消息事件处理
func messageHandle(w http.ResponseWriter, r *http.Request, business dodo.Business) {
	var msg dodo.EventBodyChannelMessage
	err := json.Unmarshal(business.EventBody, &msg)
	if err != nil {
		logger.Client.Error("unmarshal EventBodyChannelMessage err", zap.Any("EventBody", business.EventBody), zap.Error(err))
		return
	}

	if _, ok := cache.Client.Get(msg.MessageId); ok {
		logger.Client.Warn("message already reply", zap.Any("msg", msg))
		return
	}

	switch msg.MessageType {
	case 1: // 文本消息
		textMessageHandle(w, r, msg)
		break

	default:
		// 其他的暂时不适配
		break
	}
}
