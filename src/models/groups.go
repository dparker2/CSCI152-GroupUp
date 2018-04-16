package models

import "github.com/gorilla/websocket"

type group struct {
	Users []*user
	Name  string
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

func AddUserToGroup(token string, grpName string) {
	if !GroupExists(grpName) {
		AddGroup(grpName)
	}
	if UserExists(token) {
		currentUsers := groups[grpName].Users
		newUser := users[token]
		users := append(currentUsers, newUser)
		groups[grpName].Users = users
	}
}

func GetConnectionsInGroup(grpName string) (conn []*websocket.Conn) {
	if !GroupExists(grpName) {
		return nil
	}
	users := groups[grpName].Users
	for _, user := range users {
		if user.WsConn != nil {
			conn = append(conn, user.WsConn)
		}
	}
	return
}
