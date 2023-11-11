package utils

import (
	"fengqi/dodo-apex-robot/logger"
	"go.uber.org/zap"
	"time"
)

func GetLocation() *time.Location {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logger.Client.Error("load location err", zap.Error(err))
	}
	return location
}

func TimestampFormat(timestamp int64) string {
	return time.Unix(timestamp, 0).In(GetLocation()).Format("2006-01-02 15:04:05")
}

func ParseTimestamp(timestamp int64) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", TimestampFormat(timestamp), GetLocation())
}

func ParseTimeString(timeString string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", timeString, GetLocation())
}

func TimeNowString() string {
	return time.Now().In(GetLocation()).Format("2006-01-02 15:04:05")
}
