package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver for xormz
)

// Connect connects the database
func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "admin:csci1502017@tcp(csci150-mysql-sg.cvawt8ol1m2q.us-east-2.rds.amazonaws.com)/StudyGroup")
	if err != nil {
		return
	}
	/*
		db, err = sql.Open("mysql", "admin:industry1522018@tcp(csci152mysql.cp4qn9rvk7mn.us-west-1.rds.amazonaws.com)/SG_Accounts")
		if err != nil {
			return
		}*/
	return
}
