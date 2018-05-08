package models

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestNewUser(t *testing.T) {
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
		t.Error("UserExistsByUsername returns wrong value.")
	}

	b2 := UserExistsByUsername("nonexistentUsername")
	if b2 {
		t.Error("UserExistsByUsername returns true even though user doesn't exist.")
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

func TestUserHasCurrentGroups(t *testing.T) {
	grps := groupList{}
	grp := &group{
		Name: "testgroup",
	}
	grps.add(grp)
	groups["testgroup"] = grp
	groups["testgroup"] = &group{
		Name: "testgroup",
	}
	users["testtoken"] = &user{
		CurrentGroups: grps,
	}
	b := UserHasCurrentGroup("testtoken", "testgroup")
	if !b {
		t.Error("GetCurrentGroups returned wrong value.")
	}
}

func TestAddFriendToUser(t *testing.T) {
	users["testtoken"] = &user{}
	usr := users["testtoken"]
	friend := "Larry"
	AddFriendToUser("testtoken", friend)
	if usr.AllFriends[0] != friend {
		t.Error("AddFriendToUser did not set value correctly.")
	}
	err := AddFriendToUser("testtoken", friend)
	if err == nil {
		t.Error("AddFriendToUser did not return error when adding friend that has already been added.")
	}
}

func TestURemove(t *testing.T) {
	usrs := userList{}
	usr := &user{
		Name: "testname",
	}
	usrs = append(usrs, usr)
	length1 := len(usrs)
	usrs.remove(usr)
	length2 := len(usrs)
	if length1 == length2 {
		t.Error("User.remove did not remove user from users list.")
	}
}

func TestUContains(t *testing.T) {
	usrs := userList{}
	usr := &user{}
	usr2 := &user{}
	usrs = append(usrs, usr)
	if !usrs.contains(usr) {
		t.Error("User.contains did not find user in users list.")
	}
	if usrs.contains(usr2) {
		t.Error("User.contains found user in users list that doesn't actually exist in users list.")
	}
}

func TestUContainsUsername(t *testing.T) {
	usrs := userList{}
	usr := &user{
		Name: "testname",
	}
	usrs = append(usrs, usr)
	if !usrs.containsUsername(usr.Name) {
		t.Error("User.containsUsername did not find username in users list.")
	}
	if usrs.containsUsername("wrongname") {
		t.Error("User.containsUsername found username in users list that doesn't actually exist in users list.")
	}

}
