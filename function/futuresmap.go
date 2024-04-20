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

// FuturesTypeMap 获取开平仓类型文字描述
func FuturesTypeMap(FuturesType int) string {
	FuturesTypes := map[int]string{
		1: "限价",
		5: "市价",
	}
	if TypeStr, ok := FuturesTypes[FuturesType]; ok {
		return TypeStr
	}
	return "未知"
}

// FuturesOperaTypeMap 获取仓位方向文字描述
func FuturesOperaTypeMap(FuturesOperaType int) string {
	FuturesOperaTypes := map[int]string{
		1: "开",
		2: "平",
	}
	if OperaTypeStr, ok := FuturesOperaTypes[FuturesOperaType]; ok {
		return OperaTypeStr
	}
	return "未知"
}

// FuturesSideMap 获取仓位方向文字描述
func FuturesSideMap(FuturesSide int) string {
	FuturesSides := map[int]string{
		1: "开多",
		2: "平空",
		3: "开空",
		4: "平多",
	}
	if sideStr, ok := FuturesSides[FuturesSide]; ok {
		return sideStr
	}
	return "未知"
}
