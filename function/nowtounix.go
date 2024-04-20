package function

import "time"

func NowToUnix() float64 {
	// 获取当前时间
	currentTime := time.Now()

	// 将时间转换为Unix时间戳（以秒为单位）
	unixTimestamp := currentTime.Unix()

	// 获取纳秒部分并转换为浮点数
	nanoSeconds := float64(currentTime.Nanosecond()) / 1e9

	// 将秒和纳秒部分组合在一起
	result := float64(unixTimestamp) + nanoSeconds
	return result
}
