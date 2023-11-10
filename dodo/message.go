package dodo

import (
	"bytes"
	"encoding/json"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/message"
	"fengqi/dodo-apex-robot/model"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	retry chan message.SendChannelRequest
)

func init() {
	retry = make(chan message.SendChannelRequest, 1000)
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

func InQueue(data message.SendChannelRequest) {
	retry <- data
}

// SetChannelMessageSend 发送频道消息
func SetChannelMessageSend(data message.SendChannelRequest) error {
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

	fmt.Printf("response: %s\n", content)
	var mr model.MessageResponse
	if err := json.Unmarshal(content, &mr); err != nil {
		return err
	}

	if mr.Status == -9999 { // 失败重试一次
		go func() {
			time.Sleep(time.Second * 1)
			retry <- data
		}()
	}

	return nil
}
