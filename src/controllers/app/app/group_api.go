package app

import (
	"groupup/src/models"
)

func groupCreate(args wsAPIstruct) error {
	args.Msg.Groupid = models.AddGroup(args.Msg.Groupid)
	usrConn := models.GetConnection(args.UserToken)

	usrConn.WriteJSON(args.Msg)

	return nil
}

func groupJoin(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	userToken := args.UserToken
	usrConn := models.GetConnection(args.UserToken)
	putInUsername(&args)

	//if models.GroupExists(groupid) { // When we get DB setup, this should check it
	err := models.AddUserToGroup(userToken, groupid)
	if err != nil {
		usrConn.WriteJSON(&wsMessage{
			Code: "group", // No other args shows failure to join
		})
		return err
	}
	usrConn.WriteJSON(&wsMessage{
		Code:    "group",
		Groupid: groupid, // "Okay to render"
	})
	writeJSONToGroup(groupid, args.Msg)
	return nil
	//}
	//return errors.New("Group does not exist")
}

func groupChat(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	putInUsername(&args)

	writeJSONToGroup(groupid, args.Msg)
	return nil
}

func groupWhiteboardDraw(args wsAPIstruct) error {
	return nil
}

func groupFlashcardNew(args wsAPIstruct) error {
	return nil
}

func groupFlashcardEdit(ars wsAPIstruct) error {
	return nil
}

func putInUsername(args *wsAPIstruct) {
	args.Msg.Username = models.GetUsername(args.UserToken)
}
