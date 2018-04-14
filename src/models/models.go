package models

import (
	"github.com/gorilla/websocket"

	DB "groupup/src/system/db"
)

// TODO: Separate files for users/groups/etc

type user struct {
	Name           string
	WsConn         *websocket.Conn
	CurrentGroups  []*group
	FavoriteGroups []*group
	RecentGroups   []*group
}

type group struct {
	Users []user
	Name  string
}

var groups map[string]*group

func Init() {
	groups = make(map[string]*group)

	db, err := DB.Connect()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GroupExists(name string) bool {
	if _, exists := groups[name]; exists {
		return true
	} else {
		return false
	}
}

// AddGroup adds a group with name
func AddGroup(name string) {
	groups[name] = &group{
		Users: nil,
		Name:  name,
	}
}

func AddUserToGroup(name string, conn *websocket.Conn, grpName string) {
	newUser := user{
		Name:   name,
		WsConn: conn,
	}
	currentUsers := groups[grpName].Users
	users := append(currentUsers, newUser)
	groups[grpName].Users = users
}

func GetConnectionsInGroup(grpName string) (conn []*websocket.Conn) {
	users := groups[grpName].Users
	for _, user := range users {
		conn = append(conn, user.WsConn)
	}
	return
}
