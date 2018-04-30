package models

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/gorilla/websocket"
)

type group struct {
	Users []*user
	Name  string
}

// GroupExists returns whether or not group "name" exists
func GroupExists(name string) (exists bool) {
	_, exists = groups[name]
	return
}

// UserExistsInGroup returns whether or not a user exists in a group, and an error if the group does not exist
func UserExistsInGroup(token string, grpName string) (b bool) {
	b = false
	if !GroupExists(grpName) {
		return
	}
	group := groups[grpName]
	grpUsers := group.Users
	usr := users[token]
	for _, u := range grpUsers {
		if u == usr {
			b = true
			return
		}
	}
	return
}

// AddGroup adds a group with name
func AddGroup(name string) (groupid string) {
	for {
		// Generate random 4 digits
		randFour := 1000 + rand.Intn(9999-1000)
		randID := strconv.Itoa(randFour)
		groupid = name + "_" + randID

		// Add group if it doesn't already exist
		if !GroupExists(groupid) {
			groups[groupid] = &group{
				Users: nil,
				Name:  groupid,
			}
			return
		}
	}
}

// AddUserToGroup adds a user to a group
func AddUserToGroup(token string, grpName string) error {
	if !UserExists(token) {
		return errors.New("AddUserToGroup - user given does not exist")
	}
	if !GroupExists(grpName) {
		return errors.New("AddUserToGroup - group given does not exist")
	}
	if UserExistsInGroup(token, grpName) {
		return errors.New("AddUserToGroup - user given already exists in group")
	}
	newUser := users[token]
	grp := groups[grpName]
	grp.addUser(newUser)
	return nil
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

func GetOtherConnectionsInGroup(token string, grpName string) (conn []*websocket.Conn) {
	if !GroupExists(grpName) {
		return nil
	}
	users := groups[grpName].Users
	for _, user := range users {
		if user.Token != token && user.WsConn != nil {
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
