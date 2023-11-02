package function

// StepValue 步进值
func StepValue(amount float64) float64 {
	// 初始化步进值
	step := 0.01
	if amount >= 0.0001 && amount < 0.001 {
		step = 0.00001
	} else if amount >= 0.001 && amount < 0.01 {
		step = 0.0001
	} else if amount >= 0.01 && amount < 0.1 {
		step = 0.001
	} else if amount >= 0.1 && amount < 10 {
		step = 0.01
	} else if amount >= 10 {
		step = 1
	}
	return step
}
