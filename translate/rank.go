package translate

import "strings"

var RankMap = map[string]string{
	"ROOKIE":        "菜鸟",
	"BRONZE":        "青铜",
	"SILVER":        "白银",
	"GOLD":          "黄金",
	"PLATINUM":      "白金",
	"DIAMOND":       "钻石",
	"MASTER":        "大师",
	"APEX PREDATOR": "猎杀",
	"PREDATOR":      "猎杀",
}

func RankNameZh(rank string) string {
	if zh, ok := RankMap[strings.ToUpper(rank)]; ok {
		return zh
	}
	return rank
}
