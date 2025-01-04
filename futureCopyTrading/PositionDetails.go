package futureCopyTrading

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/futureCopyTrading"
	"myapp/function"
)

type DaiDanCurrentPositionResponse struct {
	Code  int                       `json:"code"`
	Msg   string                    `json:"msg"`
	Data  DaiDanCurrentPositionData `json:"data"`
	Extra *interface{}              `json:"extra"` // 可以根据需要更改类型
}

type DaiDanCurrentPositionData struct {
	List  []DaiDanCurrentPositionInfo `json:"list"`
	Total int                         `json:"total"`
}

type DaiDanCurrentPositionInfo struct {
	ClosePrice     string          `json:"close_price"`
	ContractID     int             `json:"contract_id"`
	ContractName   string          `json:"contract_name"`
	CreatedAt      int64           `json:"created_at"`
	EntryPrice     string          `json:"entry_price"`
	InitialMargin  string          `json:"initial_margin"`
	Leverage       int             `json:"leverage"`
	LiquidatePrice string          `json:"liquidate_price"`
	Margin         string          `json:"margin"`
	MarginMode     int             `json:"margin_mode"`
	PositionID     int             `json:"position_id"`
	PositionType   int             `json:"position_type"`
	ReleasePNL     string          `json:"release_pnl"`
	Size           decimal.Decimal `json:"size"`
	UID            int             `json:"uid"`
	UpdatedAt      int64           `json:"updated_at"`
}

