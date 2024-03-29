package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
func SearchGroupsInDB(str string) (grps [][]string, err error) {
	rows, err := db.Query("SELECT * FROM GroupIndex WHERE GroupID LIKE CONCAT('%', ? ,'%') ORDER BY GroupID ASC LIMIT 20", str)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var groupname, subs, creator string
		err = rows.Scan(&groupname, &subs, &creator)
		if err != nil {
			panic(err)
		}
		var result []string
		result = append(result, groupname, subs, creator)
		grps = append(grps, result)
	}
	return
}

// PutAdminInGroupDB adds admin to database
func PutAdminInGroupDB(groupid string, admin string) (err error) {
	adminstmt, err := db.Prepare("INSERT INTO " + groupid + " (Admin) VALUES (?)")
	if err != nil {
		panic(err)
	}
	_, err = adminstmt.Exec(admin)
	PutInGroupIndex(groupid, admin)
	return
}

func PutInGroupIndex(groupid string, creator string) (err error) {
	stmt, err := db.Prepare("INSERT INTO GroupIndex (GroupID, SubbedUsers, Creator) VALUES (?, 0, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(groupid, creator)
	return
}

func RemoveFromGroupIndex(groupid string) (err error) {
	stmt, err := db.Prepare("DELETE FROM GroupIndex WHERE GroupID = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(groupid)
	return
}

func IncreaseGroupIndexSubs(groupid string) (err error) {
	stmt, err := db.Prepare("UPDATE GroupIndex SET SubbedUsers = SubbedUsers + 1 WHERE GroupID = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(groupid)
	return
}

func DecreaseGroupIndexSubs(groupid string) (err error) {
	stmt, err := db.Prepare("UPDATE GroupIndex SET SubbedUsers = SubbedUsers - 1 WHERE GroupID = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(groupid)
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
		IncreaseGroupIndexSubs(groupid)
	case err != nil:
		panic(err)
	default:
		if username == userinDB {
			fmt.Println("User already exists in DB userList")
		}
	}
	return
}

func RemoveUserFromGroupDB(groupid string, username string) (err error) {
	stmt, err := db.Prepare("DELETE FROM " + groupid + " WHERE userList = ?")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(username)
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

func InsertCardToDB(groupid string, uuid int) (index string, err error) {
	stmt, err := db.Prepare("INSERT INTO Flashcards (GroupID, UserID) VALUES (?, ?)")
	if err != nil {
		return
	}
	result, err := stmt.Exec(groupid, uuid)
	id, err := result.LastInsertId()
	index = strconv.FormatInt(id, 10)

	return
}

func GetFlashcardsFromDB(groupid string) (flashcards [][]string, err error) {
	rows, err := db.Query("SELECT FlashcardIndex, Front, Back FROM Flashcards WHERE GroupID = ?", groupid)
	if err != nil {
		return nil, err
	}
	fmt.Println("Get flashcards was called")
	fmt.Println(rows)

	for rows.Next() {
		var index string
		var front string
		var back string
		err = rows.Scan(&index, &front, &back)
		fmt.Println("TEST", index, front, back)
		if err != nil {
			return nil, err
		}
		var card []string
		card = append(card, index, front, back)
		flashcards = append(flashcards, card)
	}
	fmt.Println(flashcards)
	return
}

func UpdateFlashcardFront(groupid string, index string, front string, uuid int) (err error) {
	stmt, err := db.Prepare("UPDATE Flashcards SET Front = ?, UserID = ? WHERE (GroupID = ? AND FlashcardIndex = ?)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(front, uuid, groupid, index)
	return
}

func UpdateFlashcardBack(groupid string, index string, back string, uuid int) (err error) {
	stmt, err := db.Prepare("UPDATE Flashcards SET Back = ?, UserID = ? WHERE (GroupID = ? AND FlashcardIndex = ?)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(back, uuid, groupid, index)
	return
}
