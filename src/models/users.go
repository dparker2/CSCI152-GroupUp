package models

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type user struct {
	Name           string
	WsConn         *websocket.Conn
	Status         int
	Token          string
	CurrentGroups  []*group
	FavoriteGroups []*group
	RecentGroups   []*group
}

var userMutex = &sync.Mutex{}

// NewUser makes a new user with the given username, token, and IP
func NewUser(username string) (userToken string) {
	userToken, err := generateRandomString(32)
	if err != nil {
		log.Println(err.Error())
	}
	u := user{
		Name:  username,
		Token: userToken,
	}
	// Thread safe users modification.
	userMutex.Lock()
	users[userToken] = &u
	userMutex.Unlock()
	return
}

// RemoveUser assumes existense of user, and removes it
func RemoveUser(token string) {
	userMutex.Lock()
	delete(users, token)
	userMutex.Unlock()
	return
}

// UserExists checks if a user currently (ie connection is alive)
func UserExists(token string) bool {
	if _, exists := users[token]; exists {
		return true
	} else {
		return false
	}
}

// GetUserToken returns the token of a given user, assumes it exists
func GetUserToken(ip string) (token string) {
	token = users[ip].Token
	return
}

// GetUsername returns the username of a user associated with the token
func GetUsername(token string) (username string) {
	userMutex.Lock()
	if UserExists(token) {
		username = users[token].Name
	}
	userMutex.Unlock()
	return
}

// SetUserStatus sets the status of user associated with token
func SetUserStatus(token string, status int) {
	userMutex.Lock()
	if UserExists(token) {
		users[token].Status = status
	}
	userMutex.Unlock()
}

// SetUserConn sets the connection of the user associated with token
func SetUserConn(token string, conn *websocket.Conn) {
	userMutex.Lock()
	if UserExists(token) {
		users[token].WsConn = conn
	}
	userMutex.Unlock()
}
