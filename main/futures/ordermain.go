package main

import (
	"fmt"
	"myapp/future"
)

func main() {
	for i := 0; i < 10; i++ {
		_, f, err := future.Order(100, 1, 1, 1, "65887.1", "BTC_USDT", "1")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(f)
	}
}
