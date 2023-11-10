package webhook

import (
	"regexp"
	"strings"
)

const (
	CmdUser      = ".apex u"
	CmdUserName  = "user"
	CmdMap       = ".apex m"
	CmdMapName   = "map"
	CmdShop      = ".apex s"
	CmdShopName  = "shop"
	CmdCraft     = ".apex c"
	CmdCraftName = "craft"
	CmdTime      = ".apex t"
	CmdTimeName  = "time"
	CmdHelp      = ".apex h"
	CmdHelpName  = "help"
	CmdDist      = ".apex d"
	CmdDistName  = "dist"
	CmdPick      = ".apex p"
	CmdPickName  = "pick"
)

var playerMatch, _ = regexp.Compile(`^(.|。)apex\s+u\s+(.*)$`)

/*var cmdMap = map[string]string{
	CmdUser:  CmdUserName,  // 玩家信息
	CmdMap:   CmdMapName,   // 地图
	CmdShop:  CmdShopName,  // 商店
	CmdCraft: CmdCraftName, // 复制器
	CmdTime:  CmdTimeName,  // 赛季时间
	CmdDist:  CmdDistName,  // 段位分布
	CmdPick:  CmdPickName,  // 英雄选择率
	CmdHelp:  CmdHelpName,  // 帮助
}*/

var cmdMapZh = map[string]string{
	CmdUser:  "查询玩家信息, 用法: .apex u EA用户名",
	CmdMap:   "查询地图",
	CmdShop:  "查询商店",
	CmdCraft: "查询复制器",
	CmdTime:  "查询赛季时间",
	CmdHelp:  "查询帮助",
	CmdDist:  "查询段位分布",
	CmdPick:  "查询英雄选择率",
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

	for k, _ := range cmdMapZh {
		if len(cmd) >= 7 && k == cmd[0:7] {
			return k
		}
	}

	if cmd[0:5] == ".apex" { // 对于不匹配的全部返回帮助
		return CmdHelp
	}

	return ""
}
