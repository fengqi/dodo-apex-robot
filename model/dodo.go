package model

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
