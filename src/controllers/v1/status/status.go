package status

import "database/sql"

var db *sql.DB

func Init(DB *sql.DB) {
	db = DB
}
