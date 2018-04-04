package db

import (
	_ "github.com/go-sql-driver/mysql" // MySQL driver for xormz
	"github.com/go-xorm/xorm"
)

// Connect connects the database
func Connect() (db *xorm.Engine, err error) {
	db, err = xorm.NewEngine("mysql", "test:test@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		return
	}

	db.DBMetas()

	return
}
