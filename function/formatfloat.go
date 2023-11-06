package function

import (
	"github.com/shopspring/decimal"
)

// FormatFloat 将string类型转换为float类型
func FormatFloat(data string) float64 {
	//float, err := strconv.ParseFloat(data, 64)
	//if err != nil {
	//	return -1
	//}
	// 如果 data 是科学计数法，将格式化为 decimal 类型
	dec, err := decimal.NewFromString(data)
	if err != nil {
		return -1
	}
	float, _ := dec.Float64()
	return float
}
