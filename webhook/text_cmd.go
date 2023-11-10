package webhook

import (
	"regexp"
	"strings"
)

const (
	CmdUser  = ".apex u"
	CmdMap   = ".apex m"
	CmdShop  = ".apex s"
	CmdCraft = ".apex c"
	CmdTime  = ".apex t"
	CmdHelp  = ".apex h"
	CmdDist  = ".apex d"
	CmdPick  = ".apex p"
)

//var playerMatch, _ = regexp.Compile(`^(.|。)apex\s+u\s+(.*)$`)

var cmdMap = map[string]string{
	CmdUser:  "查询玩家信息, 用法: .apex u EA用户名",
	CmdMap:   "查询地图",
	CmdCraft: "查询复制器",
	CmdShop:  "~~查询商店~~",
	CmdTime:  "~~查询赛季时间~~",
	CmdDist:  "~~查询段位分布~~",
	CmdPick:  "~~查询英雄选择率~~",
	CmdHelp:  "显示帮助信息",
}

// ClearCmd 清理命令, 去除前缀空格, 全角句号, 多余空格
// 限定命令最小长度为5
func ClearCmd(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if len(cmd) < 5 {
		return ""
	}

	if cmd[0:3] == "。" {
		cmd = "." + cmd[3:]
	}

	re := regexp.MustCompile(`\s+`)
	cmd = re.ReplaceAllString(cmd, " ")

	if len(cmd) < 5 {
		return ""
	}

	return cmd
}

func ParseCmd(cmd string) string {
	cmd = ClearCmd(cmd)
	if cmd == "" {
		return ""
	}

	for k, _ := range cmdMap {
		if len(cmd) >= 7 && k == cmd[0:7] {
			return k
		}
	}

	if cmd[0:5] == ".apex" { // 对于不匹配的全部返回帮助
		return CmdHelp
	}

	return ""
}
