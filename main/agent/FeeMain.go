package main

import (
	"fmt"
	"myapp/agentBackend"
)

func main() {
	err, fee := agentBackend.CalcFee()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fee)
}
