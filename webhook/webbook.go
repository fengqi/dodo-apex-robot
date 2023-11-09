package webhook

import (
	"encoding/json"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/message"
	"fengqi/dodo-apex-robot/utils"
	"fmt"
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
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)

	fmt.Printf("webhook: %s\n", bytes)

	var req message.Request
	if err := json.Unmarshal(bytes, &req); err != nil {
		panic(err)
	}

	// 解密回调数据
	bytes, err = utils.AesDecrypt(config.Dodo.SecretKey, req.Payload)
	if err != nil {
		panic(err)
	}

	var subject message.Subject
	if err := json.Unmarshal(bytes, &subject); err != nil {
		panic(err)
	}
	fmt.Printf("subject: %s\n", bytes)

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
	var business message.Business
	if err := json.Unmarshal(raw, &business); err != nil {
		panic(err)
	}

	switch business.EventType {
	case "2001": // 消息事件
		messageHandle(w, r, business)
		return

	default: // 其他的暂时不适配
		return
	}
}
