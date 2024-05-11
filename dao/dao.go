package dao

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"myapp/function"
	"time"
)

type kickstarterCreatePosition struct {
	Code int `json:"code"`
	Wait int `json:"wait"`
}

func KickstarterCreatePosition(id string) {
	// uid 51316260-51317260
	for uid := 51316750; uid < 51316780; uid++ {
		nowTime := time.Now()

		idDec, _ := decimal.NewFromString(id)
		idInt := idDec.IntPart()

		// 设置随机数种子，以确保每次运行都得到不同的随机数序列
		rand.Seed(time.Now().UnixNano())

		// 生成100000000到500000000之间的随机数
		randomNumber := rand.Intn(400000001) + 100000000 // 生成0到400000000之间的随机数，然后加上100000000
		format := map[string]interface{}{
			"row[created_at]":                nowTime,
			"row[updated_at]":                nowTime,
			"row[user_id]":                   uid,
			"row[sunshine_activity_id]_text": idInt,
			"row[sunshine_activity_id]":      idInt,
			"row[commit_amount]":             randomNumber,
			"row[is_robot]":                  1,
		}
		url := "https://dao.biconomy.vip/houtai.php/sunshine_user_hold/add?dialog=1"
		var kickstarterPositionData = kickstarterCreatePosition{}
		if err := function.PostFormData(url, format, &kickstarterPositionData); err != nil {
			fmt.Println(err.Error())
			return
		}

		if kickstarterPositionData.Code == 1 {
			fmt.Println("uid:", uid, "参与量:", randomNumber)
		}
	}
}

type kickstarterPosition struct {
	Rows []positionDetails `json:"rows"`
}

type positionDetails struct {
	SunshineActivityID string  `json:"sunshine_activity_id"`
	IsRobot            string  `json:"is_robot"`
	ValidCommitAmount  float64 `json:"valid_committ_amount"`
}

// KickstarterPosition 用户阳光普照持仓情况
func KickstarterPosition(id string) (totalValidCommitAmountNoRobot, totalValidCommitAmount float64) {
	url1 := "https://dao.biconomy.vip/houtai.php/sunshine_user_hold/index?addtabs=1&sort=id&order=desc&offset=0&limit=1000&filter=%7B%22sunshine_activity_id%22%3A%22"

	url2 := "%22%7D&op=%7B%22sunshine_activity_id%22%3A%22%3D%22%7D&_=1703384034642"
	url := fmt.Sprintf("%s%s%s", url1, id, url2)
	responseTest, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var kickstarterPositionData kickstarterPosition
	if err := json.Unmarshal(responseTest, &kickstarterPositionData); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if kickstarterPositionData.Rows != nil {
		var totalValidCommitAmount float64
		var totalValidCommitAmountNoRobot float64
		for _, positionData := range kickstarterPositionData.Rows {
			if positionData.SunshineActivityID == id {
				totalValidCommitAmount += positionData.ValidCommitAmount
			}
			if positionData.SunshineActivityID == id && positionData.IsRobot == "0" {
				totalValidCommitAmountNoRobot += positionData.ValidCommitAmount
			}
		}

		totalValidCommitAmountString := decimal.NewFromFloat(totalValidCommitAmount)
		totalValidCommitAmountNoRobotString := decimal.NewFromFloat(totalValidCommitAmountNoRobot)
		fmt.Println("真人总有效投入数量:", totalValidCommitAmountNoRobotString, "总有效投入数量:", totalValidCommitAmountString)
		return totalValidCommitAmountNoRobot, totalValidCommitAmount
	}
	return
}

type kickstarterProject struct {
	Total int              `json:"total"`
	Rows  []ProjectDetails `json:"rows"`
}

type ProjectDetails struct {
	ID             string `json:"id"`
	PrizeOneToken  string `json:"prize_one_token"`
	PrizeOneAmount string `json:"prize_one_amount"`
	PrizeTwoToken  string `json:"prize_two_token"`
	PrizeTwoAmount string `json:"prize_two_amount"`
}

// KickstarterProject 用户阳光普照项目
func KickstarterProject(id string) (projectData ProjectDetails) {
	url := "https://dao.biconomy.vip/houtai.php/sunshine_projects/index?addtabs=1&sort=id&order=desc&offset=0&limit=100&filter=%7B%7D&op=%7B%7D&_=1702968000765"
	responseText, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	var kickstarterProjectData kickstarterProject
	if err := json.Unmarshal(responseText, &kickstarterProjectData); err != nil {
		fmt.Println("解析JSON响应时发生错误:", err)
		return
	}

	if kickstarterProjectData.Rows != nil {
		for _, projectData := range kickstarterProjectData.Rows {
			if projectData.ID == id {
				return projectData
			}
		}
	}
	return
}

// KickstarterPositionIncome 用户阳光普照真人收益
func KickstarterPositionIncome(totalValidCommitAmountNoRobot, totalValidCommitAmount float64, projectDetails ProjectDetails) {

	if projectDetails.PrizeOneToken != "" {
		oneAmountFloat := function.FormatFloat(projectDetails.PrizeOneAmount)
		oneIncome := totalValidCommitAmountNoRobot / totalValidCommitAmount * oneAmountFloat
		oneIncomeDec := decimal.NewFromFloat(oneIncome)
		fmt.Printf("阳光产品%s真人总收益为: %s %s", projectDetails.ID, oneIncomeDec, projectDetails.PrizeOneToken)
		if projectDetails.PrizeTwoToken != "" {
			twoAmountFloat := function.FormatFloat(projectDetails.PrizeTwoAmount)
			twoIncome := totalValidCommitAmountNoRobot / totalValidCommitAmount * twoAmountFloat
			twoIncomeDec := decimal.NewFromFloat(twoIncome)
			fmt.Println("真人总收益为:", twoIncomeDec, projectDetails.PrizeTwoToken)
		}
	}
}
