package models

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/websocket"
)

type group struct {
	Users userList
	Name  string
	Admin string
	Mutex sync.Mutex
}
type groupList []*group

var groupsMutex = &sync.Mutex{}

// GroupExists returns whether or not group "name" exists
func GroupExists(name string) (exists bool) {
	_, exists = groups[name]
	if exists == false { //If group doesn't exists in group struct, double checks DB to confirm.
		var groupexists string
		err := db.QueryRow("SELECT table_name FROM information_schema.tables WHERE TABLE_SCHEMA='StudyGroup' AND table_name= ?", name).Scan(&groupexists)
		switch {
		case err == sql.ErrNoRows:
			fmt.Println("Group does not exist")
		case err != nil:
			panic(err)
		default:
			if name == groupexists {
				fmt.Println("Group" + groupexists + "exists in the DB already")
				exists = true
			}
		}
	}
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
func AddGroup(name string, token string) (groupid string) {
	for {
		// Generate random 4 digits
		randFour := 1000 + rand.Intn(9999-1000)
		fmt.Println(randFour)
		randID := strconv.Itoa(randFour)
		groupid = name + "_" + randID
		username := GetUsername(token)

		// Add group if it doesn't already exist
		if !GroupExists(groupid) {
			groupsMutex.Lock()
			defer groupsMutex.Unlock()

			createGroupInDB(groupid)
			putAdminInGroupDB(groupid, username)

			groups[groupid] = &group{
				Users: nil,
				Name:  groupid,
				Admin: username,
			}
			fmt.Println(groups[groupid])
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

func RemoveUserFromGroup(token string, grpName string) error {
	if !UserExists(token) {
		return errors.New("RemoveUserFromGroup - user given does not exist")
	}
	if !GroupExists(grpName) {
		return errors.New("RemoveUserFromGroup - group given does not exist")
	}
	if !UserExistsInGroup(token, grpName) {
		return errors.New("RemoveUserFromGroup - user given does not exist in group")
	}
	u := users[token]
	grp := groups[grpName]
	grp.removeUser(u)
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
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	currentUsers := g.Users
	g.Users = append(currentUsers, u)
	return
}

func (g *group) removeUser(u *user) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	for i := range g.Users {
		if g.Users[i] == u {
			g.Users[len(g.Users)-1], g.Users[i] = g.Users[i], g.Users[len(g.Users)-1]
			g.Users = g.Users[:len(g.Users)-1]
			break
		}
	}
}

func (gl groupList) add(g *group) {
	gl = append(gl, g)
}

// Remove first occurance of g
func (gl groupList) remove(g *group) {
	for i := range gl {
		if gl[i] == g {
			gl[len(gl)-1], gl[i] = gl[i], gl[len(gl)-1]
			gl = gl[:len(gl)-1]
			break
		}
	}
}

func (gl groupList) contains(g *group) (b bool) {
	b = false
	for i := range gl {
		if gl[i] == g {
			b = true
			return
		}
	}
	return
}
