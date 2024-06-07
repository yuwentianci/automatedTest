package main

import "myapp/earn"

func main() {
	//BTC 0.00211
	//amount1 := decimal.NewFromFloat(0.0009)
	//amount2 := decimal.NewFromFloat(0.0015)
	//earn.BuyEarn("104", amount1, amount2, 0.0001)
	//USDT 304.6202
	//amount3 := decimal.NewFromFloat(100.0)
	//amount4 := decimal.NewFromFloat(120.0)
	//earn.BuyEarn("95", amount3, amount4, 1)
	//BIT 3434592.89
	//amount5 := decimal.NewFromFloat(1500001.0)
	//amount6 := decimal.NewFromFloat(1500005.0)
	//earn.BuyEarn("88", amount5, amount6, 1)
	//BONE 204.6472
	//amount7 := decimal.NewFromFloat(100.0)
	//amount8 := decimal.NewFromFloat(100.5)
	//earn.BuyEarn("82", amount7, amount8, 0.1)
	//ETH 0.024445
	//amount9 := decimal.NewFromFloat(0.01)
	//amount10 := decimal.NewFromFloat(0.05)
	//earn.BuyEarn("23", amount9, amount10, 0.01)

	//earn.AllEarnProduct(1, 10, "BONE")
	//earn.MyFixedEarn(10, "")

	//理财购买历史记录
	earn.Profit(10)
}
