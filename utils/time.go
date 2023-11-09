package utils

import "time"

func GetLocation() *time.Location {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	return location
}

func TimestampFormat(timestamp int64) string {
	return time.Unix(timestamp, 0).In(GetLocation()).Format("2006-01-02 15:04:05")
}
