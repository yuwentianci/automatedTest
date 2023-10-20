package function

import (
	"encoding/json"
	"fmt"
)

// ParseJsonRe 解析JSON响应或处理raw格式响应，具体取决于API的响应类型
func ParseJsonRe(responseText []byte, details interface{}) error {
	if err := json.Unmarshal(responseText, details); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return err
	}
	return nil
}
