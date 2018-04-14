package models

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	testIP := "1.1.1.1:1000"
	testUsername := "testusername"
	NewUser(testIP, testUsername)
	testuser := users["1.1.1.1:1000"]
	if testuser == nil {
		t.Error("Test user not created in package variable users in NewUser")
	}
	if testuser.IP != testIP {
		t.Error("IP not set properly in NewUser")
	}
	if testuser.Name != testUsername {
		t.Error("Username not set properly in NewUser")
	}
	if testuser.Token == "" {
		t.Error("Token not generated in NewUser")
	}
}

func TestUserExists(t *testing.T) {
	testIP := "1.1.1.1:1000"
	NewUser(testIP, "test")
	b := UserExists(testIP)
	if !b {
		t.Error("User created but UserExists returns false.")
	}
	if b != (users[testIP] != nil) {
		t.Error("UserExists returns wrong value.")
	}
}
