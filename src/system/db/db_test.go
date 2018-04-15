package db

import (
	"testing"
)

func TestConnect(t *testing.T) {
	DB, err := Connect()
	if err != nil {
		t.Error(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		t.Error(err.Error())
	}
}
