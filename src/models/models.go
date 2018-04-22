package models

import (
	"fmt"
	DB "groupup/src/system/db"
)

var users map[string]*user
var groups map[string]*group

func init() {
	// Package variables for state
	users = make(map[string]*user)
	groups = make(map[string]*group)

	// Connect to both databases
	db, dbAcc, err, errAcc := DB.Connect()

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
