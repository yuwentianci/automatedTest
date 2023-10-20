package function

import "time"

// ConvertTo24HourFormat 将时间戳转换为24小时制
func ConvertTo24HourFormat(timestamp int64) string {
	twentyFourHour := time.Unix(timestamp, 0)
	return twentyFourHour.Format("2006-01-02 15:04:05")
}
