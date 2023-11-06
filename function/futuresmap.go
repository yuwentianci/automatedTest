package function

// FuturesPatternMap 获取保证金模式文字描述
func FuturesPatternMap(FuturesPattern int) string {
	FuturesPatterns := map[int]string{
		1: "逐仓",
		2: "全仓",
	}
	if patternStr, ok := FuturesPatterns[FuturesPattern]; ok {
		return patternStr
	}
	return "未知"
}

// FuturesSideMap 获取仓位方向文字描述
func FuturesSideMap(FuturesSide int) string {
	FuturesSides := map[int]string{
		1: "空",
		2: "多",
	}
	if sideStr, ok := FuturesSides[FuturesSide]; ok {
		return sideStr
	}
	return "未知"
}
