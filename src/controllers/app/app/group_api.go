package app

import (
	"groupup/src/models"
	"log"
)

func groupCreate(args wsAPIstruct) error {
	args.Msg.Groupid = models.AddGroup(args.Msg.Groupid, args.UserToken)
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

func groupWhiteboard(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	msgJSON := args.Msg

	for _, c := range models.GetOtherConnectionsInGroup(args.UserToken, groupid) {
		err := c.WriteJSON(msgJSON)
		if err != nil {
			log.Println(err.Error())
		}
	}

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
