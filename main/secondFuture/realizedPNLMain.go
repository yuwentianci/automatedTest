package main

import (
	"fmt"
	"myapp/secondFuture"
)

func main() {
	err, profit := secondFuture.RealizedPNL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("profit:", profit)
}
