package message

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
