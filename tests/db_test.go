package tests

import (
	"groupup/src/system/db"
	"testing"
)

func TestConnect(t *testing.T) {
	DB, err := db.Connect()
	if err != nil {
		t.Error(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		t.Error(err.Error())
	}
}
