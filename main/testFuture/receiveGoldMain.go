package main

import (
	"fmt"
	"myapp/future"
)

func main() {
	err, receiveGold := future.ReceiveGold()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(receiveGold)
}
