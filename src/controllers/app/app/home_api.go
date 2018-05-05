package app

import (
	"groupup/src/models"
	"log"
)

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

	offlineFriendsList := models.GetOfflineFriendsList(userToken)
	for _, friendName := range offlineFriendsList {
		usrConn.WriteJSON(&wsMessage{
			Code:     "app/friends/offline",
			Username: friendName,
		})
	}

	onlineFriendsList := models.GetOnlineFriendsList(userToken)
	for _, friendName := range onlineFriendsList {
		usrConn.WriteJSON(&wsMessage{
			Code:     "app/friends/online",
			Username: friendName,
		})
	}

	letFollowersKnow(userToken, &wsMessage{
		Code:     "app/friends/online",
		Username: args.Msg.Username,
	})

	return nil
}

func searchUsers(args wsAPIstruct) error {
	usrConn := models.GetConnection(args.UserToken)
	query := args.Msg.Query

	usernames, err := models.SearchUsersInDB(query)
	if err != nil {
		return err
	}
	for _, name := range usernames {
		log.Println("writing ", name, " from query: ", query)
		usrConn.WriteJSON(&wsMessage{
			Code:     "app/search/users",
			Username: name,
			Query:    query,
		})
	}

	return nil
}

func friendsAdd(args wsAPIstruct) error {
	usrConn := models.GetConnection(args.UserToken)
	friendName := args.Msg.Username

	err := models.AddFriendToUser(args.UserToken, friendName)
	if err != nil {
		return err
	}

	code := "app/friends/offline"
	if models.UserExistsByUsername(friendName) {
		code = "app/friends/online"
	}
	usrConn.WriteJSON(&wsMessage{
		Code:     code,
		Username: friendName,
	})

	return nil
}

func friendsRemove(args wsAPIstruct) error {

	return nil
}
