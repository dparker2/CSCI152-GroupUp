package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect connects the database
func Connect() (db *sql.DB, dbAcc *sql.DB, err error, errAcc error) {
	db, err = sql.Open("mysql", "admin:csci1502017@tcp(csci150-mysql-sg.cvawt8ol1m2q.us-east-2.rds.amazonaws.com)/StudyGroup")
	if err != nil {
		panic(err.Error())
	}

	dbAcc, errAcc = sql.Open("mysql", "admin:industry1522018@tcp(csci152mysql.cp4qn9rvk7mn.us-west-1.rds.amazonaws.com)/SG_Accounts")
	if errAcc != nil {
		panic(errAcc.Error())
	}

	return db, dbAcc, err, errAcc
}
