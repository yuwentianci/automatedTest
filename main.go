package main

import (
	_ "github.com/go-sql-driver/mysql"
	"myapp/text"
)

func main() {
	//a := new(text.PostSpotCreateData)
	//a.LimitBuy()
	//a.MarketBuy()
	//a.LimitSell()
	ee := new(text.PostEarnCreateData)
	//ee.MyEarnAssetsDetails(ee.MyEarn(1, 2, "BTC"))
	//ee.EarnProduct(0, 10, "")
	ee.Profit()
}
