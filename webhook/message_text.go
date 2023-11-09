package webhook

import (
	"encoding/json"
	als "fengqi/dodo-apex-robot/apex_legends_status"
	"fengqi/dodo-apex-robot/dodo"
	"fengqi/dodo-apex-robot/message"
	"fengqi/dodo-apex-robot/utils"
	"fmt"
	"net/http"
	"strings"
)

// 3002 文本消息
func textMessageHandle(w http.ResponseWriter, r *http.Request, msg message.EventBodyChannelMessage) {
	var textBody message.TextBody
	if err := json.Unmarshal(msg.MessageBody, &textBody); err != nil {
		panic(err)
	}

	// 识别命令
	cmd := ClearCmd(textBody.Content)
	if cmd == "" {
		return
	}

	var card message.CardMessage
	switch ParseCmd(cmd) {
	case CmdUser:
		player := strings.TrimSpace(cmd[7:])
		card = cmdQueryPlayer(player)
		break

	case CmdMap:
		card = cmdQueryMap(cmd)
		break

	default:
		card = cmdHelp()
		break
	}

	send := message.SendChannelRequest{
		ChannelId:           msg.ChannelId,
		MessageType:         6,
		MessageBody:         card,
		ReferencedMessageId: msg.MessageId,
	}
	if err := dodo.SetChannelMessageSend(send); err != nil {
		panic(err)
	}
}

func cmdQueryPlayer(player string) (card message.CardMessage) {
	//player := utils.MatchPlayerName(cmd)
	if player == "" {
		return
	}

	fmt.Printf("cmd query player: %s\n", player)

	bridge, err := als.GetBridge(player, "pc")
	if err != nil {
		panic(err)
	}

	title := player + "的EA信息"
	if bridge.Global.Name != "" {
		title = bridge.Global.Name + "的EA信息"
	}

	content := fmt.Sprintf(
		"等级：%d\n分数：%d\n段位：%s%d\n排名：%d\n排名：%v%",
		bridge.Global.Level,
		bridge.Global.Rank.RankScore,
		utils.RankNameZh(bridge.Global.Rank.RankName),
		bridge.Global.Rank.RankDiv,
		bridge.Global.Rank.ALStopInt,
		bridge.Global.Rank.ALStopPercent,
	)
	card = message.CardMessage{
		Content: "",
		Card: message.CardBody{
			Type:  "card",
			Title: title,
			Theme: "default",
			Components: []any{
				struct {
					Type      string `json:"type"`
					Align     string `json:"align"`
					Text      any    `json:"text"`
					Accessory any    `json:"accessory"`
				}{
					Type:  "section",
					Align: "left",
					Text: struct {
						Type    string `json:"type"`
						Content string `json:"content"`
					}{
						Type:    "dodo-md",
						Content: content,
					},
					Accessory: struct {
						Type string `json:"type"`
						Src  string `json:"src"`
					}{
						Type: "image",
						Src:  bridge.Global.Rank.RankImg,
					},
				},
			},
		},
	}
	return
}

func cmdQueryMap(cmd string) (card message.CardMessage) {
	match := utils.MatchIsMap(cmd)
	if !match {
		return
	}

	fmt.Println("cmd query map")

	rot, err := als.GetMapRotation()
	if err != nil {
		panic(err)
	}

	card = message.CardMessage{
		Content: "",
		Card: message.CardBody{
			Type:  "card",
			Title: "排位地图轮换",
			Theme: "default",
			Components: []any{
				struct {
					Type string `json:"type"`
					Text any    `json:"text"`
				}{
					Type: "section",
					Text: struct {
						Type    string `json:"type"`
						Content string `json:"content"`
					}{
						Type:    "dodo-md",
						Content: fmt.Sprintf("当前排位地图：%s，剩余时间：%s", rot.Ranked.Current.Map, rot.Ranked.Current.RemainingTimer),
					},
				},
				struct {
					Type string `json:"type"`
					Src  string `json:"src"`
				}{
					Type: "image",
					Src:  rot.Ranked.Current.Asset,
				},
				struct {
					Type string `json:"type"`
					Text any    `json:"text"`
				}{
					Type: "section",
					Text: struct {
						Type    string `json:"type"`
						Content string `json:"content"`
					}{
						Type:    "dodo-md",
						Content: fmt.Sprintf("下一轮排位地图：%s，开启时间：%s", rot.Ranked.Next.Map, utils.TimestampFormat(rot.Ranked.Next.Start)),
					},
				},
				struct {
					Type string `json:"type"`
					Src  string `json:"src"`
				}{
					Type: "image",
					Src:  rot.Ranked.Next.Asset,
				},
			},
		},
	}

	return
}

func cmdHelp() (card message.CardMessage) {
	content := ""
	for k, v := range cmdMapZh {
		content += fmt.Sprintf("%s: %s\n", k, v)
	}

	card = message.CardMessage{
		Content: "",
		Card: message.CardBody{
			Type:  "card",
			Title: "使用帮助",
			Theme: "default",
			Components: []any{
				struct {
					Type string `json:"type"`
					Text any    `json:"text"`
				}{
					Type: "section",
					Text: struct {
						Type    string `json:"type"`
						Content string `json:"content"`
					}{
						Type:    "dodo-md",
						Content: content,
					},
				},
			},
		},
	}

	return
}