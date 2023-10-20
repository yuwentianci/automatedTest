package function

import "fmt"

// FormatPercentage 格式百分比
func FormatPercentage(value float64) string {
	return fmt.Sprintf("%.2f%%", value*100)
}
