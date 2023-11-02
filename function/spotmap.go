package function

// SpotTypeMap 获取买卖文字描述
func SpotSideMap(side int) string {
	spotSide := map[int]string{
		1: "卖出",
		2: "买入",
	}
	if sideStr, ok := spotSide[side]; ok {
		return sideStr
	}
	return "未知"
}

// SpotRoleMap 获取是否做市文字描述
func SpotRoleMap(role string) string {
	spotRole := map[string]string{
		"1": "Taker",
		"2": "Maker",
	}
	if roleStr, ok := spotRole[role]; ok {
		return roleStr
	}
	return "未知"
}
