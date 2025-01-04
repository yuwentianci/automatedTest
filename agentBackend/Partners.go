package agentBackend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/agentBackend"
	"myapp/function"
	"myapp/future"
)

type PartnersResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  PartnersResult `json:"result"`
}

type PartnersResult struct {
	Item         []PartnersInfo `json:"item"`
	Total        int            `json:"total"`
	Registration int            `json:"registration"`
	PerpetualFee string         `json:"perpetual_fee"`
	PerpetualVol string         `json:"perpetual_vol"`
	Deposited    string         `json:"deposited"`
	Withdrawal   string         `json:"withdrawal"`
}

type PartnersInfo struct {
	UID                  int    `json:"uid"`
	InviterUID           int    `json:"inviter_uid"`
	Nickname             string `json:"nickname"`
	RegistrationTime     string `json:"registration_time"`
	TeamNum              int    `json:"team_num"`
	Level                int    `json:"level"`
	UserType             string `json:"user_type"`
	DepositUser          int    `json:"deposit_user"`
	Deposit              string `json:"deposit"`
	WithdrawUser         int    `json:"withdraw_user"`
	Withdraw             string `json:"withdraw"`
	TotalBalance         string `json:"total_balance"`
	PerpRe               string `json:"perp_re"`
	Profit               string `json:"profit"`
	Loss                 string `json:"loss"`
	PersonalPerpetualFee string `json:"personal_perpetual_fee"`
	PersonalPerpetualVol string `json:"personal_perpetual_vol"`
	TotalPerpetualFee    string `json:"total_perpetual_fee"`
	TotalPerpetualVol    string `json:"total_perpetual_vol"`
	PerpetualTraders     int    `json:"perpetual_traders"`
}

func Partners(pageSize, uid, inviterUid, startTime, endTime, startTimeLocal, endTimeLocal, level int, field string) (error, decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	totalDeposits := decimal.Zero
	totalFees := decimal.Zero
	totalTraders := decimal.Zero
	totalWithdraws := decimal.Zero
	totalTransactionAmount := decimal.Zero

	currentPage := 1
	for {
		// 构建当前页的URL
		currentPageURL := fmt.Sprintf("%spageNumber=%d&pageSize=%d&uid=%d&inviter_uid=%d&startTime=%d&endTime=%d&startTimeLocal=%d&endTimeLocal=%d&level=%d&field=%s", agentBackend.PartnerUrl, currentPage, pageSize, uid, inviterUid, startTime, endTime, startTimeLocal, endTimeLocal, level, field)
		responseTest, err := function.GetDetails(currentPageURL)
		if err != nil {
			fmt.Println(err)
		}

		var partnersResponse PartnersResponse
		if err := json.Unmarshal(responseTest, &partnersResponse); err != nil {
			return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero
		}

		// 检查响应是否成功
		if partnersResponse.Code == 0 {
			if partnersResponse.Result.Item != nil {
				for _, partners := range partnersResponse.Result.Item {
					// 转换并累计 deposits
					depositDec, err := decimal.NewFromString(partners.Deposit)
					if err == nil {
						totalDeposits = totalDeposits.Add(depositDec)
					}

					// 转换并累计 perpetual fees
					feeDec, err := decimal.NewFromString(partners.PersonalPerpetualFee)
					if err == nil {
						totalFees = totalFees.Add(feeDec)
					}

					// 转换并累计 withdrawals
					withdrawDec, err := decimal.NewFromString(partners.Withdraw)
					if err == nil {
						totalWithdraws = totalWithdraws.Add(withdrawDec)
					}

					// 累计 perpetual traders
					if partners.PersonalPerpetualVol != "0" {
						totalTraders = totalTraders.Add(future.One)
					}

					// 计算交易总额（存款+取款）
					personalPerpetualVolDec, err := decimal.NewFromString(partners.PersonalPerpetualVol)
					if err == nil {
						totalTransactionAmount = totalTransactionAmount.Add(personalPerpetualVolDec)
					}
				}
			}

			// 检查是否还有更多的页面
			totalCurrentPage := pageSize * currentPage
			if totalCurrentPage < partnersResponse.Result.Total {
				// 更新当前页数
				currentPage++
			} else {
				// 所有页面都已遍历完成
				break
			}
		} else {
			return errors.New("请求失败"), decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero
		}
	}

	return nil, totalDeposits, totalFees, totalTraders, totalWithdraws, totalTransactionAmount
}
