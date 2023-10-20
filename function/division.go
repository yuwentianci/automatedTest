package function

import (
	"fmt"
)

// Divide 除法运算
func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	result := dividend / divisor
	return result, nil
}
