package common

// Request 请求数据
type Request struct {
	ClientId string `json:"clientId"` // 机器人唯一标识
	Payload  string `json:"payload"`  // 事件消息，默认进行了AES加密，需要进行数据解密
}

// Response 返回数据
type Response struct {
	Status  int         `json:"status"`         // 返回码，0：成功，-9999：失败
	Message string      `json:"message"`        // 返回信息，一般作为错误信息
	Data    interface{} `json:"data,omitempty"` // 返回数据
}
