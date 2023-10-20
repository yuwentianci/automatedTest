package function

// GetTermDescription 获取文字描述
func GetTermDescription(earnType string) string {
	earnTypes := map[string]string{
		"Fixed":          "定期",
		"Flexible":       "活期",
		"Flexible/Fixed": "活期/定期",
	}
	if typeStr, ok := earnTypes[earnType]; ok {
		return typeStr
	}
	return "未知"
}
