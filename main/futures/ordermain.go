package main

import (
	"fmt"
	"myapp/future"
)

func main() {
	for i := 0; i < 10; i++ {
		_, openLong, err := future.Order(200, 2, 1, 5, "71196.8", "BTC_USDT", "4")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(openLong)
		_, openShort, err := future.Order(84, 1, 3, 5, "71196.8", "BTC_USDT", "3")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(openShort)

		_, closeLong, err := future.Order(200, 2, 4, 5, "71196.8", "BTC_USDT", "2")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(closeLong)

		_, closeShort, err := future.Order(84, 1, 2, 5, "71196.8", "BTC_USDT", "1")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(closeShort)
	}
}
