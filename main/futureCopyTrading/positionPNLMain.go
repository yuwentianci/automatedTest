package main

import (
	"fmt"
	"myapp/futureCopyTrading"
)

func main() {
	//trader := 51404556
	//err, danDanHistoryPosition := futureCopyTrading.DaiDanHistoryPosition(10, 51373300, 26920)
	//if err != nil {
	//	fmt.Println("带单者历史仓位错误信息：", err)
	//}

	// 获取交易对详情
	//err, symbol := future.SymbolDetails()
	//if err != nil {
	//	fmt.Println("symbol错误信息：", err)
	//	return
	//}
	// 获取杠杆详情
	//err, lvg := future.LvgMmrDetails()
	//if err != nil {
	//	fmt.Println("杠杆错误信息：", err)
	//}

	// 计算带单者历史仓位盈亏和盈亏率
	//err, danDanHistoryPositionPNL := futureCopyTrading.DaiDanHistoryPositionPNL(danDanHistoryPosition, symbol)
	//if err != nil {
	//	fmt.Println("pnl错误信息", err)
	//}
	//fmt.Println("带单者历史仓位盈亏:", danDanHistoryPositionPNL)
	//
	//err, daiDanHistoryPositionPNLRoi := futureCopyTrading.DaiDanHistoryPositionPNLRoi(danDanHistoryPosition, symbol, lvg)
	//if err != nil {
	//	fmt.Println("带单者历史仓位盈亏率错误：", err)
	//}
	//fmt.Println("带单者历史仓位盈亏率：", daiDanHistoryPositionPNLRoi)

	// 计算带单者总已实现盈亏
	//err, daiDanCurrentPosition := futureCopyTrading.DaiDanCurrentPosition(10, trader)
	//if err != nil {
	//	fmt.Println("带单者当前仓位报错信息:", err)
	//}
	//
	//err, daiDanCurrentPositionFee := futureCopyTrading.DaiDanCurrentPositionPNL(daiDanCurrentPosition, symbol)
	//if err != nil {
	//	fmt.Println("带单者当前仓位手续费报错信息:", err)
	//}
	//fmt.Println(daiDanCurrentPositionFee)
	//err, daiDanHistoryPositionTotalPNL := futureCopyTrading.DaiDanHistoryPositionTotalPNL(10, trader, 0)
	//if err != nil {
	//	fmt.Println("带单者历史仓位总盈亏错误信息:", err)
	//}
	//fmt.Println("手续费:", daiDanCurrentPositionFee, "历史仓位总盈亏：", daiDanHistoryPositionTotalPNL)
	//total := daiDanHistoryPositionTotalPNL.Sub(daiDanCurrentPositionFee)
	//fmt.Println("带单员总盈亏:", total)

	//计算跟单者历史仓位盈亏和盈亏率
	//err, ganDanHistoryPosition := futureCopyTrading.GenDanHistoryPosition(10, 51404556, 2715)
	//if err != nil {
	//	fmt.Println("错误信息：", err)
	//}
	//err, ganDanHistoryPositionPNL := futureCopyTrading.GenDanHistoryPositionPNL(ganDanHistoryPosition, symbol)
	//if err != nil {
	//	fmt.Println("跟单者历史仓位盈亏错误信息：", err)
	//}
	//fmt.Println("跟单者历史仓位盈亏:", ganDanHistoryPositionPNL)
	//err, genDanHistoryPositionPNLRoi := futureCopyTrading.GenDanHistoryPositionPNLRoi(ganDanHistoryPosition, symbol)
	//if err != nil {
	//	fmt.Println("跟单者历史仓位盈亏率错误信息：", err)
	//}
	//fmt.Println("跟单者历史仓位盈亏率：", genDanHistoryPositionPNLRoi)

	err, ganDanHistoryAllPosition := futureCopyTrading.GenDanHistoryAllPosition(10, 51404556)
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	err, totalPNL := futureCopyTrading.GenDanHistoryPositionTotalPNL(ganDanHistoryAllPosition)
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	fmt.Println("跟单者历史仓位总盈亏：", totalPNL)
}
