package function

import (
	"fmt"
)

// FormatString float类型返回为string类型
func FormatString(value float64) string {
	intValue := int(value)
	if value == float64(intValue) {
		return fmt.Sprintf("%d", intValue)
	} else if value > 1e6 {
		return fmt.Sprintf("%e", value)
	} else {
		return fmt.Sprintf("%f", value)
	}
}
