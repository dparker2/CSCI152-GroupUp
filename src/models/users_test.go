package models

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestnewUser(t *testing.T) {
	testUsername := "testusername"
	testUser := newUser(testUsername)
	if testUser.Token == "" {
		t.Error("Token not set properly in NewUser")
	}
	if testUser.Name != testUsername {
		t.Error("Username not set properly in NewUser")
	}
}

func TestRemoveUser(t *testing.T) {
	users["testtoken"] = &user{
		Name: "testUsername",
	}
	RemoveUser("testtoken")
	if UserExists("testtoken") {
		t.Error("RemoveUser did not remove the user.")
	}
}

func TestUserExists(t *testing.T) {
	users["testtoken"] = &user{
		Name: "testUsername",
	}
	b := UserExists("testtoken")
	if !b {
		t.Error("User created but UserExists returns false.")
	}
	if b != (users["testtoken"] != nil) {
		t.Error("UserExists returns wrong value.")
	}
}

func TestUserExistsByUsername(t *testing.T) {
	users["testtoken"] = &user{
		Name: "testUsername",
	}
	userTokens["testUsername"] = "testtoken"
	b := UserExistsByUsername("testUsername")
	if !b {
		t.Error("User created but UserExistsByUsername returns false.")
	}
	if b != (users["testtoken"] != nil) {
		t.Error("UserExists returns wrong value.")
	}
}

func TestGetUsername(t *testing.T) {
	users["testtoken"] = &user{
		Name: "testUsername",
	}
	name := GetUsername("testtoken")
	if name != "testUsername" {
		t.Error("GetUsername returned wrong value.")
	}
}

func TestGetConnection(t *testing.T) {
	conn := &websocket.Conn{}
	users["testtoken"] = &user{
		WsConn: conn,
	}
	gotConn := GetConnection("testtoken")
	if gotConn != conn {
		t.Error("GetConnection returned wrong value.")
	}
}

func TestGetConnectionByUsername(t *testing.T) {
	conn := &websocket.Conn{}
	users["testtoken"] = &user{
		WsConn: conn,
	}
	userTokens["testUsername"] = "testtoken"
	gotConn := GetConnectionByUsername("testUsername")
	if gotConn != conn {
		t.Error("GetConnection returned wrong value.")
	}
}

func TestGetCurrentGroups(t *testing.T) {
	grps := groupList{}
	grp := &group{
		Name: "testgroup",
	}
	grps.add(grp)
	users["testtoken"] = &user{
		CurrentGroups: grps,
	}
	gotGrps := GetCurrentGroups("testtoken")
	if !(gotGrps[0] == "testgroup") {
		t.Error("GetCurrentGroups returned wrong value.")
	}
}

func TestGetOnlineFriendsList(t *testing.T) {
	usrList := userList{}
	usr := &user{
		Name: "testfriend",
	}
	usrList.add(usr)
	users["testtoken"] = &user{
		OnlineFriends: usrList,
	}
	gotFriends := GetOnlineFriendsList("testtoken")
	if !(gotFriends[0] == "testfriend") {
		t.Error("GetCurrentGroups returned wrong value.")
	}
}

func TestGetOfflineFriendsList(t *testing.T) {
	var usrList []string
	usrList = append(usrList, "testfriend")
	users["testtoken"] = &user{
		AllFriends: usrList,
	}
	gotFriends := GetOfflineFriendsList("testtoken")
	if !(gotFriends[0] == "testfriend") {
		t.Error("GetOfflineFriendsList returned wrong value.")
	}
}

func TestSetUserStatus(t *testing.T) {
	users["testtoken"] = &user{}
	SetUserStatus("testtoken", 1)
	if users["testtoken"].Status != 1 {
		t.Error("SetUserStatus did not set value correctly.")
	}
}

func TestSetUserConn(t *testing.T) {
	users["testtoken"] = &user{}
	conn := &websocket.Conn{}
	SetUserConn("testtoken", conn)
	if users["testtoken"].WsConn != conn {
		t.Error("SetUserConn did not set value correctly.")
	}
}
