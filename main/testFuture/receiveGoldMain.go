package main

import (
	"fmt"
	"myapp/testFuture"
)

func main() {
	err, receiveGold := testFuture.ReceiveGold()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(receiveGold)
}
