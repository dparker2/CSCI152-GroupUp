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
			insertIntoUsers(u)
		} // add incorrect password response
	}
	//send email
	return
}

func insertIntoUsers(u user) {
	userMutex.Lock()
	users[u.Token] = &u
	userMutex.Unlock()
}

// VerifyRegister checks for existence of username and email in dbAcc
func VerifyRegister(username string, email string) (verify bool) {
	return (!usernameTaken(username) && !emailTaken(email))
}

func usernameTaken(username string) (exists bool) {
	row := dbAcc.QueryRow("SELECT EXISTS(SELECT 1 FROM UserInfo WHERE Username = ?)", username)
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}

func emailTaken(email string) (exists bool) {
	row := dbAcc.QueryRow("SELECT EXISTS(SELECT 1 FROM UserInfo WHERE Email = ?)", email)
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}

func CreateAccount(username, password, email string) bool {
	/*
		stmt, err := db.Prepare("INSERT INTO UserInfo (Username, Pass, Email) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(username, password, email)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
	*/
	_, err := dbAcc.Exec("INSERT INTO UserInfo (Username, Pass, Email) VALUES (?, ?, ?)", username, password, email)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Something went wrong inserting the data.")
		return false
	case err != nil:
		panic(err)
	default:
		return true
	}
}
