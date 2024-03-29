package dodo

import (
	"bytes"
	"encoding/json"
	"fengqi/dodo-apex-robot/cache"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/logger"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	retry chan SendChannelRequest
)

func init() {
	retry = make(chan SendChannelRequest, 1000)
	go runRetry()
}

func runRetry() {
	for {
		select {
		case msg := <-retry:
			err := SetChannelMessageSend(msg)
			if err != nil {
				log.Printf("retry error: %s\n", err)
			}
			break
		}
	}
}

func InQueue(data SendChannelRequest) {
	retry <- data
}

// SetChannelMessageSend 发送频道消息
func SetChannelMessageSend(data SendChannelRequest) error {
	logger.Client.Debug("SetChannelMessageSend", zap.Any("data", data))

	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}

	api := config.Dodo.Host + ApiSendMessage
	req, err := http.NewRequest(http.MethodPost, api, bytes.NewReader(marshal))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bot "+config.Dodo.ClientId+"."+config.Dodo.BotToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	logger.Client.Debug("response", zap.ByteString("content", content))

	var mr MessageResponse
	if err := json.Unmarshal(content, &mr); err != nil {
		return err
	}

	if data.Retry > 0 && mr.Status == -9999 { // 失败重试一次
		data.Retry -= 1
		go func() {
			time.Sleep(time.Second * 3)
			retry <- data
		}()
	}

	if mr.Status == 0 {
		cache.Client.Set(data.ReferencedMessageId, 1, 1*time.Hour)
	}

	return nil
}
