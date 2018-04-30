package app

import (
	"groupup/src/models"
	"log"
)

func groupJoin(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	userToken := args.UserToken

	//if models.GroupExists(groupid) { // When we get DB setup, this should check it
	models.AddUserToGroup(userToken, groupid)
	return nil
	//}
	//return errors.New("Group does not exist")
}

func groupChat(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	args.Msg.Username = models.GetUsername(args.UserToken) // Fill in username to send out
	msgJSON := args.Msg

	for _, c := range models.GetConnectionsInGroup(groupid) {
		err := c.WriteJSON(msgJSON)
		if err != nil {
			log.Println(err.Error())
		}
	}
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
