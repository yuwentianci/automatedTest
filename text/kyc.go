package text

import (
	"fmt"
	"myapp/function"
	"myapp/redata"
)

type Identification struct {
}

func (t *Identification) Identity() {

	//创建要发送的 JSON 数据
	formData := map[string]interface{}{
		"number":     "14018112",
		"last_name":  "树的影",
		"type":       "id",
		"first_name": "人的名",
		"country_id": "44",
	}

	// 创建一个POST请求
	url := "https://www.biconomy.com/api/v1/user/identity"
	// 解析JSON响应或处理raw格式响应，具体取决于API的响应类型
	var Identity = redata.Identity{}
	if err := function.PostFormData(url, formData, &Identity); err != nil {
		fmt.Println(err.Error())
		return
	}
	if Identity.Code == 0 {
		println("method:", Identity.Result.Method, "id:", Identity.Result.ID, "url:", Identity.Result.Url)
	} else {
		println("身份认证信息提交失败，失败原因:", Identity.Message)
	}
}
