package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/future"
	"myapp/futureCopyTrading"
)

func main() {
	trader := 51404524
	ganDanCloseBeforeAsset, _ := decimal.NewFromString("16015.48668037604")
	//genDanOpenBeforeAsset, _ := decimal.NewFromString("34485.987493159220295172")
	err, beforeAsset := futureCopyTrading.CopyTradingAsset()
	if err != nil {
		fmt.Println("可用资产报错:", err)
	}
	fmt.Println("操作前资产:", beforeAsset)
	fmt.Println()

	// 获取交易对详情
	err, symbol := future.SymbolDetails()
	if err != nil {
		fmt.Println("symbol错误信息：", err)
	}

	// 开仓后资产变动
	//err, genDanPositionDetails := futureCopyTrading.GenDanPositionDetails(10, trader, 2715)
	//if err != nil {
	//	fmt.Println("跟单者仓位明细错误信息:", err)
	//}
	//err, openAfterAsset := futureCopyTrading.OpenAfterAsset(genDanOpenBeforeAsset, genDanPositionDetails, symbol)
	//if err != nil {
	//	fmt.Println("跟单者开仓后资产报错信息：", err)
	//}
	//fmt.Println("跟单者开仓后资产:", openAfterAsset)
	//fmt.Println()

	// 平仓后资产变动
	err, ganDanHistoryPosition := futureCopyTrading.GenDanHistoryPosition(50, trader, 3112)
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	err, closeAfterAsset := futureCopyTrading.CloseAfterAsset(ganDanCloseBeforeAsset, ganDanHistoryPosition, symbol)
	if err != nil {
		fmt.Println("平仓后资产错误信息:", err)
	}
	fmt.Println("平仓后资产:", closeAfterAsset)
	fmt.Println()

	// 调低杠杆后资产变动
	//err, asset := futureCopyTrading.AdjustLeverageAfterAsset(beforeAsset, 28.0, genDanPositionDetails, symbol)
	//if err != nil {
	//	fmt.Println("错误信息：", err)
	//}
	//fmt.Println("调整杠杆后资产:", asset)
}
