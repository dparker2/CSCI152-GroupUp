package app

import (
	"errors"
	"groupup/src/models"
	"log"
)

func groupCreate(args wsAPIstruct) error {
	args.Msg.Groupid = models.AddGroup(args.Msg.Groupid, args.UserToken)
	usrConn := models.GetConnection(args.UserToken)

	usrConn.WriteJSON(args.Msg)

	return nil
}

func groupRemove(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	userToken := args.UserToken
	usrConn := models.GetConnection(args.UserToken)
	username := models.GetUsername(userToken)

	models.RemoveGroupFromUser(userToken, groupid)

	usrConn.WriteJSON(&wsMessage{
		Code:    "app/current/remove",
		Groupid: groupid,
	})

	writeJSONToGroup(groupid, &wsMessage{
		Code:     "group/remove",
		Username: username,
		Groupid:  groupid,
	})
	return nil
}

func groupJoin(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	userToken := args.UserToken
	usrConn := models.GetConnection(args.UserToken)
	putInUsername(&args)

	if models.GroupExists(groupid) { // When we get DB setup, this should check it
		err := models.AddUserToGroup(userToken, groupid)
		if err != nil {
			usrConn.WriteJSON(&wsMessage{
				Code: "group", // No other args shows failure to join
			})
			return err
		}

		if !models.UserHasCurrentGroup(userToken, groupid) {
			models.AddGroupToUsersCurrentGroups(userToken, groupid)
			usrConn.WriteJSON(&wsMessage{
				Code:    "app/current/add",
				Groupid: groupid,
			})
		}

		usrConn.WriteJSON(&wsMessage{
			Code:    "group",
			Groupid: groupid, // "Okay to render"
		})

		fullChatLog := models.GetChatLogFromDB(groupid)
		for _, chat := range fullChatLog {
			timestamp := chat[0]
			username := chat[1]
			message := chat[2]
			usrConn.WriteJSON(&wsMessage{
				Code:      "group/chat",
				Groupid:   groupid,
				Username:  username,
				Timestamp: timestamp,
				Chat:      message,
			})
		}

		fullUserlist := models.GetFullUserListWithStatus(groupid)
		for _, user := range fullUserlist {
			username := user[0]
			status := user[1]
			usrConn.WriteJSON(&wsMessage{
				Code:     "group/join",
				Groupid:  groupid,
				Username: username,
				Status:   status,
			})
		}
		// get flashcards from db here
		flashcards, _ := models.GetFlashcardsFromDB(groupid)
		for _, card := range flashcards {
			index := card[0]
			front := card[1]
			back := card[2]
			usrConn.WriteJSON(&wsMessage{
				Code:    "group/flashcards/new",
				Groupid: groupid,
				Front:   front,
				Back:    back,
				Index:   index,
			})
		}
		writeJSONToOthersInGroup(groupid, userToken, args.Msg)
		return nil
	}
	return errors.New("Group does not exist")
}

func groupLeave(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	userToken := args.UserToken
	putInUsername(&args)

	err := models.RemoveUserFromGroup(userToken, groupid)
	if err != nil {
		return err
	}
	writeJSONToGroup(groupid, args.Msg)
	return nil
}

func groupChat(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	putInUsername(&args)

	writeJSONToGroup(groupid, args.Msg)
	models.WriteChatToDB(groupid, args.Msg.Username, args.Msg.Chat)
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

func groupFlashcardNew(args wsAPIstruct) (err error) {
	groupid := args.Msg.Groupid
	uuid := models.GetUserID(args.UserToken)
	args.Msg.Index, err = models.InsertCardToDB(groupid, uuid)
	if err != nil {
		return err
	}
	writeJSONToGroup(groupid, args.Msg)
	return
}

func groupFlashcardEditFront(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	index := args.Msg.Index
	front := args.Msg.Front
	uuid := models.GetUserID(args.UserToken)

	err := models.UpdateFlashcardFront(groupid, index, front, uuid)
	if err != nil {
		return err
	}
	writeJSONToGroup(groupid, args.Msg)

	return nil
}

func groupFlashcardEditBack(args wsAPIstruct) error {
	groupid := args.Msg.Groupid
	index := args.Msg.Index
	back := args.Msg.Back
	uuid := models.GetUserID(args.UserToken)

	err := models.UpdateFlashcardBack(groupid, index, back, uuid)
	if err != nil {
		return err
	}
	writeJSONToGroup(groupid, args.Msg)

	return nil
}

func putInUsername(args *wsAPIstruct) {
	args.Msg.Username = models.GetUsername(args.UserToken)
}
