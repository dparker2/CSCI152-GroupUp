package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// CreateGroupInDB creates the group in the database
func CreateGroupInDB(groupid string) (err error) {
	fmt.Println("Creating group " + groupid)
	_, err = db.Exec("CREATE TABLE " + groupid + " (Admin varchar(50), userList varchar(20), ipAddress varchar(50), User varchar(20), Clock datetime, Message varchar(255), Whiteboard LONGTEXT)")
	if err != nil {
		panic(err)
	}
	return
}

// SearchGroupsInDB returns a list of string arrays where
//   0 index is groupname
//   1 index is number of users
//   2 index is creator
func SearchGroupsInDB(str string) (usrs [][]string, err error) {
	/*stmt, err := dbAcc.Prepare("SELECT Username FROM UserInfo WHERE Username LIKE CONCAT('%', ? ,'%') ORDER BY Username ASC LIMIT 20")
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
	}*/
	return
}

// PutAdminInGroupDB adds admin to database
func PutAdminInGroupDB(groupid string, admin string) (err error) {
	adminstmt, err := db.Prepare("INSERT INTO " + groupid + " (Admin) VALUES (?)")
	if err != nil {
		panic(err)
	}
	_, err = adminstmt.Exec(admin)

	return
}

// AddUserToGroupDB checks if the user is already in the userList in the database and adds them if not.
func AddUserToGroupDB(groupid string, username string) (err error) {
	addstmt, err := db.Prepare("INSERT INTO " + groupid + " (userList) VALUES (?)")
	if err != nil {
		panic(err)
	}
	defer addstmt.Close()
	var userinDB string
	err = db.QueryRow("SELECT userList FROM "+groupid+" WHERE userList = ?", username).Scan(&userinDB)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("User doesn't already exist in userList, adding...")
		_, err = addstmt.Exec(username)
	case err != nil:
		panic(err)
	default:
		if username == userinDB {
			fmt.Println("User already exists in DB userList")
		}
	}
	return
}

// WriteChatToDB stores chat into the DB
func WriteChatToDB(groupid string, timestamp string, username string, msg string) (err error) {
	chatstmt, err := db.Prepare("INSERT INTO " + groupid + " (user, Clock, Message) VALUES (?, ?, ?)")
	fmt.Println("Inserting chat to DB...")
	if err != nil {
		return err
	}
	_, err = chatstmt.Exec(username, timestamp, msg)

	if err != nil {
		return err
	}
	return
}

// GetChatLogFromDB will fetch all the messages in the group to return to the user
func GetChatLogFromDB(groupid string) (chatLog [][]string) {
	chatstmt, err := db.Query("SELECT user, Clock, Message FROM " + groupid + " WHERE Message IS NOT NULL")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer chatstmt.Close()
	for chatstmt.Next() {
		var timestamp string
		var username string
		var message string
		err = chatstmt.Scan(&username, &timestamp, &message)
		if err != nil {
			log.Println(err.Error())
			return
		}
		ts, _ := time.Parse("2006-01-02 15:04:05", timestamp)
		timestamp = ts.Format(time.RFC3339)
		var aChat []string
		aChat = append(aChat, timestamp, username, message)
		chatLog = append(chatLog, aChat)
	}
	return
}

// GetFullUserListFromDB returns the userList in the database in alphabetical order
func GetFullUserListFromDB(groupid string) (fullUserList []string) {
	fullusers, err := db.Query("select userList from " + groupid + " WHERE userList IS NOT NULL ORDER BY userList ASC")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer fullusers.Close()
	for fullusers.Next() {
		var currentUser string
		err = fullusers.Scan(&currentUser)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fullUserList = append(fullUserList, currentUser)
	}
	return
}
