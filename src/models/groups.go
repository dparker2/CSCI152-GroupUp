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
	// Delete this when create group functionality is there (DB setup too)
	if !GroupExists(grpName) {
		AddGroup(grpName)
	}
	if UserExists(token) {
		newUser := users[token]
		grp := groups[grpName]
		grp.addUser(newUser)
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

func (g *group) addUser(u *user) {
	currentUsers := g.Users
	g.Users = append(currentUsers, u)
	return
}
