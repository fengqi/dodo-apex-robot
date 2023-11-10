package dodo

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

// Personal 个人信息
type Personal struct {
	NickName  string `json:"nickName"`  // DoDo昵称
	AvatarUrl string `json:"avatarUrl"` //头像
	Sex       int    `json:"sex"`       // 	性别，-1：保密，0：女，1：男
}

// Member 成员信息
type Member struct {
	NickName string `json:"nickName"` // 群昵称
	JoinTime string `json:"joinTime"` // 加群时间
}

// Reference 回复信息
type Reference struct {
	MessageId    string `json:"messageId"`    // 被回复消息ID
	DodoSourceId string `json:"dodoSourceId"` // 被回复者DoDoID
	NickName     string `json:"nickName"`     // 被回复者群昵称
}

// SendChannelRequest 发送频道消息
type SendChannelRequest struct {
	ChannelId           string `json:"channelId"`           // 频道ID
	MessageType         int    `json:"messageType"`         // 1：文字消息，2：图片消息，3：视频消息，6：卡片消息
	MessageBody         any    `json:"messageBody"`         // 消息内容
	ReferencedMessageId string `json:"referencedMessageId"` // 回复消息ID
	DodoSourceId        string `json:"dodoSourceId"`        // DoDoID，非必传，如果传了，则给该成员发送频道私信
	Retry               bool   `json:"-"`
}

// TextMessage 文字消息
type TextMessage struct {
	Content string `json:"content"`
}

// CardMessage 卡片消息
type CardMessage struct {
	Content string `json:"content"` // 附加文本，支持Markdown语法、菱形语法
	Card    any    `json:"card"`    // 卡片，限制10000个字符，支持Markdown语法，不支持菱形语法
}

// CardBody 卡片
type CardBody struct {
	Type       string `json:"type"`       // 类型，固定填写card
	Title      string `json:"title"`      // 卡片标题，只支持普通文本，可以为空字符串
	Theme      string `json:"theme"`      // 卡片风格，grey，red，orange，yellow ，green，indigo，blue，purple，black，default
	Components []any  `json:"components"` // 内容组件
}

// MessageResponse 消息Api响应
type MessageResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    MessageResponseData `json:"data"`
}

// MessageResponseData 消息Api响应数据
type MessageResponseData struct {
	MessageId string `json:"messageId"`
}

type ImageCard struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

type ImageGroupCard struct {
	Type     string      `json:"type"`
	Elements []ImageCard `json:"elements"`
}

type TextCard struct {
	Type string   `json:"type"`
	Text TextData `json:"text"`
}

type TextData struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// RemarkCard 备注卡片
type RemarkCard struct {
	Type     string           `json:"type"`     // 组件类型，当前填写remark
	Elements []RemarkCardData `json:"elements"` // 数据列表
}

// RemarkCardData 备注数据列表
type RemarkCardData struct {
	Type    string `json:"type"`              // 数据类型，image：图片，plain-text：普通文本，dodo-md：markdown文本
	Content string `json:"content,omitempty"` // 文本内容，数据类型为文本时必填
	Src     string `json:"src,omitempty"`     // 图片地址，一般用作备注头部图标显示，数据类型为图片时必填
}

// TextBody 消息内容
type TextBody struct {
	Content string `json:"content"` // 文字内容，限制10000个字符，频道文字消息支持多种 消息语法
}

// CountdownCard 倒计时卡片
type CountdownCard struct {
	Type    string `json:"type"`    // 组件类型，当前填写countdown
	Style   string `json:"style"`   // 显示样式，day：按天显示，hour ：按小时显示
	Title   string `json:"title"`   // 倒计时标题
	EndTime int64  `json:"endTime"` // 结束时间戳，单位秒
}
