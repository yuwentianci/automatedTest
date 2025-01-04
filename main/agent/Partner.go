package main

import (
	"fmt"
	"myapp/agentBackend"
)

func main() {
	pageSize, uid, inviterUid, startTime, endTime, startTimeLocal, endTimeLocal, level, field := 50, 0, 0, 1731283200, 1731628799, 1731254400, 1731599999, 1, "total_perpetual_vol"
	err, totalDeposits, totalFees, totalTraders, totalWithdraws, totalTransactionAmount := agentBackend.Partners(pageSize, uid, inviterUid, startTime, endTime, startTimeLocal, endTimeLocal, level, field)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("总充值:", totalDeposits, "总手续费:", totalFees)
	fmt.Println("总交易人数:", totalTraders, "总提现:", totalWithdraws, "总交易额:", totalTransactionAmount)
}
