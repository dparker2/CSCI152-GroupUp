package models

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type user struct {
	Name          string
	Email         string
	UserID        int
	LockoutStatus sql.NullInt64
	SecQuestions  [3]sql.NullString
	SecAnswers    [3]sql.NullString
	WsConn        *websocket.Conn
	Status        int
	Token         string
	OnlineFriends userList
	AllFriends    []string
	CurrentGroups groupList
}
type userList []*user

var userMutex = &sync.Mutex{}

// NewUser makes a new user with the given username, token, and IP
func newUser(username string) (u user) {
	userToken, err := generateRandomString(32)
	if err != nil {
		log.Println(err.Error())
	}
	u = user{
		Name:  username,
		Token: userToken,
	}
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

func UserExistsByUsername(username string) bool {
	if token, exists := userTokens[username]; exists {
		return UserExists(token)
	} else {
		return false
	}
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

func GetConnection(token string) (conn *websocket.Conn) {
	userMutex.Lock()
	if UserExists(token) {
		conn = users[token].WsConn
	}
	userMutex.Unlock()
	return
}

func GetConnectionByUsername(username string) (conn *websocket.Conn) {
	userMutex.Lock()
	if UserExistsByUsername(username) {
		conn = users[userTokens[username]].WsConn
	}
	userMutex.Unlock()
	return
}

func GetCurrentGroups(token string) (list []string) {
	usr := users[token]
	currGrps := usr.CurrentGroups
	for _, grp := range currGrps {
		list = append(list, grp.Name)
	}
	return
}

func GetOnlineFriendsList(token string) (list []string) {
	for _, usr := range users[token].OnlineFriends {
		list = append(list, usr.Name)
	}
	return
}

func GetOfflineFriendsList(token string) (list []string) {
	for _, friend := range users[token].AllFriends {
		if !users[token].OnlineFriends.containsUsername(friend) {
			list = append(list, friend)
		}
	}
	return
}

func GetOnlineFollowers(token string) (list []string, err error) {
	totalFollowers, err := db_GetUsersFollowersDB(users[token].UserID)
	if err != nil {
		return nil, err
	}
	for _, follower := range totalFollowers {
		if UserExistsByUsername(follower) {
			list = append(list, follower)
		}
	}
	return
}
func GetUserID(token string) (id int) {
	if UserExists(token) {
		return users[token].UserID
	}
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

func UserHasCurrentGroup(token string, grpName string) (b bool) {
	grp := groups[grpName]
	u := users[token]
	b = u.CurrentGroups.contains(grp)
	return
}

func AddFriendToUser(token string, friendname string) error {
	usr := users[token]
	uuid := usr.UserID

	for _, friend := range usr.AllFriends {
		if friend == friendname {
			return errors.New("User already friend")
		}
	}

	usr.AllFriends = append(usr.AllFriends, friendname)
	err := db_AddFriendToUser(uuid, friendname)
	return err
}

func AddGroupToUsersCurrentGroups(token string, grpName string) {
	u := users[token]
	grp := groups[grpName]
	u.CurrentGroups.add(grp)
	db_AddGroupToUsersGroups(u.UserID, grpName)
}

func RemoveGroupFromUser(token string, grpName string) (err error) {
	u := users[token]
	grp := groups[grpName]
	u.CurrentGroups.remove(grp)
	db_RemoveGroupFromUsersGroups(u.UserID, grpName)
	err = DecreaseGroupIndexSubs(grpName)

	if err != nil {
		return err
	}
	return
}

func (ul *userList) add(u *user) {
	*ul = append(*ul, u)
}

// Remove first occurance of u
func (ul *userList) remove(u *user) {
	for i := range *ul {
		if (*ul)[i].Name == u.Name {
			(*ul)[len((*ul))-1], (*ul)[i] = (*ul)[i], (*ul)[len((*ul))-1]
			(*ul) = (*ul)[:len((*ul))-1]
			break
		}
	}
}

func (ul userList) contains(u *user) (b bool) {
	b = false
	for i := range ul {
		if ul[i] == u {
			b = true
			return
		}
	}
	return
}

func (ul userList) containsUsername(u string) (b bool) {
	b = false
	for i := range ul {
		if ul[i].Name == u {
			b = true
			return
		}
	}
	return
}
