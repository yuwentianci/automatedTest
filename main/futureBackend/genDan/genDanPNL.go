package main

import (
	"fmt"
	"myapp/futureBackend/genDan"
)

func main() {
	err, totalPNL := genDan.GenDanPNL(10, 51300844)
	if err != nil {
		fmt.Println("错误信息:", err)
	}
	fmt.Println("总盈亏：", totalPNL)
}
