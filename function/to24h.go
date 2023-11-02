package function

import "time"

// To24H 将时间戳转换为24小时制
func To24H(timestamp int64) string {
	twentyFourHour := time.Unix(timestamp, 0)
	return twentyFourHour.Format("2006-01-02 15:04:05")
}

// To24h 将时间戳转换为24小时制
func To24h(times float64) string {
	timestamp := int64(times)
	ctime := time.Unix(timestamp, 0)
	return ctime.Format("2006-01-02 15:04:05")
}
