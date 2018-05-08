package models

import (
	"database/sql"
	"fmt"
	DB "groupup/src/system/db"
	"log"
)

var users map[string]*user
var userTokens map[string]string
var groups map[string]*group
var db *sql.DB
var dbAcc *sql.DB

func init() {
	// Package variables for state
	users = make(map[string]*user)
	userTokens = make(map[string]string)
	groups = make(map[string]*group)
	var err, errAcc error
	// Connect to both databases
	db, dbAcc, err, errAcc = DB.Connect()

	// Ping both databases to guarantee no connection errors
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to group db")
	}

	errAcc = dbAcc.Ping()
	if errAcc != nil {
		panic(errAcc)
	} else {
		fmt.Println("Successfully connected to the account db")
	}
}

//VerifyLogin verifies credentials and then creates user object if verified
func VerifyLogin(username string, password string) (u user, verify bool) {
	verify = false
	//create user object and return token
	u = newUser(username)
	var passwordDB string
	//query db for username's info.. e.g. password, email, sec ?'s, etc
	err := dbAcc.QueryRow(
		"SELECT UserID, Pass, Email, SQ1, SQ2, SQ3, SQA1, SQA2, SQA3, LockoutStatus FROM UserInfo WHERE Username = ?", username,
	).Scan(
		&u.UserID, &passwordDB, &u.Email, &u.SecQuestions[0], &u.SecQuestions[1], &u.SecQuestions[2],
		&u.SecAnswers[0], &u.SecAnswers[1], &u.SecAnswers[2], &u.LockoutStatus)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No user with that username.")
	case err != nil:
		panic(err)
	default:
		if password == passwordDB {
			verify = true

			groupNames, err := db_GetUsersGroups(u.UserID)
			if err != nil {
				log.Println(err.Error())
				return u, false
			}
			userGroups := groupNamesToObjects(groupNames)
			u.CurrentGroups = userGroups

			friendsList, err := db_GetUsersFriends(u.UserID)
			if err != nil {
				log.Println(err.Error())
				return u, false
			}
			u.AllFriends = friendsList

			onlineFriendsList := getOnlineUsers(friendsList)
			u.OnlineFriends = onlineFriendsList

			log.Println(u)
			insertIntoUsers(u)
		} // add incorrect password response
	}
	//send email
	return
}

func insertIntoUsers(u user) {
	userMutex.Lock()
	users[u.Token] = &u
	userTokens[u.Name] = u.Token
	userMutex.Unlock()
}

func groupNamesToObjects(groupNames []string) (gl groupList) {
	for _, g := range groupNames {
		if GroupExists(g) {
			gl = append(gl, groups[g])
		}
	}
	return
}

func getOnlineUsers(usernames []string) (ul userList) {
	for _, name := range usernames {
		if UserExistsByUsername(name) {
			u := users[userTokens[name]]
			ul = append(ul, u)
		}
	}
	return
}
