package app

import "groupup/src/models"

func home(args wsAPIstruct) error {
	userToken := args.UserToken
	usrConn := models.GetConnection(args.UserToken)
	putInUsername(&args)

	prevGrps := models.GetPreviousGroups(userToken)
	for _, grpName := range prevGrps {
		usrConn.WriteJSON(&wsMessage{
			Code:    "app/previous/add",
			Groupid: grpName,
		})
	}

	currGrps := models.GetCurrentGroups(userToken)
	for _, grpName := range currGrps {
		usrConn.WriteJSON(&wsMessage{
			Code:    "app/current/add",
			Groupid: grpName,
		})
	}

	return nil
}
