package main

import (
	"myapp/dao"
)

func main() {
	id2 := "19"
	dao.KickstarterCreatePosition(id2)

	projectDetails2 := dao.KickstarterProject(id2)
	totalValidCommitAmountNoRobot2, totalValidCommitAmount2 := dao.KickstarterPosition(id2)
	dao.KickstarterPositionIncome(totalValidCommitAmountNoRobot2, totalValidCommitAmount2, projectDetails2)
}
