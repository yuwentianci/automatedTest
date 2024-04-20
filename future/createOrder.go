package future

import (
	"errors"
	"fmt"
	"myapp/function"
)

const (
	openCloseUrl = "https://api-future.biconomy.com/future/api/v1/private/order/submit"
)

type orderSuccess struct {
	Success bool  `json:"success"`
	Code    int   `json:"code"`
	Data    int64 `json:"data"`
}
type orderFailure struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Order(leverage, openType, sides, types int, price, symbol, amount string) (int, float64, error) {
	// 获取当前时间戳
	currentTime := function.NowToUnix()
	// 构建 rawData
	rawData := orderRawData(leverage, openType, types, sides, price, symbol, amount)
	sideStr := function.FuturesSideMap(sides)
	typeStr := function.FuturesTypeMap(types)

	var orderSuccessData = orderSuccess{}
	if err := function.PostByteDetailsComplete(openCloseUrl, rawData, &orderSuccessData); err != nil {
		fmt.Println(typeStr, sideStr, "失败,失败原因:", err)
		return 1, currentTime, err
	}

	var orderFailureData = orderFailure{}
	if err := function.PostByteDetailsComplete(openCloseUrl, rawData, &orderFailureData); err != nil {
		fmt.Println(typeStr, sideStr, "失败,失败原因:", err)
		return 1, currentTime, err
	}

	if orderSuccessData.Code == 0 {
		fmt.Println(typeStr, sideStr, "挂单成功")
		return 0, currentTime, nil
	}

	return 1, currentTime, errors.New(typeStr + sideStr + "挂单失败:" + orderFailureData.Message)
}

func orderRawData(leverage, openType, types, sides int, price, symbol, amount string) map[string]interface{} {
	rawData := map[string]interface{}{
		"leverage":  leverage,
		"openType":  openType,
		"symbol":    symbol,
		"side":      sides,
		"orderType": types,
		"vol":       amount,
	}

	// 如果 types 等于 1，添加 price 字段
	if types == 1 {
		rawData["price"] = price
	}

	return rawData
}
