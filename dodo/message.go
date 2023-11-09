package dodo

import (
	"bytes"
	"encoding/json"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/message"
	"fmt"
	"io"
	"net/http"
)

func SetChannelMessageSend(data message.SendChannelRequest) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	api := config.Dodo.Host + ApiSendMessage
	req, err := http.NewRequest(http.MethodPost, api, bytes.NewReader(marshal))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bot "+config.Dodo.ClientId+"."+config.Dodo.BotToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %s\n", content)

	return nil
}
