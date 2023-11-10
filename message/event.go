package message

import "encoding/json"

// Subject 事件主体
type Subject struct {
	Type    int             `json:"type"`    // 数据类型，0：业务数据
	Data    json.RawMessage `json:"data"`    // 业务数据
	Version string          `json:"version"` // 业务版本
}

// Business 业务数据
type Business struct {
	EventBody json.RawMessage `json:"eventBody"` // 事件内容
	EventId   string          `json:"eventId"`   // 事件ID
	EventType string          `json:"eventType"` // 事件类型
	Timestamp int64           `json:"timestamp"` // 接收时间戳
}

// EventBodyChannelMessage MessageEvent 文字消息
type EventBodyChannelMessage struct {
	IslandSourceId string          `json:"islandSourceId"` // 来源群ID
	ChannelId      string          `json:"channelId"`      // 来源频道ID
	DodoSourceId   string          `json:"dodoSourceId"`   // 来源DoDoID
	MessageId      string          `json:"messageId"`      // 消息ID
	Personal       Personal        `json:"personal"`       // 个人信息
	Member         Member          `json:"member"`         // 成员信息
	Reference      Reference       `json:"reference"`      // 回复信息
	MessageType    int             `json:"messageType"`    // 消息类型，1：文字消息，2：图片消息，3：视频消息，4：分享消息，5：文件消息，6：卡片消息，7：红包消息
	MessageBody    json.RawMessage `json:"messageBody"`    // 消息内容
}

// ImageBody 图片消息
type ImageBody struct {
	Url        string `json:"url"`        // 图片链接，必须是官方的链接，通过上传资源图片接口可获得图片链接
	Width      int    `json:"width"`      // 图片宽度
	Height     int    `json:"height"`     // 图片高度
	IsOriginal int    `json:"isOriginal"` // 是否原图，0：压缩图，1：原图
}