// DaiDanCurrentPosition 带单者当前仓位信息
func DaiDanCurrentPosition(pageSize, trader int) (error, []*DaiDanCurrentPositionInfo) {
	var positions []*DaiDanCurrentPositionInfo
	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spage_index=%d&page_size=%d&trader=%d", futureCopyTrading.DaiDanCurrentPositionUrl, currentPage, pageSize, trader)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var daiDanCurrentPositionResponse DaiDanCurrentPositionResponse
		if err := json.Unmarshal(responseTest, &daiDanCurrentPositionResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
		}

		// 检查响应是否成功
		if daiDanCurrentPositionResponse.Code == 200 {
			if daiDanCurrentPositionResponse.Data.List != nil {
				for _, position := range daiDanCurrentPositionResponse.Data.List {
					newPositions := position
					positions = append(positions, &newPositions)
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < daiDanCurrentPositionResponse.Data.Total {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), nil
		}
	}

	return nil, positions
}

type PositionResponse struct {
	Code  int          `json:"code"`
	Msg   string       `json:"msg"`
	Data  PositionData `json:"data"`
	Extra interface{}  `json:"extra"`
}

type PositionData struct {
	List  []PositionInfo `json:"list"`
	Total int            `json:"total"`
}

type PositionInfo struct {
	ClosePrice   string          `json:"close_price"`
	CloseVolume  decimal.Decimal `json:"close_vol"`
	ContractID   int             `json:"contract_id"`
	ContractName string          `json:"contract_name"`
	CreatedAt    int64           `json:"created_at"`
	EntryPrice   string          `json:"entry_price"`
	Leverage     decimal.Decimal `json:"leverage"`
	MarginMode   int             `json:"margin_mode"`
	PositionID   int             `json:"position_id"`
	PositionType int             `json:"position_type"`
	ReleasePNL   string          `json:"release_pnl"`
	ROI          string          `json:"roi"`
	UID          int             `json:"uid"`
	UpdatedAt    int64           `json:"updated_at"`
	WithHolding  string          `json:"withholding_funding"`
}

// DaiDanHistoryPosition 带单者历史仓位信息
func DaiDanHistoryPosition(pageSize, trader, positionID int) (error, []*PositionInfo) {
	var positions []*PositionInfo
	currentPage := 1

	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spage_index=%d&page_size=%d&trader=%d", futureCopyTrading.DaiDanHistoryPositionUrl, currentPage, pageSize, trader)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			return fmt.Errorf("请求失败: %w", err), nil
		}

		var positionResponse PositionResponse
		if err := json.Unmarshal(responseTest, &positionResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
		}

		// 检查响应是否成功
		if positionResponse.Code == 200 {
			if positionResponse.Data.List != nil {
				for _, position := range positionResponse.Data.List {
					// 检查是否为目标 PositionID
					if positionID == 0 || position.PositionID == positionID {
						newPositions := position
						positions = append(positions, &newPositions)
					}
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < positionResponse.Data.Total {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), nil
		}
	}

	return nil, positions
}

// GenDanHistoryPosition 跟单者某一仓位ID历史仓位信息
func GenDanHistoryPosition(pageSize, trader, positionID int) (error, []*PositionInfo) {
	var positions []*PositionInfo
	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spage_index=%d&page_size=%d&trader=%d", futureCopyTrading.GenDanHistoryPositionUrl, currentPage, pageSize, trader)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var positionResponse PositionResponse
		if err := json.Unmarshal(responseTest, &positionResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
		}

		// 检查响应是否成功
		if positionResponse.Code == 200 {
			if positionResponse.Data.List != nil {
				for _, position := range positionResponse.Data.List {
					// 检查是否为目标 PositionID
					if position.PositionID == 0 || position.PositionID == positionID {
						newPositions := position
						positions = append(positions, &newPositions)
					}
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < positionResponse.Data.Total {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), nil
		}
	}

	return nil, positions
}

type GenDanPositionDetailsResponse struct {
	Code int                       `json:"code"`
	Msg  string                    `json:"msg"`
	Data GenDanPositionDetailsData `json:"data"`
}

type GenDanPositionDetailsData struct {
	List       []GenDanPositionDetailsInfo `json:"list"`
	Total      int                         `json:"total"`
	UsedMargin string                      `json:"used_margin"`
}

type GenDanPositionDetailsInfo struct {
	ContractID     int             `json:"contract_id"`
	ContractName   string          `json:"contract_name"`
	CreatedAt      int64           `json:"created_at"`
	EntryPrice     string          `json:"entry_price"`
	InitialMargin  string          `json:"initial_margin"`
	Leverage       decimal.Decimal `json:"leverage"`
	LiquidatePrice string          `json:"liquidate_price"`
	Margin         string          `json:"margin"`
	MarginMode     int             `json:"margin_mode"`
	PositionID     int             `json:"position_id"`
	PositionType   int             `json:"position_type"`
	ReleasePNL     string          `json:"release_pnl"`
	Size           decimal.Decimal `json:"size"`
}

// GenDanPositionDetails 跟单者某一带单者仓位明细
func GenDanPositionDetails(pageSize, trader, positionID int) (error, []*GenDanPositionDetailsInfo) {
	var positions []*GenDanPositionDetailsInfo
	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spage_index=%d&page_size=%d&trader=%d", futureCopyTrading.GenDanPositionDetailsUrl, currentPage, pageSize, trader)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var genDanPositionDetailsResponse GenDanPositionDetailsResponse
		if err := json.Unmarshal(responseTest, &genDanPositionDetailsResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
		}

		// 检查响应是否成功
		if genDanPositionDetailsResponse.Code == 200 {
			if genDanPositionDetailsResponse.Data.List != nil {
				for _, position := range genDanPositionDetailsResponse.Data.List {
					// 检查是否为目标 PositionID
					if positionID == 0 || position.PositionID == positionID {
						newPositions := position
						positions = append(positions, &newPositions)
					}
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < genDanPositionDetailsResponse.Data.Total {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), nil
		}
	}

	return nil, positions
}

// GenDanTotalPosition 跟单者所有仓位明细

//func GenDanTotalPosition() {
//	responseTest, err := function.GetDetails(futureCopyTrading.GenDanAllPositionUrl)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//
//}

// GenDanHistoryAllPosition 跟单者所有历史仓位信息
func GenDanHistoryAllPosition(pageSize, trader int) (error, []*PositionInfo) {
	var positions []*PositionInfo
	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spage_index=%d&page_size=%d&trader=%d", futureCopyTrading.GenDanHistoryPositionUrl, currentPage, pageSize, trader)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var positionResponse PositionResponse
		if err := json.Unmarshal(responseTest, &positionResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), nil
		}

		// 检查响应是否成功
		if positionResponse.Code == 200 {
			if positionResponse.Data.List != nil {
				for _, position := range positionResponse.Data.List {
					newPositions := position
					positions = append(positions, &newPositions)
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < positionResponse.Data.Total {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), nil
		}
	}

	return nil, positions
}
