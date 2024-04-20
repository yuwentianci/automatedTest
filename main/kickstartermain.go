package main

import "myapp/text"

func main() {
	id2 := "19"
	text.KickstarterCreatePosition(id2)

	projectDetails2 := text.KickstarterProject(id2)
	totalValidCommitAmountNoRobot2, totalValidCommitAmount2 := text.KickstarterPosition(id2)
	text.KickstarterPositionIncome(totalValidCommitAmountNoRobot2, totalValidCommitAmount2, projectDetails2)
}
