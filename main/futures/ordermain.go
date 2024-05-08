package main

import (
	"fmt"
	"myapp/future"
)

func main() {
	for i := 0; i < 10; i++ {
		_, openLong, err := future.Order(132, 1, 1, 2, "63999.8", "BTC_USDT", "78")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(openLong)
		_, openShort, err := future.Order(122, 1, 3, 2, "63999.8", "BTC_USDT", "78")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(openShort)

		_, closeLong, err := future.Order(132, 1, 4, 2, "63999.8", "BTC_USDT", "45")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(closeLong)

		_, closeShort, err := future.Order(122, 1, 2, 2, "63999.8", "BTC_USDT", "65")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(closeShort)
	}
}
