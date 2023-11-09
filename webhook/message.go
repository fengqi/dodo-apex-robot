package webhook

// 2001 消息时间

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/message"
	"net/http"
)

// messageHandle 文本消息事件处理
func messageHandle(w http.ResponseWriter, r *http.Request, business message.Business) {
	var msg message.EventBodyChannelMessage
	err := json.Unmarshal(business.EventBody, &msg)
	if err != nil {
		panic(err)
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
