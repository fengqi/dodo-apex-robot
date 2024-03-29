package webhook

import (
	"encoding/json"
	als "fengqi/dodo-apex-robot/apex_legends_status"
	"fengqi/dodo-apex-robot/cache"
	"fengqi/dodo-apex-robot/config"
	"fengqi/dodo-apex-robot/dodo"
	"fengqi/dodo-apex-robot/logger"
	"fengqi/dodo-apex-robot/translate"
	"fengqi/dodo-apex-robot/utils"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

// 3002 文本消息
func textMessageHandle(w http.ResponseWriter, r *http.Request, msg dodo.EventBodyChannelMessage) {
	var textBody dodo.TextBody
	if err := json.Unmarshal(msg.MessageBody, &textBody); err != nil {
		logger.Client.Error("unmarshal text body err", zap.Any("MessageBody", msg.MessageBody), zap.Error(err))
		return
	}

	// 识别命令
	cmd := ClearCmd(textBody.Content)
	parsed := ParseCmd(cmd)
	if parsed == "" {
		return
	}

	var card dodo.CardMessage
	switch parsed {
	case CmdUser:
		card = cmdQueryPlayer(cmd[7:], msg)
		break

	case CmdMap:
		card = cmdQueryMap(cmd)
		break

	case CmdCraft:
		card = cmdQueryCraft(cmd)
		break

	case CmdTime:
		card = cmdQueryTime(cmd)
		break

	case CmdDist:
		card = cmdDistribution(cmd)
		break

	case CmdPick:
		card = cmdLegendPick(cmd)
		break

	default:
		card = cmdHelp()
		break
	}

	send := dodo.SendChannelRequest{
		ChannelId:           msg.ChannelId,
		MessageType:         6,
		MessageBody:         card,
		ReferencedMessageId: msg.MessageId,
		Retry:               2,
	}
	if err := dodo.SetChannelMessageSend(send); err != nil {
		logger.Client.Error("send dodo channel message err", zap.Any("message", send), zap.Error(err))
	}
}

func cmdLegendPick(cmd string) (card dodo.CardMessage) {
	card = dodo.CardMessage{
		Content: "",
		Card: dodo.CardBody{
			Type:  "card",
			Title: "排位段位分布",
			Theme: dodo.RandTheme(),
			Components: []any{
				dodo.TextAndComponent{
					Type:  "section",
					Align: "right",
					Text: dodo.TextData{
						Type:    "dodo-md",
						Content: "暂无数据源，请前往Apex Legends Status查看。",
					},
					Accessory: dodo.ButtonCardData{
						Type:  "button",
						Name:  "点此前往",
						Color: dodo.RandColor(),
						Click: dodo.ButtonClickAction{
							Value:  "https://apex-status.speedapi.tk/game-stats/legends-pick-rates",
							Action: "link_url",
						},
					},
				},
			},
		},
	}
	return
}

func cmdDistribution(cmd string) (card dodo.CardMessage) {
	rank := als.GetRankDistribution(true)

	// 排序
	orderList := make([]als.RankData, len(rank)-1)
	for k, v := range rank {
		if order, ok := translate.RankOrder[k]; ok {
			orderList[order] = v
		}
	}

	total := 0
	content := ""
	for _, v := range orderList {
		total += v.Total
		content += fmt.Sprintf("%s\t占比：%.2f%%\t人数：%d\n", translate.RankNameZh(v.Name), v.Percent, v.Total)
	}

	card = dodo.CardMessage{
		Content: "",
		Card: dodo.CardBody{
			Type:  "card",
			Title: "排位段位分布",
			Theme: dodo.RandTheme(),
			Components: []any{
				dodo.TextCard{
					Type: "section",
					Text: dodo.TextData{
						Type:    "dodo-md",
						Content: content,
					},
				},
				dodo.RemarkCard{
					Type: "remark",
					Elements: []dodo.RemarkCardData{
						{
							Type:    "dodo-md",
							Content: fmt.Sprintf("统计人数：%d", total) + "，详细数据参考：[Ranked distribution statistics](https://apex-status.speedapi.tk/game-stats/ranked-distribution)",
						},
					},
				},
			},
		},
	}

	return
}

func cmdQueryTime(cmd string) (card dodo.CardMessage) {
	card = dodo.CardMessage{
		Content: "",
		Card: dodo.CardBody{
			Type:  "card",
			Title: "距离通行证结束还有",
			Theme: dodo.RandTheme(),
			Components: []any{
				dodo.CountdownCard{
					Type:    "countdown",
					Title:   "",
					Style:   "day",
					EndTime: config.Season.End,
				},
			},
		},
	}
	return
}

func cmdQueryCraft(cmd string) (card dodo.CardMessage) {
	logger.Client.Info("cmd query craft")

	list, err := als.GetCrafting()
	if err != nil {
		logger.Client.Error("get crafting error", zap.Error(err))
		return
	}

	img := make(map[string]string, 0)
	for _, bundle := range list {
		end, _ := utils.ParseTimestamp(bundle.End)
		if bundle.End > 0 && end.Before(time.Now()) {
			continue
		}
		for _, item := range bundle.BundleContent {
			img[item.ItemType.Name] = item.ItemType.Asset
		}
	}

	var imageTitle = ""
	var imageGroup []dodo.ImageCard
	for k, v := range img {
		// 过滤子弹、大血包、大电池、进化盾
		if k == "evo_armor" || k == "ammo" || k == "med_kit" || k == "large_shield_cell" {
			continue
		}
		imageTitle += k + "、"
		imageGroup = append(imageGroup, dodo.ImageCard{
			Type: "image",
			Src:  cache.CacheImage(v),
		})
		if len(imageGroup) >= 9 {
			break
		}
	}

	card = dodo.CardMessage{
		Content: "",
		Card: dodo.CardBody{
			Type:  "card",
			Title: "当前复制器道具",
			Theme: dodo.RandTheme(),
			Components: []any{
				dodo.TextCard{
					Type: "section",
					Text: dodo.TextData{
						Type:    "dodo-md",
						Content: strings.TrimRight(imageTitle, "、"),
					},
				},
				dodo.ImageGroupCard{
					Type:     "image-group",
					Elements: imageGroup,
				},
				dodo.RemarkCard{
					Type: "remark",
					Elements: []dodo.RemarkCardData{
						{
							Type:    "dodo-md",
							Content: "详细数据参考：[Replicators crafting rotation](https://apex-status.speedapi.tk/crafting-rotation)",
						},
					},
				},
			},
		},
	}

	return card
}

func cmdQueryPlayer(player string, msg dodo.EventBodyChannelMessage) (card dodo.CardMessage) {
	player = strings.TrimSpace(player)
	if len(player) == 0 {
		if val, ok := cache.Client.Get("dodo-" + msg.DodoSourceId); ok {
			player = val.(string)
		}
		if len(player) == 0 {
			return cmdHelp()
		}
	}

	cache.Client.Set("dodo-"+msg.DodoSourceId, player, -1)
	logger.Client.Info("cmd query player", zap.String("player", player))

	bridge, err := als.GetBridge(player, "pc")
	if err != nil {
		logger.Client.Error("get bridge error", zap.String("player", player), zap.Error(err))
		return
	}

	title := player + "的EA信息"
	if bridge.Global.Name != "" {
		title = bridge.Global.Name + "的EA信息"
	}

	content := fmt.Sprintf(
		"等级：%d\n分数：%d\n段位：%s%d\n排名：%d\nTOP：%v%%",
		bridge.Global.Level,
		bridge.Global.Rank.RankScore,
		translate.RankNameZh(bridge.Global.Rank.RankName),
		bridge.Global.Rank.RankDiv,
		bridge.Global.Rank.ALStopInt,
		bridge.Global.Rank.ALStopPercent,
	)
	card = dodo.CardMessage{
		Content: "",
		Card: dodo.CardBody{
			Type:  "card",
			Title: title,
			Theme: dodo.RandTheme(),
			Components: []any{
				dodo.TextAndComponent{
					Type:  "section",
					Align: "left",
					Text: dodo.TextData{
						Type:    "dodo-md",
						Content: content,
					},
					Accessory: dodo.ImageCard{
						Type: "image",
						Src:  cache.CacheImage(bridge.Global.Rank.RankImg),
					},
				},
			},
		},
	}
	return
}

func cmdQueryMap(cmd string) (card dodo.CardMessage) {
	logger.Client.Info("cmd query map")

	rot, err := als.GetMapRotation()
	if err != nil {
		logger.Client.Error("get map rotation error", zap.Error(err))
		return
	}

	card = dodo.CardMessage{
		Content: "",
		Card: dodo.CardBody{
			Type:  "card",
			Title: "排位地图轮换",
			Theme: dodo.RandTheme(),
			Components: []any{
				dodo.TextCard{
					Type: "section",
					Text: dodo.TextData{
						Type:    "dodo-md",
						Content: fmt.Sprintf("当前地图：%s，剩余时间：%s", translate.MapNameZh(rot.Ranked.Current.Map), rot.Ranked.Current.RemainingTimer),
					},
				},
				dodo.ImageCard{
					Type: "image",
					Src:  cache.ResizeImage(cache.CacheImage(rot.Ranked.Current.Asset), 1024000),
				},
				dodo.TextCard{
					Type: "section",
					Text: dodo.TextData{
						Type:    "dodo-md",
						Content: fmt.Sprintf("下一轮地图：%s，开启时间：%s", translate.MapNameZh(rot.Ranked.Next.Map), utils.TimestampFormat(rot.Ranked.Next.Start)),
					},
				},
				dodo.ImageCard{
					Type: "image",
					Src:  cache.ResizeImage(cache.CacheImage(rot.Ranked.Next.Asset), 1024000),
				},
				dodo.RemarkCard{
					Type: "remark",
					Elements: []dodo.RemarkCardData{
						{
							Type:    "dodo-md",
							Content: "详细数据参考：[Current Apex Legends map in rotation](https://apex-status.speedapi.tk/current-map)",
						},
					},
				},
			},
		},
	}

	return
}

func cmdHelp() (card dodo.CardMessage) {
	content := ""
	for k, v := range cmdMap {
		content += fmt.Sprintf("%s: %s\n", k, v)
	}

	card = dodo.CardMessage{
		Content: "",
		Card: dodo.CardBody{
			Type:  "card",
			Title: "使用帮助",
			Theme: dodo.RandTheme(),
			Components: []any{
				dodo.TextCard{
					Type: "section",
					Text: dodo.TextData{
						Type:    "dodo-md",
						Content: content,
					},
				},
				dodo.RemarkCard{
					Type: "remark",
					Elements: []dodo.RemarkCardData{
						{
							Type:    "dodo-md",
							Content: "注：部分指令需要查询第三方系统，请耐心等待，不要重复发送指令。",
						},
					},
				},
			},
		},
	}

	return
}

func waitNotice(w http.ResponseWriter, r *http.Request, msg dodo.EventBodyChannelMessage) {
	send := dodo.SendChannelRequest{
		ChannelId:           msg.ChannelId,
		MessageType:         1,
		MessageBody:         dodo.TextBody{Content: "正在查询，请稍后..."},
		ReferencedMessageId: msg.MessageId,
	}
	if err := dodo.SetChannelMessageSend(send); err != nil {
		logger.Client.Error("send dodo channel message err", zap.Any("message", send), zap.Error(err))
	}
}
