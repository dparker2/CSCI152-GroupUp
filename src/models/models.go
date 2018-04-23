package models

import (
	DB "groupup/src/system/db"
)

var users map[string]*user
var groups map[string]*group

func init() {
	// Package variables for state
	users = make(map[string]*user)
	groups = make(map[string]*group)

	// Connect to the DB
	db, err := DB.Connect()
	err = db.Ping()
	if err != nil {
		panic(err) // If no DB just fail
	}
}
