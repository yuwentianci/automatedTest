package futureBackend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/futureBackend"
	"myapp/function"
)

type TradingVolumeResponse struct {
	Code    int               `json:"code"`
	Data    TradingVolumeData `json:"data"`
	Message string            `json:"msg"`
	Success bool              `json:"success"`
}

type TradingVolumeData struct {
	ResultList  []TradingVolumeInfo `json:"resultList"`
	TotalResult int                 `json:"totalResult"`
}

type TradingVolumeInfo struct {
	ID         int64  `json:"id"`
	RoundID    int64  `json:"round_id"`
	Underlying string `json:"underlying"`
	UID        int64  `json:"uid"`
	Amount     string `json:"amount"`
	Side       int    `json:"side"`
	PnL        string `json:"pnl"`
	Status     int    `json:"status"`
	EndTime    int64  `json:"end_time"`
	Fee        string `json:"fee"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func CalcTradingVolume(uid, pageSize int, startTime, endTime int64) (error, decimal.Decimal) {

	tradingVolume := decimal.Zero
	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spageIndex=%d&pageSize=%d&uid=%d", futureBackend.TradingVolumeUrl, currentPage, pageSize, uid)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var tradingVolumeResponse TradingVolumeResponse
		if err := json.Unmarshal(responseTest, &tradingVolumeResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
		}

		// 检查响应是否成功
		if tradingVolumeResponse.Code == 200 {
			// 遍历结果列表
			for _, item := range tradingVolumeResponse.Data.ResultList {
				if item.Status == 1 && item.EndTime >= startTime && item.EndTime < endTime && item.PnL != "0.000000000000000000" {
					amountDec, err := decimal.NewFromString(item.Amount)
					if err != nil {
						return errors.New("金额转换错误:" + err.Error()), decimal.Zero
					}

					tradingVolume = tradingVolume.Add(amountDec)
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < tradingVolumeResponse.Data.TotalResult {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), decimal.Zero
		}
	}

	return nil, tradingVolume
}

func SecondPNL(uid, pageSize, roundId int, startTime, endTime int64) (error, decimal.Decimal, decimal.Decimal) {
	var (
		risePnl         = decimal.Zero
		fallPnl         = decimal.Zero
		amountRiseUser  = decimal.Zero
		amountFallUser  = decimal.Zero
		totalAmountRise = decimal.Zero
		totalAmountFall = decimal.Zero
		currentPage     = 1
	)

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spageIndex=%d&pageSize=%d&roundId=%d&startTime=%d&endTime=%d",
			futureBackend.TradingVolumeUrl, currentPage, pageSize, roundId, startTime, endTime)

		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			return fmt.Errorf("请求错误: %w", err), decimal.Zero, decimal.Zero
		}

		var tradingVolumeResponse TradingVolumeResponse
		if err := json.Unmarshal(responseTest, &tradingVolumeResponse); err != nil {
			return fmt.Errorf("解析JSON响应时发生错误: %w", err), decimal.Zero, decimal.Zero
		}

		// 检查响应是否成功
		if tradingVolumeResponse.Code != 200 {
			return errors.New("请求失败"), decimal.Zero, decimal.Zero
		}

		// 遍历结果列表
		for _, item := range tradingVolumeResponse.Data.ResultList {
			amountDec, err := decimal.NewFromString(item.Amount)
			if err != nil {
				return fmt.Errorf("金额转换错误: %w", err), decimal.Zero, decimal.Zero
			}

			if item.UID == int64(uid) {
				if item.Side == 0 {
					amountRiseUser = amountRiseUser.Add(amountDec)
				} else {
					amountFallUser = amountFallUser.Add(amountDec)
				}
			}

			if item.Side == 0 {
				totalAmountRise = totalAmountRise.Add(amountDec)
			} else {
				totalAmountFall = totalAmountFall.Add(amountDec)
			}
		}

		// 检查是否还有更多的页面
		totalCurrentPage := pageSize * currentPage
		if totalCurrentPage < tradingVolumeResponse.Data.TotalResult {
			currentPage++
		} else {
			break
		}
	}

	if !totalAmountRise.IsZero() {
		risePnl = amountRiseUser.Div(totalAmountRise).Mul(totalAmountFall)
	}
	if !totalAmountFall.IsZero() {
		fallPnl = amountFallUser.Div(totalAmountFall).Mul(totalAmountRise)
	}

	return nil, risePnl, fallPnl
}

type BackendPositionResponse struct {
	Success bool                `json:"success"`
	Code    int                 `json:"code"`
	Msg     string              `json:"msg"`
	Data    BackendPositionData `json:"data"`
}

type BackendPositionData struct {
	CurrentPage int                   `json:"currentPage"`
	ShowCount   int                   `json:"showCount"`
	ResultList  []BackendPositionInfo `json:"resultList"`
	TotalResult int                   `json:"totalResult"`
}

type BackendPositionInfo struct {
	ID                int64  `json:"id"`
	UID               int64  `json:"uid"`
	ContractID        int    `json:"contractId"`
	PositionType      string `json:"positionType"`
	OpenType          string `json:"openType"`
	State             string `json:"state"`
	OpenAvgPrice      string `json:"openAvgPrice"`
	CloseAvgPrice     string `json:"closeAvgPrice"`
	CloseVol          string `json:"closeVol"`
	HoldVol           string `json:"holdVol"`
	AvailableCloseVol string `json:"availableCloseVol"`
	Leverage          int    `json:"leverage"`
	FairPrice         string `json:"fairPrice"`
	Oim               string `json:"oim"`
	Im                string `json:"im"`
	Mm                string `json:"mm"`
	LiquidatePrice    string `json:"liquidatePrice"`
	BankruptcyPrice   string `json:"bankruptcyPrice"`
	MarginRatio       string `json:"marginRatio"`
	FloatingPL        string `json:"floatingPL"`
	Profit            string `json:"profit"`
	PositionCost      string `json:"positionCost"`
	ReturnRate        string `json:"returnRate"`
	Realised          string `json:"realised"`
	AutoAddIm         bool   `json:"autoAddIm"`
	CreateTime        int64  `json:"createTime"`
	UpdateTime        int64  `json:"updateTime"`
}

// CalcPositionVolume 统计用户当前的总持仓量
func CalcPositionVolume() (error, decimal.Decimal) {
	totalHoldVol := decimal.Zero
	// 构造请求体
	requestData := map[string]interface{}{
		"pageIndex":     1,
		"pageSize":      100,
		"isMarketMaker": "false",
		"state":         "HOLDING",
		"contractIds":   []int{10},
		"type":          "LONG",
	}

	// 将数据序列化为 JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("JSON 编码失败:", err)
		return err, decimal.Zero
	}

	// 创建一个POST请求，将rawBytes作为请求主体数据
	responseText, err := function.PostByteDetails(futureBackend.PositionListUrl, jsonData)
	if err != nil {
		return fmt.Errorf("POST 请求失败: %w", err), decimal.Zero
	}

	// 解析JSON响应或处理raw格式响应，具体取决于API的响应类型
	var backendPositionResponse BackendPositionResponse
	if err := json.Unmarshal(responseText, &backendPositionResponse); err != nil {
		return fmt.Errorf("解析JSON响应时发生错误: %w", err), decimal.Zero
	}

	for _, item := range backendPositionResponse.Data.ResultList {
		amountDec, err := decimal.NewFromString(item.HoldVol)
		if err != nil {
			return fmt.Errorf("金额转换错误: %w", err), decimal.Zero
		}
		totalHoldVol = totalHoldVol.Add(amountDec.Mul(decimal.NewFromFloat32(0.0001)))
	}

	return nil, totalHoldVol
}
