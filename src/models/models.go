package models

import (
	"database/sql"
	"fmt"
	DB "groupup/src/system/db"
)

var users map[string]*user
var groups map[string]*group
var db *sql.DB
var dbAcc *sql.DB

func init() {
	// Package variables for state
	users = make(map[string]*user)
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
func VerifyLogin(username string, password string) (userToken string, verify bool) {
	verify = false
	//create user object and return token
	userToken = newUser(username)
	var passwordDB, email string
	var SQ1, SQ2, SQ3, SQA1, SQA2, SQA3 sql.NullString
	var userID int
	var lockoutStatus sql.NullInt64
	var secQuestions, secAnswers []sql.NullString
	//query db for username's info.. e.g. password, email, sec ?'s, etc
	err := dbAcc.QueryRow(
		"SELECT UserID, Pass, Email, SQ1, SQ2, SQ3, SQA1, SQA2, SQA3, LockoutStatus FROM UserInfo WHERE Username = ?", username,
	).Scan(
		&userID, &passwordDB, &email, &SQ1, &SQ2, &SQ3, &SQA1, &SQA2, &SQA3, &lockoutStatus)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No user with that username.")
	case err != nil:
		panic(err)
	default:
		if password == passwordDB {
			verify = true
			secQuestions = append(secQuestions, SQ1, SQ2, SQ3)
			secAnswers = append(secAnswers, SQA1, SQA2, SQA3)
			readIntoUser(userToken, userID, email, secQuestions[:], secAnswers[:], lockoutStatus)
		} // add incorrect password response
	}
	//send email
	return
}

func readIntoUser(userToken string, userID int, email string, secQuestions []sql.NullString, secAnswers []sql.NullString, lockoutStatus sql.NullInt64) {
	users[userToken].UserID = userID
	users[userToken].Email = email
	for i := 0; i < 3; i++ {
		users[userToken].SecQuestions = append(users[userToken].SecQuestions, secQuestions[i])
		users[userToken].SecAnswers = append(users[userToken].SecAnswers, secAnswers[i])
	}
	users[userToken].LockoutStatus = lockoutStatus
}
