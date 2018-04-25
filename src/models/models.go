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

	doThing()
}

func doThing() {
	result, _ := dbAcc.Exec("INSERT INTO UserInfo (Username) VALUES ('groupup')")

	result2, _ := db.Exec("INSERT INTO a_1739 (Admin) VALUES ('spongebob')")
	fmt.Println(result, result2)
}
