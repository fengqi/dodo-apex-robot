package translate

import (
	"fengqi/dodo-apex-robot/logger"
	"go.uber.org/zap"
)

func MapNameZh(name string) string {
	switch name {
	case "Kings Canyon":
		return "国王峡谷"

	case "World's Edge":
		return "世界边缘"

	case "Broken Moon":
		return "残月"

	case "Olympus":
		return "奥林匹斯"

	case "Overflow":
		return "溢出"

	case "Phase runner":
		return "相位穿梭器"

	case "Storm Point":
		return "风暴点"

	case "Estates":
		return "不动产"

	case "Barometer":
		return "气压计"

	default:
		logger.Client.Warn("unknown map name", zap.String("name", name))
		return name
	}
}
