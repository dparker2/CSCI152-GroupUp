package app

import "groupup/src/models"

func home(args wsAPIstruct) error {
	userToken := args.UserToken
	usrConn := models.GetConnection(args.UserToken)
	putInUsername(&args)

	currGrps := models.GetCurrentGroups(userToken)
	for _, grpName := range currGrps {
		usrConn.WriteJSON(&wsMessage{
			Code:    "app/current/add",
			Groupid: grpName,
		})
	}

	return nil
}
