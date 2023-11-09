package utils

import (
	"regexp"
	"strings"
)

var (
	playerMatch *regexp.Regexp
	mapMatch    *regexp.Regexp
	shopMatch   *regexp.Regexp
	craftMatch  *regexp.Regexp
)

func init() {
	// todo 大概率不需要正则
	playerMatch, _ = regexp.Compile(`^(.|。)apex\s+u\s+(.*)$`)
	mapMatch, _ = regexp.Compile(`^(?:.|。)apex\s+m$`)
	shopMatch, _ = regexp.Compile(`^(?:.|。)apex\s+s$`)
	craftMatch, _ = regexp.Compile(`^(?:.|。)apex\s+c$`)
}

// MatchPlayerName 匹配玩家名称
// TODO 对齐ea对用户名的限制
func MatchPlayerName(text string) string {
	text = strings.TrimSpace(text)
	match := playerMatch.FindStringSubmatch(text)
	if len(match) == 3 {
		return match[2]
	}
	return ""
}

// MatchIsMap 是否是地图查询
func MatchIsMap(text string) bool {
	text = strings.TrimSpace(text)
	return mapMatch.MatchString(text)
}

// MatchIsShop 是否是商店查询
func MatchIsShop(text string) bool {
	text = strings.TrimSpace(text)
	return shopMatch.MatchString(text)
}

// MatchIsCraft 是否是合成查询
func MatchIsCraft(text string) bool {
	text = strings.TrimSpace(text)
	return craftMatch.MatchString(text)
}
