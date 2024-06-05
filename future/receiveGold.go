package future

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func ReceiveGold() (error, decimal.Decimal) {
	err, totalPositionMargin := CalcPositionMargin()
	if err != nil {
		fmt.Println(err)
	}

	err, totalOrderMargin := CalcOrderMargin()
	if err != nil {
		fmt.Println(err)
	}

	err, avlBal := AssetsData()
	if err != nil {
		fmt.Println(err)
	}

	totalGold := decimal.NewFromFloat(50000.0)
	receiveGold := totalGold.Sub(totalPositionMargin).Sub(totalOrderMargin).Sub(avlBal)
	return nil, receiveGold
}
