package config

type Config struct {
	Dodo dodo               `json:"dodo"`
	ALS  apexLegendsStatus  `json:"apex_legends_status"`
	ALT  apexLegendsTracker `json:"apex_legends_tracker"`
}

type dodo struct {
	ClientId  string `json:"client_id"`  // Client Id（应用唯一标识）
	BotToken  string `json:"bot_token"`  // Bot Token（机器人鉴权Token）
	Host      string `json:"host"`       // Host（机器人鉴权Token）
	SecretKey string `json:"secret_key"` // Secret Key（事件密钥）
}

type apexLegendsStatus struct {
	Host   string `json:"host"`
	ApiKey string `json:"api_key"`
}

type apexLegendsTracker struct {
	Host   string `json:"host"`
	ApiKey string `json:"api_key"`
}
