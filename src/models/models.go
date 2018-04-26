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

func VerifyLogin(username string, password string) bool {
	verify := false
	//create user object? add to current user object? temporary solution...
	var passwordDB, email string

	//query db for username's info.. e.g. password, email, sec ?'s, etc
	err := dbAcc.QueryRow("SELECT Pass, Email FROM UserInfo WHERE Username = ?", username).Scan(&passwordDB, &email)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No user with that username.")
	case err != nil:
		panic(err)
	default:
		if password == passwordDB {
			verify = true
		} // add incorrect password response
	}
	return verify
	//compare input password with db password

	//send email
}
