package als

// Bridge ALS桥接
type Bridge struct {
	Global   Global   `json:"global"`
	Realtime Realtime `json:"realtime"`
}

// MapRotation 地图轮换
type MapRotation struct {
	BattleRoyale BattleMap `json:"battleRoyale"`
	Arenas       BattleMap `json:"arenas"`
	Ranked       BattleMap `json:"ranked"`
	ArenasRanked BattleMap `json:"arenasRanked"`
	Ltm          BattleMap `json:"ltm"`
}

// Global 全局信息
type Global struct {
	Name                string  `json:"name"`
	Uid                 string  `json:"uid"`
	Platform            string  `json:"platform"`
	Avatar              string  `json:"avatar"`
	Level               int     `json:"level"`
	ToNextLevelPercent  float64 `json:"toNextLevelPercent"`
	InternalUpdateCount int     `json:"internalUpdateCount"`
	Bans                Bans    `json:"bans"`
	Rank                Rank    `json:"rank"`
}

// Rank 段位信息
type Rank struct {
	RankScore           int     `json:"rankScore"`
	RankName            string  `json:"rankName"`
	RankDiv             int     `json:"rankDiv"`
	RankImg             string  `json:"rankImg"`
	RankedSeason        string  `json:"rankedSeason"`
	ALStopPercent       float64 `json:"ALStopPercent"`
	ALStopInt           int     `json:"ALStopInt"`
	ALSTopPercentGlobal float64 `json:"ALSTopPercentGlobal"`
	ALSTopIntGlobal     int     `json:"ALSTopIntGlobal"`
	ALSFlag             bool    `json:"ALSFlag"`
}

// Bans 封禁信息
type Bans struct {
	IsActive      bool   `json:"isActive"`
	Remaining     int    `json:"remaining"`
	LastBanReason string `json:"lastBanReason"`
}

// Realtime 实时信息
type Realtime struct {
	IsOnline                  int    `json:"isOnline"`
	LobbyState                string `json:"lobbyState"`
	IsInGame                  int    `json:"isInGame"`
	CanJoin                   int    `json:"canJoin"`
	PartyFull                 int    `json:"partyFull"`
	SelectedLegend            string `json:"selectedLegend"`
	CurrentState              string `json:"currentState"`
	CurrentStatSinceTimestamp int    `json:"currentStatSinceTimestamp"`
	CurrentStateAsText        string `json:"currentStateAsText"`
}

// BattleMap 比赛地图
type BattleMap struct {
	Current Current `json:"current"`
	Next    Next    `json:"next"`
}

// Current 当前地图
type Current struct {
	Start             int64  `json:"start"`
	End               int64  `json:"end"`
	ReadableDateStart string `json:"readableDate_start"`
	ReadableDateEnd   string `json:"readableDate_end"`
	Map               string `json:"map"`
	Code              string `json:"code"`
	DurationInSecs    int    `json:"durationInSecs"`
	DurationInMinutes int    `json:"durationInMinutes"`
	IsActive          bool   `json:"isActive"`  // 仅限LTM
	EventName         string `json:"eventName"` // 仅限LTM
	Asset             string `json:"asset"`
	RemainingSecs     int    `json:"remainingSecs"`
	RemainingMins     int    `json:"remainingMins"`
	RemainingTimer    string `json:"remainingTimer"`
}

// Next 下一地图
type Next struct {
	Start             int64  `json:"start"`
	End               int64  `json:"end"`
	ReadableDateStart string `json:"readableDate_start"`
	ReadableDateEnd   string `json:"readableDate_end"`
	Map               string `json:"map"`
	Code              string `json:"code"`
	DurationInSecs    int    `json:"DurationInSecs"`
	DurationInMinutes int    `json:"DurationInMinutes"`
	IsActive          bool   `json:"isActive"`  // 仅限LTM
	EventName         string `json:"eventName"` // 仅限LTM
	Asset             string `json:"asset"`
}
