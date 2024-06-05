package main

import (
	"fmt"
	"myapp/future"
)

func main() {
	err, totalPNL := future.TotalRealizedPNL("TRB_USDT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("totalPNL:", totalPNL)
}
