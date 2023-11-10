package translate

import "strings"

func RankNameZh(rank string) string {
	switch strings.ToUpper(rank) {
	case "ROOKIE":
		return "菜鸟"
	case "BRONZE":
		return "青铜"
	case "SILVER":
		return "白银"
	case "GOLD":
		return "黄金"
	case "PLATINUM":
		return "白金"
	case "DIAMOND":
		return "钻石"
	case "MASTER":
		return "大师"
	case "APEX PREDATOR":
		return "猎杀"
	case "PREDATOR":
		return "猎杀"
	default:
		return "未知"
	}
}
