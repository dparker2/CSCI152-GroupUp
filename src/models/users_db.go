package models

import (
	"database/sql"
	"errors"
	"log"
)

func SearchUsersInDB(str string) (usrs []string, err error) {
	stmt, err := dbAcc.Prepare("SELECT Username FROM UserInfo WHERE Username LIKE CONCAT('%', ? ,'%') ORDER BY Username ASC LIMIT 20")
	if err != nil {
		return nil, err
	}
	usernames, err := stmt.Query(str)
	if err != nil {
		return nil, err
	}
	for usernames.Next() {
		var u string
		err = usernames.Scan(&u)
		if err != nil {
			log.Println(err.Error())
			return
		}
		usrs = append(usrs, u)
	}
	return
}

func db_AddFriendToUser(uuid int, friendname string) error {
	friendid, err := db_GetUsersID(friendname)
	if friendid == 0 || err != nil {
		return err
	}

	stmt, err := dbAcc.Prepare("INSERT INTO FriendTest (followerID, followedID) VALUES (?, ?)")

	if err != nil {
		return err
	}
	log.Println(uuid, friendid)
	_, err = stmt.Exec(uuid, friendid)
	return err
}

func db_GetUsersID(username string) (id int, err error) {
	stmt, err := dbAcc.Prepare("SELECT UserID FROM UserInfo WHERE (Username = ?)")
	// FIX
	if err != nil {
		return 0, err
	}
	row, err := stmt.Query(username)
	if !row.Next() {
		return 0, errors.New("No user with name " + username)
	}
	row.Scan(&id)
	return
}

func db_AddGroupToUsersGroups(uuid int, groupid string) error {
	stmt, err := dbAcc.Prepare("INSERT INTO GroupMapping (UserID, SubbedGroup) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid, groupid)
	return err
}

func db_RemoveGroupFromUsersGroups(uuid int, groupid string) error {
	stmt, err := dbAcc.Prepare("DELETE FROM GroupMapping WHERE (UserID = ?) AND (SubbedGroup = ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid, groupid)
	return err
}

func db_GetUsersGroups(uuid int) (sl []string, err error) {
	stmt, err := dbAcc.Prepare("SELECT g.SubbedGroup FROM GroupMapping g INNER JOIN UserInfo u ON u.UserID = g.UserID WHERE g.UserID = ?")
	if err != nil {
		return nil, err
	}
	groupNames, err := stmt.Query(uuid)
	if err != nil {
		return nil, err
	}
	for groupNames.Next() {
		var g string
		err = groupNames.Scan(&g)
		if err != nil {
			log.Println(err.Error())
			return
		}
		sl = append(sl, g)
	}
	return
}

func db_GetUsersFriends(uuid int) (sl []string, err error) {
	stmt, err := dbAcc.Prepare("SELECT Username FROM UserInfo WHERE UserID IN (SELECT f.followedID FROM FriendTest f INNER JOIN UserInfo u ON u.UserID = f.followerID WHERE f.followerID = ?)")
	if err != nil {
		return nil, err
	}
	sl, err = followQuery(stmt, uuid)
	return
}

func followQuery(stmt *sql.Stmt, uuid int) (sl []string, err error) {
	friends, err := stmt.Query(uuid)
	if err != nil {
		return nil, err
	}
	for friends.Next() {
		var f string
		err = friends.Scan(&f)
		if err != nil {
			log.Println(err.Error())
			return
		}
		sl = append(sl, f)
	}
	return
}

func db_GetUsersFollowersDB(uuid int) (sl []string, err error) {
	stmt, err := dbAcc.Prepare("SELECT Username FROM UserInfo WHERE UserID IN (SELECT f.followerID FROM FriendTest f INNER JOIN UserInfo u ON u.UserID = f.followedID WHERE f.followedID = ?)")
	if err != nil {
		return nil, err
	}
	sl, err = followQuery(stmt, uuid)
	return
}
